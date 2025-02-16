package config

import "github.com/ethereum/go-ethereum/common"

const (
	VerifyRandVRF = "0x9BA23eDAdC4A8c4Ee896B736bCCBafe2A18da2D2"
)

var (
	VerifyRandVRFAddr = common.HexToAddress(VerifyRandVRF)
)
