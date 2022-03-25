package repository

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"strings"
	"time"
)

var (
	ErrNoSuchIndex = errors.New("no such index")
)

func search(searchService *elastic.SearchService) (*elastic.SearchResult, error) {
	result, err := searchService.Do(context.Background())
	if err != nil {
		if err.Error() == "elastic: Error 429 (Too Many Requests)" {
			zap.L().Warn("Elastic: 429 (Too Many Requests)")
			time.Sleep(5 * time.Second)
			return search(searchService)
		}
		if strings.Contains(err.Error(), "no such index") {
			return nil, ErrNoSuchIndex
		}
	}


	return result, err
}

func count(countService *elastic.CountService) (int64, error) {
	result, err := countService.Do(context.Background())
	if err != nil {
		if err.Error() == "elastic: Error 429 (Too Many Requests)" {
			zap.L().Warn("Elastic: 429 (Too Many Requests)")
			time.Sleep(5 * time.Second)
			return count(countService)
		}
		if strings.Contains(err.Error(), "no such index") {
			return 0, ErrNoSuchIndex
		}
	}

	return result, err
}
