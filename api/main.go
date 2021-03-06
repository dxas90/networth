package main

import (
	_ "github.com/networth-app/networth/dotenv"

	"github.com/gorilla/mux"
	"github.com/networth-app/networth/lib"
)

// NetworthAPI nw api struct
type NetworthAPI struct {
	db     *nwlib.DynamoDBClient
	router *mux.Router
	plaid  *nwlib.PlaidClient
}

var (
	username       = ""
	plaidClientID  = nwlib.GetEnv("PLAID_CLIENT_ID")
	plaidSecret    = nwlib.GetEnv("PLAID_SECRET")
	plaidPublicKey = nwlib.GetEnv("PLAID_PUBLIC_KEY")
	plaidEnv       = nwlib.GetEnv("PLAID_ENV")
	plaidClient    = nwlib.NewPlaidClient(plaidClientID, plaidSecret, plaidPublicKey, plaidEnv)
	kmsClient      = nwlib.NewKMSClient()
)

func main() {
	apiHost := nwlib.GetEnv("API_HOST", ":8000")
	dbClient := nwlib.NewDynamoDBClient()

	api := &NetworthAPI{
		db:     dbClient,
		router: mux.NewRouter(),
		plaid:  plaidClient,
	}
	api.Start(apiHost)
}
