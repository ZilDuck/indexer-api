package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dantudor/zilkroad-txapi/internal/elastic_cache"
	"github.com/dantudor/zilkroad-txapi/internal/entity"
	"github.com/gin-contrib/cache/persistence"
	"github.com/olivere/elastic/v7"
)

type TransactionRepository interface {
	GetBestBlock() (uint64, error)
}

type transactionRepository struct {
	elastic *elastic_cache.Index
	cache   persistence.CacheStore
}

var (
	ErrTxNotFound = errors.New("tx not found")
)

func NewTransactionRepository(elastic *elastic_cache.Index, cache persistence.CacheStore) TransactionRepository {
	return transactionRepository{elastic, cache}
}

func (txRepo transactionRepository) GetBestBlock() (uint64, error) {
	result, err := txRepo.elastic.Client.
		Search(elastic_cache.TransactionIndex.Get()).
		Size(1).
		Sort("BlockNum", false).
		Do(context.Background())

	tx, err := txRepo.findOne(result, err)
	if err != nil {
		return 0, err
	}

	return tx.BlockNum, nil
}

func (txRepo transactionRepository) findOne(results *elastic.SearchResult, err error) (entity.Transaction, error) {
	if err != nil {
		return entity.Transaction{}, err
	}

	if len(results.Hits.Hits) == 0 {
		return entity.Transaction{}, ErrTxNotFound
	}

	var tx entity.Transaction
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &tx)

	return tx, err
}
