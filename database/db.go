package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/qiaopengjun5162/VRFService/config"
	"github.com/qiaopengjun5162/VRFService/database/common"
	"github.com/qiaopengjun5162/VRFService/database/event"
	_ "github.com/qiaopengjun5162/VRFService/database/utils/serializers"
	"github.com/qiaopengjun5162/VRFService/database/worker"
	"github.com/qiaopengjun5162/VRFService/synchronizer/retry"
)

type DB struct {
	gorm *gorm.DB

	Blocks          common.BlocksDB
	ContractEvent   event.ContractEventDB
	EventBlocks     worker.EventBlocksDB
	FillRandomWords worker.FillRandomWordsDB
	RequestSend     worker.RequestSendDB
	PoxyCreated     worker.PoxyCreatedDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})

	if err != nil {
		return nil, err
	}

	db := &DB{
		gorm:            gorm,
		Blocks:          common.NewBlocksDB(gorm),
		ContractEvent:   event.NewContractEventsDB(gorm),
		EventBlocks:     worker.NewEventBlocksDB(gorm),
		FillRandomWords: worker.NewFillRandomWordsDB(gorm),
		RequestSend:     worker.NewRequestSendDB(gorm),
		PoxyCreated:     worker.NewPoxyCreatedDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:            tx,
			Blocks:          common.NewBlocksDB(tx),
			ContractEvent:   event.NewContractEventsDB(tx),
			EventBlocks:     worker.NewEventBlocksDB(tx),
			FillRandomWords: worker.NewFillRandomWordsDB(tx),
			RequestSend:     worker.NewRequestSendDB(tx),
			PoxyCreated:     worker.NewPoxyCreatedDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}

		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
