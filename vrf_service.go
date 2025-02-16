package VRFService

import (
	"context"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/qiaopengjun5162/VRFService/common"
	"github.com/qiaopengjun5162/VRFService/config"
	"github.com/qiaopengjun5162/VRFService/database"
	"github.com/qiaopengjun5162/VRFService/driver"
	"github.com/qiaopengjun5162/VRFService/event"
	"github.com/qiaopengjun5162/VRFService/synchronizer"
	"github.com/qiaopengjun5162/VRFService/synchronizer/node"
	"github.com/qiaopengjun5162/VRFService/worker"
)

type VerifyRandVRF struct {
	db *database.DB

	synchronizer  *synchronizer.Synchronizer
	eventsHandler *event.EventsHandler
	worker        *worker.Worker

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewVerifyRandVRF(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*VerifyRandVRF, error) {
	ethClient, err := node.DialEthClient(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth client fail", "err", err)
		return nil, err
	}

	db, err := database.NewDB(ctx, cfg.MasterDB)
	if err != nil {
		log.Error("new database fail", "err", err)
		return nil, err
	}

	synchronizerS, err := synchronizer.NewSynchronizer(cfg, db, ethClient, shutdown)
	if err != nil {
		log.Error("new synchronizer fail", "err", err)
		return nil, err
	}

	eventConfig := &event.EventsHandlerConfig{
		VerifyRandVRFAddress:        cfg.Chain.VerifyRandVRFContractAddress,
		VerifyRandVRFFactoryAddress: cfg.Chain.VerifyRandVRFFactoryContractAddress,
		LoopInterval:                cfg.Chain.EventInterval,
		StartHeight:                 big.NewInt(int64(cfg.Chain.StartingHeight)),
		Epoch:                       500,
	}

	eventHandler, err := event.NewEventsHandler(db, eventConfig, shutdown)
	if err != nil {
		return nil, err
	}

	ethclient, err := driver.EthClientWithTimeout(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth cli fail", "err", err)
		return nil, err
	}

	callerPrivateKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"ContractCaller", cfg.Chain.Mnemonic, cfg.Chain.CallerHDPath,
		cfg.Chain.PrivateKey, cfg.Chain.VerifyRandVRFContractAddress, cfg.Chain.Passphrase,
	)

	driverConfig := &driver.DriverEngineConfig{
		ChainClient:               ethclient,
		ChainId:                   big.NewInt(int64(cfg.Chain.ChainId)),
		VerifyRandVRFAddress:      common.HexToAddress(cfg.Chain.VerifyRandVRFContractAddress),
		CallerAddress:             common.HexToAddress(cfg.Chain.CallerAddress),
		PrivateKey:                callerPrivateKey,
		NumConfirmations:          cfg.Chain.Confirmations,
		SafeAbortNonceTooLowCount: cfg.Chain.SafeAbortNonceTooLowCount,
	}
	engine, err := driver.NewDriverEngine(ctx, driverConfig)
	if err != nil {
		log.Error("new driver engine fail", "err", err)
		return nil, err
	}

	workerConfig := &worker.Config{
		LoopInterval: cfg.Chain.CallInterval,
	}

	workerProcessor, err := worker.NewWorker(db, engine, workerConfig, shutdown)
	if err != nil {
		log.Error("new event processor fail", "err", err)
		return nil, err
	}

	return &VerifyRandVRF{
		db:            db,
		synchronizer:  synchronizerS,
		eventsHandler: eventHandler,
		worker:        workerProcessor,
		shutdown:      shutdown,
	}, nil
}

func (vrf *VerifyRandVRF) Start(ctx context.Context) error {
	err := vrf.synchronizer.Start()
	if err != nil {
		return err
	}
	err = vrf.eventsHandler.Start()
	if err != nil {
		return err
	}
	err = vrf.worker.Start()
	if err != nil {
		return err
	}
	return nil
}

func (vrf *VerifyRandVRF) Stop(ctx context.Context) error {
	err := vrf.synchronizer.Close()
	if err != nil {
		return err
	}
	err = vrf.eventsHandler.Close()
	if err != nil {
		return err
	}
	err = vrf.worker.Close()
	if err != nil {
		return err
	}
	return nil
}

func (vrf *VerifyRandVRF) Stopped() bool {
	return vrf.stopped.Load()
}
