package main

import (
	"encoding/json"
	"log"

	"github.com/mpessas/fastmailsync/jmap"
)

func fetchEmailsForMailbox(maildir string, accountId string, auth BasicAuthInfo, mailbox jmap.Mailbox) {
	emailQuery := jmap.NewEmailQueryJmapQuery(accountId, mailbox.Id, 10)
	mcEmailQuery := jmap.NewEmailQueryMethodCall(emailQuery, "a")
	emailGet := jmap.NewEmailGetJmapQuery(accountId, "a")
	mcEmailGet := jmap.NewEmailGetMethodCall(emailGet, "b")
	payload := jmap.NewPayload([]jmap.MethodCall{mcEmailQuery, mcEmailGet})
	body := post(jmap.ApiUrl, auth, payload, false)

	var response jmap.Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalln("Could not parse JSON response:", err)
	}
	mailboxInfo := jmap.NewMailboxInfo(response.MethodResponses[0].Result)
	emails := jmap.NewEmailList(response.MethodResponses[1].Result)
	log.Println(mailboxInfo)
	log.Println(emails)
	for _, email := range emails.List {
		info := MessageInfo{accountId, email.Id, email.BlobId}
		download(maildir, info, auth)
	}
}

func getMailboxesInfo(accountId string, auth BasicAuthInfo) jmap.AccountMailboxesInfo {
	query := jmap.NewMailboxGetQuery(accountId)
	mc := jmap.NewMailboxGetMethodCall(query, "a")
	payload := jmap.NewPayload([]jmap.MethodCall{mc})
	body := post(jmap.ApiUrl, auth, payload, false)

	var response jmap.Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalln("Could not parse JSON response:", err)
	}
	return jmap.NewAccountMailboxesInfo(response.MethodResponses[0].Result)
}

func initialSync(conf Configuration, auth BasicAuthInfo) {
	log.Println("Fetching mailboxes...")
	mailboxInfo := getMailboxesInfo(conf.AccountId, auth)
	for _, mailbox := range mailboxInfo.List {
		log.Printf("Fetching emails for %s ...", mailbox.Name)
		fetchEmailsForMailbox(conf.MailDir, conf.AccountId, auth, mailbox)
	}
}

func main() {
	conf := ReadConfiguration()
	auth := BasicAuthInfo{conf.Username, conf.Password}
	initialSync(conf, auth)
}
