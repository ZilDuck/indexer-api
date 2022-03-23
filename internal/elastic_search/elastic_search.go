package elastic_search

import (
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type Index struct {
	Client *elastic.Client
}

func New(elasticConfig config.ElasticSearchConfig, awsConfig config.AwsConfig) (Index, error) {
	client, err := newClient(elasticConfig, awsConfig)
	if err != nil {
		zap.L().With(zap.Error(err)).Fatal("ElasticCache: Failed to create client")
	}

	return Index{client}, err
}

func newClient(elasticConfig config.ElasticSearchConfig, awsConfig config.AwsConfig) (*elastic.Client, error) {
	zap.S().Infof("Using elastic search instance: %s", elasticConfig.Host)

	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(elasticConfig.Host),
		elastic.SetSniff(elasticConfig.Sniff),
		elastic.SetHealthcheck(elasticConfig.HealthCheck),
	}

	if elasticConfig.Debug {
		opts = append(opts, elastic.SetTraceLog(ElasticLogger{}))
	}

	if elasticConfig.Username != "" {
		opts = append(opts, elastic.SetBasicAuth(
			elasticConfig.Username,
			elasticConfig.Password,
		))
	}

	return elastic.NewClient(opts...)
}
