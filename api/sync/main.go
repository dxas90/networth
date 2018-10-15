package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/networth-app/networth/api/lib"
)

var (
	plaidClientID  = nwlib.GetEnv("PLAID_CLIENT_ID")
	plaidSecret    = nwlib.GetEnv("PLAID_SECRET")
	plaidPublicKey = nwlib.GetEnv("PLAID_PUBLIC_KEY")
	plaidEnv       = nwlib.GetEnv("PLAID_ENV")
	plaidClient    = nwlib.NewPlaidClient(plaidClientID, plaidSecret, plaidPublicKey, plaidEnv)
	kms            = nwlib.NewKMSClient()
	db             = nwlib.NewDynamoDBClient()
	snsARN         = nwlib.GetEnv("SNS_TOPIC_ARN")
	slackURL       = nwlib.GetEnv("SLACK_WEBHOOK_URL")
)

func handleScheduledEvent(ctx context.Context, e events.CloudWatchEvent) {
	if e.Source != "aws.events" {
		fmt.Printf("Invalid source: %s\n", e.Source)
		return
	}

	// TODO: get all active username
	if err := nwlib.SyncNetworth(db, "c1fa7e12-529e-4b63-8c64-855ba23690ff"); err != nil {
		fmt.Println("Problem syncing networth ", err)
	}

	fmt.Println("Sync done.")
}

func main() {
	lambda.Start(handleScheduledEvent)
}
