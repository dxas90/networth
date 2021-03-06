package main

import (
	"log"
	"net/http"

	"github.com/networth-app/networth/lib"
)

// AccountResp - account payload
type AccountResp struct {
	Accounts        []nwlib.Account `json:"accounts"`
	InstitutionName string          `json:"institution_name"`
}

func (s *NetworthAPI) handleAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		grouped := make(map[string]AccountResp)
		// var payload []AccountResp

		accounts, err := s.db.GetAccounts(username)

		if err != nil {
			log.Printf("Problem getting accounts: %+v\n", err)
			nwlib.ErrorResp(w, err.Error())
			return
		}

		for _, account := range accounts {
			insID := account.InstitutionID
			insName := account.InstitutionName

			if insID == "" {
				insID = "unknown"
			}

			if insName == "" {
				insName = "Unknown / Manual"
			}

			var existingAccounts []nwlib.Account
			if _, ok := grouped[insID]; ok {
				existingAccounts = grouped[insID].Accounts
			}

			grouped[insID] = AccountResp{
				Accounts:        append(existingAccounts, account),
				InstitutionName: insName,
			}
		}

		// TODO: sort by Institution Name
		// var keys []string
		// for insID := range grouped {
		// 	keys = append(keys, grouped[insID].InstitutionName)
		// }
		// sort.Strings(keys)

		// for key, val := range keys {
		// 	fmt.Println(key, val)
		// }

		nwlib.SuccessResp(w, grouped)
	}
}
