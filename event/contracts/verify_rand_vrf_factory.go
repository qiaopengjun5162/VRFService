package contracts

import (
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/qiaopengjun5162/VRFService/bindings"
	"github.com/qiaopengjun5162/VRFService/database"
	"github.com/qiaopengjun5162/VRFService/database/event"
	"github.com/qiaopengjun5162/VRFService/database/worker"
)

type VerifyRandVRFFactory struct {
	VerifyRandVRFFactoryAbi    *abi.ABI
	VerifyRandVRFFactoryFilter *bindings.VerifyRandVRFFactoryFilterer
}

func NewVerifyRandVRFFactory() (*VerifyRandVRFFactory, error) {
	verifyRandVRFFactoryAbi, err := bindings.VerifyRandVRFFactoryMetaData.GetAbi()
	if err != nil {
		log.Error("get VerifyRandVRFFactory vrf factory abi fail", "err", err)
		return nil, err
	}
	verifyRandVRFFactoryFilter, err := bindings.NewVerifyRandVRFFactoryFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new VerifyRandVRFFactory vrf factory filter fail", "err", err)
		return nil, err
	}
	return &VerifyRandVRFFactory{
		VerifyRandVRFFactoryAbi:    verifyRandVRFFactoryAbi,
		VerifyRandVRFFactoryFilter: verifyRandVRFFactoryFilter,
	}, nil
}

func (vff *VerifyRandVRFFactory) ProcessVerifyRandVRFFactoryEvent(db *database.DB, verifyRandVRFFactoryAddress string, startHeight, endHeight *big.Int) ([]worker.PoxyCreated, error) {
	var proxyCreatedList []worker.PoxyCreated
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(verifyRandVRFFactoryAddress)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startHeight, endHeight)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return proxyCreatedList, err
	}
	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == vff.VerifyRandVRFFactoryAbi.Events["ProxyCreated"].ID.String() {
			proxyCreated, err := vff.VerifyRandVRFFactoryFilter.ParseProxyCreated(*contractEvent.RLPLog)
			if err != nil {
				log.Error("proxy created fail", "err", err)
				return proxyCreatedList, err
			}
			log.Info("proxy created event", "MintProxyAddress", proxyCreated.MintProxyAddress)
			pc := worker.PoxyCreated{
				GUID:         uuid.New(),
				ProxyAddress: proxyCreated.MintProxyAddress,
				Timestamp:    uint64(time.Now().Unix()),
			}
			proxyCreatedList = append(proxyCreatedList, pc)
		}
	}
	return proxyCreatedList, nil
}
