package main

import (
	"context"

	"github.com/networth-app/networth/lib"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	slackWebhookURL = nwlib.GetEnv("SLACK_WEBHOOK_URL")
	slackChannel    = nwlib.GetEnv("SLACK_CHANNEL", "sns")
)

func handleSNSNotification(ctx context.Context, snsEvent events.SNSEvent) {
	messages := ""
	for _, record := range snsEvent.Records {
		messages += record.SNS.Message + "\n"
	}

	nwlib.PublishSlack(slackWebhookURL, messages, slackChannel)
}

func main() {
	lambda.Start(handleSNSNotification)
}
