package main

import (
	"log"
	"time"
)

// sync last 12 months
func syncTransactions(username string, token string) error {
	endDate := time.Now().UTC()
	endDateStr := endDate.Format("2006-01-02")
	startDate := endDate.AddDate(0, -12, 0)
	startDateStr := startDate.Format("2006-01-02")

	trans, err := plaidClient.GetTransactions(token, startDateStr, endDateStr)

	if err != nil {
		log.Printf("syncTransactions() Problem getting trans ", err)
		return err
	}

	for _, tran := range trans.Transactions {
		if err := db.SetTransaction(username, tran); err != nil {
			log.Printf("Problem saving this transaction to db: %+v", tran)
		}
	}

	return nil
}
