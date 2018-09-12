package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func syncAccounts(username string, institutionID string, token string) error {
	accounts, err := plaidClient.GetAccounts(token)

	if err != nil {
		log.Println("syncAccounts() Problem getting accounts ", err)
		return err
	}

	for _, account := range accounts.Accounts {
		db.SetAccount(username, institutionID, &account)
	}

	return nil
}

// append token from single institution to the "all" institution sort key
func appendAccount(username string, accounts []events.DynamoDBAttributeValue) error {
	for _, account := range accounts {
		fmt.Println(account)
		// account := &plaid.Account{
		// 	ItemID:          tokenMap["item_id"].String(),
		// 	AccessToken:     tokenMap["access_token"].String(),
		// 	InstitutionID:   tokenMap["institution_id"].String(),
		// 	InstitutionName: tokenMap["institution_name"].String(),
		// }

		// db.SetAccount(username, nwlib.DefaultSortValue, account)
		// db.SetAccount(username, institutionID, &account)
	}

	return nil
}
