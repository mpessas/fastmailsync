package main

import (
	"log"
	"os"

	"github.com/mpessas/fastmailsync/queries"
)

func getEmails(accountId string, auth BasicAuthInfo) {
	log.Println("Fetching emails for account ID:", accountId)
	query := queries.NewJmapQuery(accountId, 10)
	mc := queries.NewEmailQueryMethodCall(query)
	payload := queries.NewEmailQueryPayload([]queries.MethodCall{mc})
	post(auth, payload)
}

func main() {
	accountId := os.Getenv("FASTMAIL_ACCOUNT_ID")
	username := os.Getenv("FASTMAIL_USERNAME")
	password := os.Getenv("FASTMAIL_PASSWORD")
	auth := BasicAuthInfo{username, password}
	getEmails(accountId, auth)
}
