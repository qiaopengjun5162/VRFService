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

type VerifyRandVRF struct {
	VerifyRandVRFAbi    *abi.ABI
	VerifyRandVRFFilter *bindings.VerifyRandVRFFilterer
}

func NewVerifyRandVRF() (*VerifyRandVRF, error) {
	verifyRandVRFAbi, err := bindings.VerifyRandVRFMetaData.GetAbi()
	if err != nil {
		log.Error("get VerifyRandVRF abi fail", "err", err)
		return nil, err
	}
	verifyRandVRFFilterer, err := bindings.NewVerifyRandVRFFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new VerifyRandVRF filter fail", "err", err)
		return nil, err
	}
	return &VerifyRandVRF{
		VerifyRandVRFAbi:    verifyRandVRFAbi,
		VerifyRandVRFFilter: verifyRandVRFFilterer,
	}, nil
}

func (vf *VerifyRandVRF) ProcessVerifyRandVRFEvent(db *database.DB, verifyRandVRFAddress string, startHeight, endHeight *big.Int) ([]worker.RequestSend, []worker.FillRandomWords, error) {
	var RequestSentList []worker.RequestSend
	var FillRandomWordList []worker.FillRandomWords
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(verifyRandVRFAddress)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startHeight, endHeight)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return RequestSentList, FillRandomWordList, err
	}

	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == vf.VerifyRandVRFAbi.Events["RequestSent"].ID.String() {
			requestSentEvent, err := vf.VerifyRandVRFFilter.ParseRequestSent(*contractEvent.RLPLog)
			if err != nil {
				log.Error("parse request sent fail", "err", err)
				return RequestSentList, FillRandomWordList, err
			}
			log.Info("Request sent event", "RequestId", requestSentEvent.RequestId, "NumWords", requestSentEvent.NumWords, "Current", requestSentEvent.Current)
			rs := worker.RequestSend{
				GUID:       uuid.New(),
				RequestId:  requestSentEvent.RequestId,
				VrfAddress: requestSentEvent.Current,
				NumWords:   requestSentEvent.NumWords,
				Status:     0,
				Timestamp:  uint64(time.Now().Unix()),
			}
			RequestSentList = append(RequestSentList, rs)
		}

		if contractEvent.EventSignature.String() == vf.VerifyRandVRFAbi.Events["FillRandomWords"].ID.String() {
			fillRandomWords, err := vf.VerifyRandVRFFilter.ParseFillRandomWords(*contractEvent.RLPLog)
			if err != nil {
				log.Error("parse fill random fail", "err", err)
				return RequestSentList, FillRandomWordList, err
			}
			log.Info("Fill random words event", "RequestId", fillRandomWords.RequestId, "RandomWords", fillRandomWords.RandomWords)
			var randomWords string
			for _, rWord := range fillRandomWords.RandomWords {
				randomWords = rWord.String()
			}
			frw := worker.FillRandomWords{
				GUID:        uuid.New(),
				RequestId:   fillRandomWords.RequestId,
				RandomWords: randomWords,
				Timestamp:   uint64(time.Now().Unix()),
			}
			FillRandomWordList = append(FillRandomWordList, frw)
		}
	}
	return RequestSentList, FillRandomWordList, nil
}
