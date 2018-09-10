package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/networth-app/networth/api/lib"
)

var kms = nwlib.NewKMSClient()

func decryptTokens() tokens []string {

}

// append token from single institution to the "all" institution sort key
func appendToken(token) error {

}

func tokens(record events.DynamoDBEventRecord) (username string, tokens []string) {
	var email string
	var decryptedAccessTokens []string

	for name, value := range record.Change.NewImage {
		if name == "email" {
			email = value.String()
		}

		if value.DataType() == events.DataTypeMap {
			val := value.Map()
			tokens := val["access_tokens"].List()

			for _, token := range tokens {
				decryptedToken := kms.Decrypt(token.String())
				decryptedAccessTokens = append(decryptedAccessTokens, decryptedToken)
			}
		}
	}

	return email, decryptedAccessTokens
}
