package driver

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/qiaopengjun5162/VRFService/bindings"
	"github.com/qiaopengjun5162/VRFService/tx_manager"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1_500_000_000)
)

type DriverEngineConfig struct {
	ChainClient               *ethclient.Client
	ChainId                   *big.Int
	VerifyRandVRFAddress      common.Address
	CallerAddress             common.Address
	PrivateKey                *ecdsa.PrivateKey // CallerAddress 和 PrivateKey 是一一对应的
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
}

type DriverEngine struct {
	Ctx                      context.Context
	Cfg                      *DriverEngineConfig
	VerifyRandVRFContract    *bindings.VerifyRandVRF
	RawVerifyRandVRFContract *bind.BoundContract
	VerifyRandVRFContractAbi *abi.ABI
	TxMgr                    tx_manager.TxManager
	cancel                   func()
	wg                       sync.WaitGroup
}

func NewDriverEngine(ctx context.Context, cfg *DriverEngineConfig) (*DriverEngine, error) {
	_, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	verifyRandVRFContract, err := bindings.NewVerifyRandVRF(cfg.VerifyRandVRFAddress, cfg.ChainClient)
	if err != nil {
		log.Error("new VerifyRandVRF fail", "err", err)
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(bindings.VerifyRandVRFMetaData.ABI))
	if err != nil {
		log.Error("parsed abi fail", "err", err)
		return nil, err
	}

	verifyRandVRFContractAbi, err := bindings.VerifyRandVRFMetaData.GetAbi()
	if err != nil {
		log.Error("get VerifyRandVRF meta data fail", "err", err)
		return nil, err
	}

	rawVerifyRandVRFContract := bind.NewBoundContract(cfg.VerifyRandVRFAddress, parsed, cfg.ChainClient, cfg.ChainClient, cfg.ChainClient)

	txManagerConfig := tx_manager.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txManager := tx_manager.NewSimpleTxManager(txManagerConfig, cfg.ChainClient)

	return &DriverEngine{
		Ctx:                      ctx,
		Cfg:                      cfg,
		VerifyRandVRFContract:    verifyRandVRFContract,
		RawVerifyRandVRFContract: rawVerifyRandVRFContract,
		VerifyRandVRFContractAbi: verifyRandVRFContractAbi,
		TxMgr:                    txManager,
		cancel:                   cancel,
	}, nil
}

func (de *DriverEngine) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	opts, err = bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true

	finalTx, err := de.RawVerifyRandVRFContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	case de.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return de.RawVerifyRandVRFContract.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

// SendTransaction sends the given transaction to the network.
//
// It's a wrapper around ethclient.Client.SendTransaction, which is used to send
// a signed transaction to the network. The context is passed through to the
// underlying ethclient call.
//
// The function returns an error if the transaction could not be sent, or nil if
// the transaction was sent successfully.
func (de *DriverEngine) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return de.Cfg.ChainClient.SendTransaction(ctx, tx)
}

func (de *DriverEngine) isMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound.Error())
}

func (de *DriverEngine) fulfillRandomWords(ctx context.Context, requestId *big.Int, randomList []*big.Int) (*types.Transaction, error) {
	nonce, err := de.Cfg.ChainClient.NonceAt(ctx, de.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get nonce error", "err", err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	tx, err := de.VerifyRandVRFContract.FulfillRandomWords(opts, requestId, randomList)
	switch {
	case err == nil:
		return tx, nil

	case de.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return de.VerifyRandVRFContract.FulfillRandomWords(opts, requestId, randomList)

	default:
		return nil, err
	}
}

func (de *DriverEngine) FulfillRandomWords(requestId *big.Int, randomList []*big.Int) (*types.Receipt, error) {
	tx, err := de.fulfillRandomWords(de.Ctx, requestId, randomList)
	if err != nil {
		log.Error("build request random words tx fail", "err", err)
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return de.UpdateGasPrice(ctx, tx)
	}

	receipt, err := de.TxMgr.Send(de.Ctx, updateGasPrice, de.SendTransaction)
	if err != nil {
		log.Error("send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil
}
