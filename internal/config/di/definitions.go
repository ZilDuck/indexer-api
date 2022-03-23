package di

import (
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/ZilDuck/indexer-api/internal/messenger"
	"github.com/ZilDuck/indexer-api/internal/metadata"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sarulabs/dingo/v4"
	"go.uber.org/zap"
)

var Definitions = []dingo.Def{
	{
		Name: "elastic",
		Build: func() (elastic_search.Index, error) {
			elastic, err := elastic_search.New(config.Get().ElasticSearch, config.Get().Aws)
			if err != nil {
				zap.L().With(zap.Error(err)).Fatal("Failed to start ES")
			}

			return elastic, nil
		},
	},
	{
		Name: "sqs",
		Build: func() (*sqs.SQS, error) {
			sess := session.Must(session.NewSession(&aws.Config{
				Credentials: credentials.NewStaticCredentials(config.Get().Aws.AccessKey, config.Get().Aws.SecretKey, ""),
			}))

			return sqs.New(sess), nil
		},
	},
	{
		Name: "messenger",
		Build: func(sqs *sqs.SQS) (messenger.MessageService, error) {
			return messenger.NewMessenger(sqs), nil
		},
	},
	{
		Name: "contract.repository",
		Build: func(elastic elastic_search.Index) (repository.ContactRepository, error) {
			return repository.NewContractRepository(elastic), nil
		},
	},
	{
		Name: "nft.repository",
		Build: func(elastic elastic_search.Index) (repository.NftRepository, error) {
			return repository.NewNftRepository(elastic), nil
		},
	},
	{
		Name: "action.repository",
		Build: func(elastic elastic_search.Index) (repository.NftActionRepository, error) {
			return repository.NewActionRepository(elastic), nil
		},
	},
	{
		Name: "metadata.service",
		Build: func() (metadata.Service, error) {
			retryClient := retryablehttp.NewClient()
			retryClient.RetryMax = 3

			return metadata.NewMetadataService(retryClient), nil
		},
	},
}
