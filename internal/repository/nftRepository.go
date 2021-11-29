package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_cache"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type NftRepository interface {
	GetForAddress(network, ownerAddr string) ([]entity.NftOwner, error)
	GetForContract(network, contractAddr string, size, page int) ([]entity.NFT, int64, error)
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

func (nftRepo nftRepository) GetForAddress(network, ownerAddr string) ([]entity.NftOwner, error) {
	contracts := make([]entity.NftOwner, 0)

	query := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("owner.keyword", ownerAddr),
		elastic.NewTermQuery("burnedAt", 0),
	)

	agg := elastic.NewFilterAggregation().Filter(query).
		SubAggregation("contract", elastic.NewTermsAggregation().Field("contract.keyword").Size(10000).
			SubAggregation("tokenId", elastic.NewTermsAggregation().Field("tokenId").Size(10000)))

	results, err := search(nftRepo.elastic.Client.
		Search(elastic_cache.NftIndex.Get(network)).
		Aggregation("owner", agg).
		Size(0))

	if err != nil {
		return contracts, err
	}

	if ownerAgg, found := results.Aggregations.Filter("owner"); found {
		if contractAgg, found := ownerAgg.Aggregations.Terms("contract"); found {
			for _, contractsBucket := range contractAgg.Buckets {
				contract := entity.NftOwner{Address: contractsBucket.Key.(string)}

				if tokenIdBucket, found := contractsBucket.Terms("tokenId"); found {
					for _, tokenId := range tokenIdBucket.Buckets {
						contract.TokenIds = append(contract.TokenIds, uint64(tokenId.Key.(float64)))
					}
				}
				contracts = append(contracts, contract)
			}
		}
	}

	return contracts, nil
}

func (nftRepo nftRepository) GetForContract(network, contractAddr string, size, page int) ([]entity.NFT, int64, error) {
	from := size*page - size

	result, err := search(nftRepo.elastic.Client.
		Search(elastic_cache.NftIndex.Get(network)).
		Query(elastic.NewTermQuery("contract.keyword", contractAddr)).
		Size(size).
		Sort("tokenId", true).
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
