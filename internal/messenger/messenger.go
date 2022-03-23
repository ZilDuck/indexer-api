package messenger

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

type MessageService interface {
	SendMessage(network string, queue Queue, body []byte) error
}

type Messenger struct {
	sqsClient *sqs.SQS
}

type Queue string

const (
	MetadataRefresh Queue = "metadata_refresh"
)

func (q *Queue) Get(network string) string {
	return fmt.Sprintf("%s_%s", *q, network)
}

func NewMessenger(sqsClient *sqs.SQS) MessageService {
	return &Messenger{sqsClient}
}

func (m Messenger) SendMessage(network string, queue Queue, body []byte) error {
	zap.L().With(zap.String("queue", string(queue)), zap.String("body", string(body))).Info("Send Message")
	queueUrl, err := m.getQueueUrl(network, queue)
	if err != nil {
		return err
	}

	_, err = m.sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    queueUrl,
		MessageBody: aws.String(string(body)),
	})

	return err
}

func (m Messenger) getQueueUrl(network string, queue Queue) (*string, error) {
	queueName := queue.Get(network)
	result, err := m.sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &queueName})
	if err != nil {
		zap.L().With(zap.Error(err), zap.String("queue", queueName)).Error("Failed to get queue url")
		return nil, err
	}

	return result.QueueUrl, nil
}
