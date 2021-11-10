package di

import (
	"github.com/dantudor/zilkroad-txapi/internal/elastic_cache"
	"github.com/dantudor/zilkroad-txapi/internal/repository"
	"github.com/dantudor/zilkroad-txapi/internal/service"
	"github.com/sarulabs/dingo/v3"
	"go.uber.org/zap"
)

var Definitions = []dingo.Def{
	{
		Name: "elastic",
		Build: func() (*elastic_cache.Index, error) {
			elastic, err := elastic_cache.New()
			if err != nil {
				zap.L().With(zap.Error(err)).Fatal("Failed to start ES")
			}

			return elastic, nil
		},
	},
	{
		Name: "nft.repository",
		Build: func(elastic *elastic_cache.Index) (repository.NftRepository, error) {
			return repository.NewNftRepository(elastic), nil
		},
	},
	{
		Name: "nft.service",
		Build: func(nftRepo repository.NftRepository) (service.NFTService, error) {
			return service.NewNFTService(nftRepo), nil
		},
	},
}
