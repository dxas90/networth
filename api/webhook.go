package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/networth-app/networth/lib"
)

var (
	snsARN = nwlib.GetEnv("SNS_TOPIC_ARN")
	// Plaid webhook ips: https://support.plaid.com/customer/en/portal/articles/2546264-webhook-overview
	plaidIPs = []string{"52.21.26.131", "52.21.47.157", "52.41.247.19", "52.88.82.239"}
)

func (s *NetworthAPI) handleWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var webhook nwlib.Webhook
		ips := r.Header.Get("X-Forwarded-For")

		validIP := false
		for ip := range ips {
			for plaidIP := range plaidIPs {
				if ip == plaidIP {
					validIP = true
					break
				}
			}
		}

		if !validIP {
			log.Printf("Invalid webhook, IP does not match whitelisted IP: %+v", ips)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&webhook); err != nil {
			nwlib.ErrorResp(w, err.Error())
			return
		}

		// look up username by itemID:
		username, err := s.db.GetUsernameByItemID(webhook.ItemID)

		if err != nil || username == "" {
			log.Printf("Cannot find username based on itemID: %s", webhook.ItemID)
			nwlib.ErrorResp(w, err.Error())
			return
		}

		webhook.Username = username

		log.Printf("New webhook, username: %s, type: %s, code: %s, item: %s\n", webhook.Username, webhook.WebhookType, webhook.WebhookCode, webhook.ItemID)
		if err := s.db.SetWebhook(webhook); err != nil {
			log.Printf("Problem saving webhook to db: %+v\n", err)
			nwlib.ErrorResp(w, err.Error())
			return
		}

		nwlib.SuccessResp(w, webhook)
	}
}
