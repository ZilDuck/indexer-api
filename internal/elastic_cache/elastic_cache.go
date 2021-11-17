package elastic_cache

import (
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic/v7"
	"github.com/sha1sum/aws_signing_client"
	"go.uber.org/zap"
	"strings"
)

type Index struct {
	Client *elastic.Client
}

func New() (Index, error) {
	client, err := newClient()
	if err != nil {
		zap.L().With(zap.Error(err)).Fatal("ElasticCache: Failed to create client")
	}

	return Index{client}, err
}

func newClient() (*elastic.Client, error) {
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(strings.Join(config.Get().ElasticSearch.Hosts, ",")),
		elastic.SetSniff(config.Get().ElasticSearch.Sniff),
		elastic.SetHealthcheck(config.Get().ElasticSearch.HealthCheck),
	}

	if config.Get().ElasticSearch.Debug {
		opts = append(opts, elastic.SetTraceLog(ElasticLogger{}))
	}

	if config.Get().ElasticSearch.Aws {
		creds := credentials.NewStaticCredentials(config.Get().Aws.AccessKey, config.Get().Aws.SecretKey, config.Get().Aws.Token)
		awsClient, err := aws_signing_client.New(v4.NewSigner(creds), nil, "es", config.Get().Aws.Region)
		if err != nil {
			return nil, err
		}

		opts = append(opts, elastic.SetHttpClient(awsClient))
		opts = append(opts, elastic.SetScheme("https"))
		return elastic.NewClient(opts...)
	}

	if config.Get().ElasticSearch.Username != "" {
		opts = append(opts, elastic.SetBasicAuth(
			config.Get().ElasticSearch.Username,
			config.Get().ElasticSearch.Password,
		))
	}

	return elastic.NewClient(opts...)
}
