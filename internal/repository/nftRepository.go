package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_cache"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"time"
)

type NftRepository interface {
	GetForAddress(ownerAddr string, size, page int) ([]entity.NFT, int64, error)
}

type nftRepository struct {
	elastic elastic_cache.Index
}

var (
	ErrNftNotFound = errors.New("nft not found")
)

func NewNftRepository(elastic elastic_cache.Index) NftRepository {
	return nftRepository{elastic: elastic}
}

func (nftRepo nftRepository) GetForAddress(ownerAddr string, size, page int) ([]entity.NFT, int64, error) {
	from := size*page - size

	result, err := search(nftRepo.elastic.Client.
		Search(elastic_cache.NftIndex.Get()).
		Query(elastic.NewTermQuery("owner.keyword", ownerAddr)).
		Size(size).
		From(from).
		TrackTotalHits(true))

	return nftRepo.findMany(result, err)
}

func (nftRepo nftRepository) findOne(results *elastic.SearchResult, err error) (*entity.NFT, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrNftNotFound
	}

	var nft entity.NFT
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &nft)

	return &nft, err
}

func search(searchService *elastic.SearchService) (*elastic.SearchResult, error) {
	result, err := searchService.Do(context.Background())
	if err != nil && err.Error() == "elastic: Error 429 (Too Many Requests)" {
		zap.L().Warn("Elastic: 429 (Too Many Requests)")
		time.Sleep(5 * time.Second)
		return search(searchService)
	}

	return result, err
}

func (nftRepo nftRepository) findMany(results *elastic.SearchResult, err error) ([]entity.NFT, int64, error) {
	nfts := make([]entity.NFT, 0)

	if err != nil {
		return nfts, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var nft entity.NFT
		if err := json.Unmarshal(hit.Source, &nft); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, results.TotalHits(), nil
}
