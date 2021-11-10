package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dantudor/zilkroad-txapi/internal/elastic_cache"
	"github.com/dantudor/zilkroad-txapi/internal/entity"
	"github.com/olivere/elastic/v7"
)

type NftRepository interface {
	GetNft(contractAddr string, tokenId int) (entity.NFT, error)
	GetForAddress(ownerAddr string, from, size int) ([]entity.NFT, int64, error)
}

type nftRepository struct {
	elastic *elastic_cache.Index
}

var (
	ErrNftNotFound = errors.New("nft not found")
)

func NewNftRepository(elastic *elastic_cache.Index) NftRepository {
	return nftRepository{elastic: elastic}
}

func (nftRepo nftRepository) GetNft(contractAddr string, tokenId int) (entity.NFT, error) {
	query := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("contract.keyword", contractAddr),
		elastic.NewTermQuery("tokenId", tokenId),
	)

	result, err := nftRepo.elastic.Client.
		Search(elastic_cache.NftIndex.Get()).
		Query(query).
		Size(1).
		Do(context.Background())

	return nftRepo.findOne(result, err)
}

func (nftRepo nftRepository) GetForAddress(ownerAddr string, from, size int) ([]entity.NFT, int64, error) {
	result, err := nftRepo.elastic.Client.
		Search(elastic_cache.NftIndex.Get()).
		Query(elastic.NewTermQuery("owner.keyword", ownerAddr)).
		Size(size).
		From(from).
		TrackTotalHits(true).
		Do(context.Background())

	return nftRepo.findMany(result, err)
}

func (nftRepo nftRepository) findOne(results *elastic.SearchResult, err error) (entity.NFT, error) {
	if err != nil {
		return entity.NFT{}, err
	}

	if len(results.Hits.Hits) == 0 {
		return entity.NFT{}, ErrNftNotFound
	}

	var nft entity.NFT
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &nft)

	return nft, err
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
