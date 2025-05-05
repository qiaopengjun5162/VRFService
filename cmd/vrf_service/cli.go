package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/version"

	vrf "github.com/qiaopengjun5162/VRFService"
	"github.com/qiaopengjun5162/VRFService/common/cliapp"
	"github.com/qiaopengjun5162/VRFService/common/opio"
	"github.com/qiaopengjun5162/VRFService/config"
	"github.com/qiaopengjun5162/VRFService/database"
	flag2 "github.com/qiaopengjun5162/VRFService/flags"
)

// This function runs the VerifyRandVRF function and returns a Lifecycle and an error
func runVerifyRandVR(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	// Log that the vrf function is being run
	log.Info("run  vrf")
	// Load the config from the context
	cfg, err := config.LoadConfig(ctx)
	// If there is an error, log it and return the error
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	// Return the new VerifyRandVRF function with the context, config, and shutdown function
	return vrf.NewVerifyRandVRF(ctx.Context, &cfg, shutdown)
}

func runMigrations(ctx *cli.Context) error {
	log.Info("Running migrations...")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
}

func NewCli(GitCommit string, gitDate string) *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              withCommit(GitCommit, gitDate),
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runVerifyRandVR),
			},
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}

// Semantic holds the textual version string for major.minor.patch.
var Semantic = fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)

// WithMeta holds the textual version string including the metadata.
var WithMeta = func() string {
	v := Semantic
	if version.Meta != "" {
		v += "-" + version.Meta
	}
	return v
}()

func withCommit(gitCommit, gitDate string) string {
	vsn := WithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	if (version.Meta != "stable") && (gitDate != "") {
		vsn += "-" + gitDate
	}
	return vsn
}
