package di

import (
	"github.com/ZilDuck/indexer-api/internal/cache"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/elastic_cache"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/ZilDuck/indexer-api/internal/service"
	"github.com/gin-contrib/cache/persistence"
	"github.com/sarulabs/dingo/v4"
	"go.uber.org/zap"
	"time"
)

var Definitions = []dingo.Def{
	{
		Name: "elastic",
		Build: func() (elastic_cache.Index, error) {
			elastic, err := elastic_cache.New()
			if err != nil {
				zap.L().With(zap.Error(err)).Fatal("Failed to start ES")
			}

			return elastic, nil
		},
	},
	{
		Name: "nft.repository",
		Build: func(elastic elastic_cache.Index) (repository.NftRepository, error) {
			return repository.NewNftRepository(elastic), nil
		},
	},
	{
		Name: "tx.repository",
		Build: func(elastic elastic_cache.Index, store persistence.CacheStore) (repository.TransactionRepository, error) {
			return repository.NewTransactionRepository(elastic, store), nil
		},
	},
	{
		Name: "nft.service",
		Build: func(nftRepo repository.NftRepository) (service.NFTService, error) {
			return service.NewNFTService(nftRepo), nil
		},
	},
	{
		Name: "cache.store",
		Build: func() (persistence.CacheStore, error) {
			return persistence.NewInMemoryStore(config.Get().CacheDefaultExpiration * time.Second), nil
		},
	},
	{
		Name: "cache.validator",
		Build: func(txRepo repository.TransactionRepository, store persistence.CacheStore) (cache.Validator, error) {
			return cache.NewValidator(txRepo, store, config.Get().CacheDefaultExpiration*time.Second), nil
		},
	},
}
