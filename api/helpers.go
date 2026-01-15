package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetSQSMessage(ctx context.Context, queueURL string) (*sqs.Message, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")}, nil)
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.New(sess)

	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(20),
	}

	output, err := sqsClient.ReceiveMessageWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if len(output.Messages) == 0 {
		return nil, fmt.Errorf("no messages in queue")
	}

	return output.Messages[0], nil
}

func ProcessSQSMessage(ctx context.Context, queueURL string, processor func(context.Context, *sqs.Message) error) error {
	for {
		msg, err := GetSQSMessage(ctx, queueURL)
		if err != nil {
			return err
		}

		err = processor(ctx, msg)
		if err != nil {
			log.Printf("error processing message: %v", err)
		}

		// delete the message from the queue
		deleteMessageInput := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: msg.ReceiptHandle,
		}
		_, err = sqsClient.DeleteMessageWithContext(ctx, deleteMessageInput)
		if err != nil {
			return err
		}
	}
}

func ProcessSQSMessagesWithContext(ctx context.Context, queueURL string, processor func(context.Context, *sqs.Message) error, maxMessages int) error {
	sqsClient := sqs.New(sess)

	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(int64(maxMessages)),
		WaitTimeSeconds:     aws.Int64(20),
	}

	output, err := sqsClient.ReceiveMessageWithContext(ctx, input)
	if err != nil {
		return err
	}

	for _, msg := range output.Messages {
		err := processor(ctx, msg)
		if err != nil {
			log.Printf("error processing message: %v", err)
		}

		// delete the message from the queue
		deleteMessageInput := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: msg.ReceiptHandle,
		}
		_, err = sqsClient.DeleteMessageWithContext(ctx, deleteMessageInput)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetProcessTime() time.Duration {
	start := time.Now()
	return time.Since(start)
}