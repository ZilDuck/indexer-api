package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
)

type NftActionRepository interface {
	GetByContractAndTokenId(network, contractAddr string, tokenId uint64, actionTypes[]entity.ActionType, size, offset int) ([]entity.NftAction, int64, error)
}

type nftActionRepository struct {
	elastic elastic_search.Index
}

var (
	ErrNftActionNotFound = errors.New("no nft actions found")
)

func NewActionRepository(elastic elastic_search.Index) NftActionRepository {
	return nftActionRepository{elastic: elastic}
}

func (actionRepo nftActionRepository) GetByContractAndTokenId(network, contractAddr string, tokenId uint64, actionTypes[]entity.ActionType, size, offset int) ([]entity.NftAction, int64, error) {
	queries := []elastic.Query{
		elastic.NewTermQuery("contract.keyword", contractAddr),
		elastic.NewTermQuery("tokenId", tokenId),
	}

	if len(actionTypes) != 0 {
		queries = append(queries, elastic.NewTermsQuery("action", valuesFromActionTypes(actionTypes)...))
	}

	result, err := search(actionRepo.elastic.Client.
		Search(elastic_search.NftActionIndex.Get(network)).
		Query(elastic.NewBoolQuery().Must(queries...)).
		TrackTotalHits(true).
		Size(size).
		From(offset).
		Sort("blockNum", false))

	return actionRepo.findMany(result, err)
}

func (actionRepo nftActionRepository) findOne(results *elastic.SearchResult, err error) (*entity.NftAction, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrNftActionNotFound
	}

	var action entity.NftAction
	err = json.Unmarshal(results.Hits.Hits[0].Source, &action)

	return &action, err
}

func (actionRepo nftActionRepository) findMany(results *elastic.SearchResult, err error) ([]entity.NftAction, int64, error) {
	actions := make([]entity.NftAction, 0)

	if err != nil {
		return actions, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var action entity.NftAction
		if err := json.Unmarshal(hit.Source, &action); err == nil {
			actions = append(actions, action)
		}
	}

	return actions, results.TotalHits(), nil
}
