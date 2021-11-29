package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_cache"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type ContactRepository interface {
	GetContracts(network string) ([]*entity.Contract, int64, error)
	GetContract(network, contractAddr string) (*entity.Contract, error)
}

type contactRepository struct {
	elastic elastic_cache.Index
}

var (
	ErrContractNotFound = errors.New("contract not found")
)

func NewContractRepository(elastic elastic_cache.Index) ContactRepository {
	return contactRepository{elastic: elastic}
}

func (contractRepo contactRepository) GetContracts(network string) ([]*entity.Contract, int64, error) {
	result, err := search(contractRepo.elastic.Client.
		Search(elastic_cache.ContractIndex.Get(network)).
		Query(elastic.NewTermQuery("zrc6", true)).
		Sort("blockNum", false).
		Size(100))

	return contractRepo.findMany(result, err)
}

func (contractRepo contactRepository) GetContract(network, contractAddr string) (*entity.Contract, error) {
	result, err := search(contractRepo.elastic.Client.
		Search(elastic_cache.ContractIndex.Get(network)).
		Query(elastic.NewTermQuery("address.keyword", contractAddr)).
		Size(1))

	return contractRepo.findOne(result, err)
}

func (contractRepo contactRepository) findOne(results *elastic.SearchResult, err error) (*entity.Contract, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrContractNotFound
	}

	var contract entity.Contract
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &contract)

	return &contract, err
}

func (contractRepo contactRepository) findMany(results *elastic.SearchResult, err error) ([]*entity.Contract, int64, error) {
	contracts := make([]*entity.Contract, 0)

	if err != nil {
		return contracts, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var contract entity.Contract
		if err := json.Unmarshal(hit.Source, &contract); err == nil {
			contracts = append(contracts, &contract)
		}
	}

	return contracts, results.TotalHits(), nil
}
