package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type ContactStateRepository interface {
	GetState(network, contractAddr string) (*entity.ContractState, error)
	GetAllOwnedBy(network string, ownerAddr string) ([]entity.ContractState, error)
	GetAllAddressesOwnedBy(network string, ownerAddr string) ([]string, error)
}

type contactStateRepository struct {
	elastic elastic_search.Index
}

var (
	ErrContractStateNotFound = errors.New("contract not found")
)

func NewContactStateRepository(elastic elastic_search.Index) ContactStateRepository {
	return contactStateRepository{elastic: elastic}
}

func (contractStateRepo contactStateRepository) GetState(network, contractAddr string) (*entity.ContractState, error) {
	result, err := search(contractStateRepo.elastic.Client.
		Search(elastic_search.ContractStateIndex.Get(network)).
		Query(elastic.NewTermQuery("address.keyword", contractAddr)).
		Size(1))

	return contractStateRepo.findOne(result, err)
}

func (contractStateRepo contactStateRepository) GetAllOwnedBy(network string, ownerAddr string) ([]entity.ContractState, error) {
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

	result, err := search(contractStateRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(query).
		TrackTotalHits(true).
		Size(1000))
	if err != nil {
		return nil, err
	}

	contracts, _, err := contractStateRepo.findMany(result, err)

	return contracts, err
}

func (contractStateRepo contactStateRepository) GetAllAddressesOwnedBy(network string, ownerAddr string) ([]string, error) {
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

	result, err := search(contractStateRepo.elastic.Client.
		Search(elastic_search.ContractIndex.Get(network)).
		Query(query).
		Aggregation("contractAddr", elastic.NewTermsAggregation().Field("address.keyword")).
		TrackTotalHits(true).
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

func (contractStateepo contactStateRepository) findOne(results *elastic.SearchResult, err error) (*entity.ContractState, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrContractStateNotFound
	}

	var state entity.ContractState
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &state)

	return &state, err
}

func (contractStateRepo contactStateRepository) findMany(results *elastic.SearchResult, err error) ([]entity.ContractState, int64, error) {
	states := make([]entity.ContractState, 0)

	if err != nil {
		return states, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var state entity.ContractState
		if err := json.Unmarshal(hit.Source, &state); err == nil {
			states = append(states, state)
		}
	}

	return states, results.TotalHits(), nil
}
