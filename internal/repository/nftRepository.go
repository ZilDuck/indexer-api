package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type NftRepository interface {
	GetForAddress(network, ownerAddr string, shape string) ([]entity.NftOwner, error)
	GetForContract(network, contractAddr string, size, offset uint64) ([]entity.Nft, int64, error)
	GetForContractByTokenId(network, contractAddr string, tokenId uint64) (*entity.Nft, error)
	GetForContractAttributes(network, contractAddr string) (entity.Attributes, error)
}

type nftRepository struct {
	elastic elastic_search.Index
}

var (
	ErrNftNotFound = errors.New("nft not found")
)

func NewNftRepository(elastic elastic_search.Index) NftRepository {
	return nftRepository{elastic: elastic}
}

func (nftRepo nftRepository) GetForAddress(network, ownerAddr string, shape string) ([]entity.NftOwner, error) {
	contracts := make([]entity.NftOwner, 0)

	query := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("owner.keyword", ownerAddr),
		elastic.NewTermQuery("burnedAt", 0),
	)

	contractAgg := elastic.NewTermsAggregation().Field("contract.keyword").Size(10000).
		SubAggregation("tokenId", elastic.NewTermsAggregation().Field("tokenId").Size(10000))

	agg := elastic.NewFilterAggregation().Filter(query)
	if shape == "" || shape == string(entity.ZRC1) {
		agg.SubAggregation("zrc6", elastic.NewFilterAggregation().Filter(elastic.NewTermQuery("zrc6", true)).
			SubAggregation("contract", contractAgg))
	}
	if shape == "" || shape == string(entity.ZRC6) {
		agg.SubAggregation("zrc1", elastic.NewFilterAggregation().Filter(elastic.NewTermQuery("zrc1", true)).
			SubAggregation("contract", contractAgg))
	}

	results, err := search(nftRepo.elastic.Client.
		Search(elastic_search.NftIndex.Get(network)).
		Aggregation("owner", agg).
		Size(0))

	if err != nil {
		return contracts, err
	}

	if ownerAgg, found := results.Aggregations.Filter("owner"); found {
		if zrc1Agg, found := ownerAgg.Aggregations.Terms("zrc1"); found {
			if contractAgg, found := zrc1Agg.Aggregations.Terms("contract"); found {
				for _, contractsBucket := range contractAgg.Buckets {
					contract := entity.NftOwner{Address: contractsBucket.Key.(string), ZRC6: false}

					if tokenIdBucket, found := contractsBucket.Terms("tokenId"); found {
						for _, tokenId := range tokenIdBucket.Buckets {
							contract.TokenIds = append(contract.TokenIds, uint64(tokenId.Key.(float64)))
						}
					}
					contracts = append(contracts, contract)
				}
			}
		}
		if zrc6Agg, found := ownerAgg.Aggregations.Terms("zrc6"); found {
			if contractAgg, found := zrc6Agg.Aggregations.Terms("contract"); found {
				for _, contractsBucket := range contractAgg.Buckets {
					contract := entity.NftOwner{Address: contractsBucket.Key.(string), ZRC6: true}

					if tokenIdBucket, found := contractsBucket.Terms("tokenId"); found {
						for _, tokenId := range tokenIdBucket.Buckets {
							contract.TokenIds = append(contract.TokenIds, uint64(tokenId.Key.(float64)))
						}
					}
					contracts = append(contracts, contract)
				}
			}
		}
	}

	return contracts, nil
}

func (nftRepo nftRepository) GetForContract(network, contractAddr string, size, offset uint64) ([]entity.Nft, int64, error) {
	result, err := search(nftRepo.elastic.Client.
		Search(elastic_search.NftIndex.Get(network)).
		Query(elastic.NewTermQuery("contract.keyword", contractAddr)).
		Size(int(size)).
		Sort("tokenId", true).
		From(int(offset)-1).
		TrackTotalHits(true))

	return nftRepo.findMany(result, err)
}

func (nftRepo nftRepository) GetForContractByTokenId(network, contractAddr string, tokenId uint64) (*entity.Nft, error) {
	query := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("contract.keyword", contractAddr),
		elastic.NewTermQuery("tokenId", tokenId),
	)

	result, err := search(nftRepo.elastic.Client.
		Search(elastic_search.NftIndex.Get(network)).
		Query(query).
		Size(1))

	return nftRepo.findOne(result, err)
}

func (nftRepo nftRepository) GetForContractAttributes(network, contractAddr string) (entity.Attributes, error) {
	query := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("contract.keyword", contractAddr),
		elastic.NewNestedQuery("metadata", elastic.NewTermQuery("metadata.status.keyword", "success")),
	)

	result, err := search(nftRepo.elastic.Client.
		Search(elastic_search.NftIndex.Get(network)).
		Query(query).
		Size(50000))

	nfts, _, err := nftRepo.findMany(result, err)
	if err != nil {
		return nil, err
	}

	attributes := entity.Attributes{}
	for _, nft := range nfts {
		if nft.Metadata.Status == "success" {
			mdAttributes, err := entity.MapToZrc7Attributes(nft.Metadata.Properties["attributes"])
			if err != nil {
				continue
			}
			for _, mdAttribute := range mdAttributes {
				if !attributes.HasAttribute(mdAttribute.TraitType) {
					attributes[mdAttribute.TraitType] = map[string]int64{}
				}
				if !attributes.HasTraitValue(mdAttribute.TraitType, mdAttribute.Value) {
					attributes[mdAttribute.TraitType][fmt.Sprintf("%v", mdAttribute.Value)] = 0
				}
				attributes[mdAttribute.TraitType][fmt.Sprintf("%v", mdAttribute.Value)]++
			}
		}
	}
	return attributes, nil
}

func (nftRepo nftRepository) findOne(results *elastic.SearchResult, err error) (*entity.Nft, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrNftNotFound
	}

	var nft entity.Nft
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &nft)

	return &nft, err
}

func (nftRepo nftRepository) findMany(results *elastic.SearchResult, err error) ([]entity.Nft, int64, error) {
	nfts := make([]entity.Nft, 0)

	if err != nil {
		return nfts, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var nft entity.Nft
		if err := json.Unmarshal(hit.Source, &nft); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, results.TotalHits(), nil
}
