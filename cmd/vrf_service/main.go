package main

import (
	"context"
	"github.com/qiaopengjun5162/VRFService/common/opio"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

var (
	GitCommit = ""
	gitDate   = ""
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := NewCli(GitCommit, gitDate)
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed", err)
		os.Exit(1)
	}
}
