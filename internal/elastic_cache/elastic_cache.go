package elastic_cache

import (
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic/v7"
	"github.com/sha1sum/aws_signing_client"
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
		//elastic.SetURL(elasticConfig.Host),
		elastic.SetURL("https://search-zilkroad-index-5yazp27hjzc6jrl4cuqw5t6nau.eu-west-1.es.amazonaws.com"),
		elastic.SetSniff(elasticConfig.Sniff),
		elastic.SetHealthcheck(elasticConfig.HealthCheck),
	}

	if elasticConfig.Debug {
		opts = append(opts, elastic.SetTraceLog(ElasticLogger{}))
	}

	if elasticConfig.Aws {
		creds := credentials.NewStaticCredentials(awsConfig.AccessKey, awsConfig.SecretKey, awsConfig.Token)
		awsClient, err := aws_signing_client.New(v4.NewSigner(creds), nil, "es", awsConfig.Region)
		if err != nil {
			return nil, err
		}

		opts = append(opts, elastic.SetHttpClient(awsClient))
		opts = append(opts, elastic.SetScheme("https"))
		return elastic.NewClient(opts...)
	}

	if elasticConfig.Username != "" {
		opts = append(opts, elastic.SetBasicAuth(
			elasticConfig.Username,
			elasticConfig.Password,
		))
	}

	return elastic.NewClient(opts...)
}
