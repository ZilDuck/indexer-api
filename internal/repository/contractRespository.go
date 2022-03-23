package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type ContactRepository interface {
	GetAll(network string) ([]entity.Contract, int64, error)
	GetContract(network, contractAddr string) (*entity.Contract, error)
	GetContracts(network string, contractAddrs ...string) ([]entity.Contract, error)
	GetAllOwnedBy(network string, ownerAddr string) ([]entity.Contract, error)
	GetAllAddressesOwnedBy(network string, ownerAddr string) ([]string, error)
}

type contactRepository struct {
	elastic elastic_search.Index
}

var (
	ErrContractNotFound = errors.New("contract not found")
)

func NewContractRepository(elastic elastic_search.Index) ContactRepository {
	return contactRepository{elastic: elastic}
}

func (contractRepo contactRepository) GetAll(network string) ([]entity.Contract, int64, error) {
	result, err := search(contractRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(elastic.NewTermQuery("zrc6", true)).
		Sort("blockNum", false).
		Size(100))

	return contractRepo.findMany(result, err)
}

func (contractRepo contactRepository) GetContract(network, contractAddr string) (*entity.Contract, error) {
	result, err := search(contractRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(elastic.NewTermQuery("address.keyword", contractAddr)).
		Size(1))

	return contractRepo.findOne(result, err)
}

func (contractRepo contactRepository) GetContracts(network string, contractAddrs ...string) ([]entity.Contract, error) {
	values := make([]interface{}, len(contractAddrs))
	for i, v := range contractAddrs {
		values[i] = v
	}

	result, err := search(contractRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(elastic.NewTermsQuery("address.keyword", values...)).
		Sort("blockNum", false).
		Size(100))

	contracts, _, err := contractRepo.findMany(result, err)
	return contracts, err
}

func (contractRepo contactRepository) GetAllOwnedBy(network string, ownerAddr string) ([]entity.Contract, error) {
	query := elastic.NewNestedQuery("state", elastic.NewBoolQuery().Should(
		elastic.NewBoolQuery().Must(
			elastic.NewTermQuery("state.key.keyword", "owner"),
			elastic.NewTermQuery("state.value.keyword", ownerAddr),
		),
		elastic.NewBoolQuery().Must(
			elastic.NewTermQuery("state.key.keyword", "contractOwner"),
			elastic.NewTermQuery("state.value.keyword", ownerAddr),
		),
	).MinimumShouldMatch("1"))

	result, err := search(contractRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(query).
		Size(1000))
	if err != nil {
		return nil, err
	}

	contracts, _, err := contractRepo.findMany(result, err)

	return contracts, err
}

func (contractRepo contactRepository) GetAllAddressesOwnedBy(network string, ownerAddr string) ([]string, error) {
	query := elastic.NewNestedQuery("state", elastic.NewBoolQuery().Should(
		elastic.NewBoolQuery().Must(
			elastic.NewTermQuery("state.key.keyword", "owner"),
			elastic.NewTermQuery("state.value.keyword", ownerAddr),
		),
		elastic.NewBoolQuery().Must(
			elastic.NewTermQuery("state.key.keyword", "contractOwner"),
			elastic.NewTermQuery("state.value.keyword", ownerAddr),
		),
	).MinimumShouldMatch("1"))

	result, err := search(contractRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(query).
		Aggregation("contractAddr", elastic.NewTermsAggregation().Field("address.keyword")).
		Size(0))
	if err != nil {
		return nil, err
	}

	contracts := make([]string, 0)
	if buckets, found := result.Aggregations.Terms("contractAddr"); found {
		for _, contractAddr := range buckets.Buckets {
			contracts = append(contracts, contractAddr.Key.(string))
		}
	}

	return contracts, nil
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

func (contractRepo contactRepository) findMany(results *elastic.SearchResult, err error) ([]entity.Contract, int64, error) {
	contracts := make([]entity.Contract, 0)

	if err != nil {
		return contracts, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var contract entity.Contract
		if err := json.Unmarshal(hit.Source, &contract); err == nil {
			contracts = append(contracts, contract)
		}
	}

	return contracts, results.TotalHits(), nil
}
