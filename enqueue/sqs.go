package enqueue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsEnqueue struct {
	QueueUrl *string
	Region   *string
}

func NewSqsEnqueue(queueUrl string, region string) *SqsEnqueue {
	return &SqsEnqueue{
		QueueUrl: aws.String(queueUrl),
		Region:   aws.String(region),
	}
}

func (s *SqsEnqueue) Enqueue(message string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: s.Region,
	}))

	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    s.QueueUrl,
		MessageBody: aws.String(message),
	})
	if err != nil {
		return err
	}

	return err
}
