package cache

import (
	"github.com/dantudor/zilkroad-txapi/internal/repository"
	"github.com/gin-contrib/cache/persistence"
	"go.uber.org/zap"
	"time"
)

type Validator interface {
	Subscribe()
}

type validator struct {
	txRepo               repository.TransactionRepository
	store                persistence.CacheStore
	defaultCacheDuration time.Duration
}

func NewValidator(txRepo repository.TransactionRepository, store persistence.CacheStore, defaultCacheDuration time.Duration) Validator {
	return validator{txRepo, store, defaultCacheDuration}
}

func (v validator) Subscribe() {
	zap.L().Info("Cache Validator Subscribe")

	var bestBlockNum = uint64(0)
	if err := v.store.Set("bestBlockNum", bestBlockNum, v.defaultCacheDuration); err != nil {
		zap.L().Error("Failed to set Best BlockNum cache value")
	}

	for {
		latestBlockNum, err := v.txRepo.GetBestBlock()
		if err != nil {
			zap.L().With(zap.Error(err)).Error("Failed to get best block")
			time.Sleep(10 * time.Second)
			continue
		}
		if err := v.store.Get("bestBlockNum", &bestBlockNum); err != nil {
			bestBlockNum = 0
		}

		if latestBlockNum != bestBlockNum {
			zap.S().Infof("New best block found: %d", bestBlockNum)

			if err := v.store.Flush(); err != nil {
				zap.L().With(zap.Error(err)).Warn("Failed to flush cache")
			}

			if err := v.store.Set("bestBlockNum", latestBlockNum, v.defaultCacheDuration); err != nil {
				zap.L().With(zap.Error(err)).Error("CacheValidator: Failed to set new best block num")
			}
		}
		time.Sleep(2 * time.Second)
	}
}
