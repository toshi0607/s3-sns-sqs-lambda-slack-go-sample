package main

import (
	"fmt"
	"os"

	"github.com/toshi0607/s3-sns-sqs-lambda-slack-go-sample/handlers/notifier/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var client *slack.Client

func main() {
	lambda.Start(handler)
}

func init() {
	client = slack.NewClient(
		slack.Config{
			URL:       os.Getenv("WEBHOOK_URL"),
			Channel:   os.Getenv("CHANNEL"),
			Username:  os.Getenv("USERNAME"),
			IconEmoji: os.Getenv("ICON"),
		},
	)
}

func handler(snsEvent events.SNSEvent) error {
	record := snsEvent.Records[0]
	snsRecord := snsEvent.Records[0].SNS
	fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

	if err := client.PostMessage(snsRecord.Message); err != nil {
		return err
	}

	return nil
}
