package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type ContactMetadataRepository interface {
	GetMetadata(network, contractAddr string) (*entity.ContractMetadata, error)
}

type contactMetadataRepository struct {
	elastic elastic_search.Index
}

var (
	ErrContractMetadataNotFound = errors.New("contract metadata not found")
)

func NewContactMetadataRepository(elastic elastic_search.Index) ContactMetadataRepository {
	return contactMetadataRepository{elastic: elastic}
}

func (contactMetadataRepo contactMetadataRepository) GetMetadata(network, contractAddr string) (*entity.ContractMetadata, error) {
	result, err := search(contactMetadataRepo.elastic.Client.
		Search(elastic_search.ContractMetadataIndex.Get(network)).
		Query(elastic.NewTermQuery("contract.keyword", contractAddr)).
		Size(1))

	return contactMetadataRepo.findOne(result, err)
}

func (contactMetadataRepo contactMetadataRepository) findOne(results *elastic.SearchResult, err error) (*entity.ContractMetadata, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrContractMetadataNotFound
	}

	var md entity.ContractMetadata
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &md)

	return &md, err
}
