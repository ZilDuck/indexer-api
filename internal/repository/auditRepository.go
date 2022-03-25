package repository

import (
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/olivere/elastic/v7"
	"time"
)

type AuditRepository interface {
	GetByDateAndApiKey(t time.Time, apiKey string, size, offset uint64) ([]entity.Audit, int64, error)
	CountByDateAndApiKey(t time.Time, apiKey string) (int64, error)
}

type auditRepository struct {
	elastic elastic_search.Index
}

var (
	ErrAuditNotFound = errors.New("audit not found")
)

func NewAuditRepository(elastic elastic_search.Index) AuditRepository {
	return auditRepository{elastic: elastic}
}

func (auditRepo auditRepository) GetByDateAndApiKey(t time.Time, apiKey string, size, offset uint64) ([]entity.Audit, int64, error) {
	results, err := search(auditRepo.elastic.Client.
		Search(elastic_search.AuditIndex.GetByDate(t.Format("2006.01.02"))).
		Query(elastic.NewTermQuery("apiKey", apiKey)).
		Sort("@timestamp", false).
		TrackTotalHits(true).
		Size(int(size)).
		From(int(offset)-1))

	if errors.Is(err, ErrNoSuchIndex) {
		return []entity.Audit{}, 0, nil
	}

	return auditRepo.findMany(results, err)
}

func (auditRepo auditRepository) CountByDateAndApiKey(t time.Time, apiKey string) (int64, error) {
	return count(auditRepo.elastic.Client.
		Count(elastic_search.AuditIndex.GetByDate(t.Format("2006.01.02"))).
		Query(elastic.NewTermQuery("apiKey", apiKey)))
}


func (auditRepo auditRepository) findOne(results *elastic.SearchResult, err error) (*entity.Audit, error) {
	if err != nil {
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, ErrAuditNotFound
	}

	var audit entity.Audit
	hit := results.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &audit)

	return &audit, err
}

func (auditRepo auditRepository) findMany(results *elastic.SearchResult, err error) ([]entity.Audit, int64, error) {
	audits := make([]entity.Audit, 0)

	if err != nil {
		return audits, 0, err
	}

	for _, hit := range results.Hits.Hits {
		var audit entity.Audit
		if err := json.Unmarshal(hit.Source, &audit); err == nil {
			audits = append(audits, audit)
		}
	}

	return audits, results.TotalHits(), nil
}
