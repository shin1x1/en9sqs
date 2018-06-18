package enqueue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsEnqueue struct {
	QueueUrl *string
	sqs      *sqs.SQS
}

func NewSqsEnqueue(queueUrl string, region string) *SqsEnqueue {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return &SqsEnqueue{
		QueueUrl: aws.String(queueUrl),
		sqs:      sqs.New(sess),
	}
}

func (s *SqsEnqueue) Enqueue(message string) error {
	_, err := s.sqs.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    s.QueueUrl,
		MessageBody: aws.String(message),
	})
	if err != nil {
		return err
	}

	return err
}
