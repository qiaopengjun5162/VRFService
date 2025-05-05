package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/qiaopengjun5162/VRFService/common/tasks"
	"github.com/qiaopengjun5162/VRFService/database"
	"github.com/qiaopengjun5162/VRFService/driver"
)

type Config struct {
	LoopInterval time.Duration
}

type Worker struct {
	workerConfig *Config
	db           *database.DB
	deg          *driver.DriverEngine

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewWorker(db *database.DB, deg *driver.DriverEngine, workerConfig *Config, shutdown context.CancelCauseFunc) (*Worker, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &Worker{
		db:             db,
		deg:            deg,
		workerConfig:   workerConfig,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (wk *Worker) Start() error {
	log.Info("starting worker processor...")
	tickerEventWorker := time.NewTicker(wk.workerConfig.LoopInterval)
	wk.tasks.Go(func() error {
		for range tickerEventWorker.C {
			log.Info("start handler random for vrf")
			err := wk.ProcessCallerVrf()
			if err != nil {
				log.Error("process caller vrf fail", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (wk *Worker) Close() error {
	wk.resourceCancel()
	return wk.tasks.Wait()
}

func (wk *Worker) ProcessCallerVrf() error {
	// 1. 获取 RequestSent 合约事件
	requestUnsentList, err := wk.db.RequestSend.QueryUnHandleRequestSendList()
	if err != nil {
		log.Error("query unhandled request sent list fail", "err", err)
		return err
	}
	for _, requestUnsent := range requestUnsentList {
		// 2. 组装随机数据列表交易发到 Vrf 合约
		txReceipt, err := wk.deg.FulfillRandomWords(requestUnsent.RequestId, nil)
		if err != nil {
			log.Error("Fulfill random words fail", "err", err)
			return err
		}
		if txReceipt.Status == 1 {
			// 3. 更新 RequestSent 合约事件表的状态
			err := wk.db.RequestSend.MarkRequestSendFinish(requestUnsent)
			if err != nil {
				log.Error("mark request sent event list fail", "err", err)
				return err
			}
		}
	}
	return nil
}
