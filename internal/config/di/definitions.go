package di

import (
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/elastic_cache"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/sarulabs/dingo/v4"
	"go.uber.org/zap"
)

var Definitions = []dingo.Def{
	{
		Name: "elastic",
		Build: func() (elastic_cache.Index, error) {
			elastic, err := elastic_cache.New(config.Get().ElasticSearch, config.Get().Aws)
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
}
