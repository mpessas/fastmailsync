package jmap

import "encoding/json"

// Querying for Email/query
type EmailQueryJmapQuery struct {
	AccountID string
	MailboxId string
	Limit     int8
}

func NewEmailQueryJmapQuery(accountID string, mailboxId string, limit int8) JmapQuery {
	return EmailQueryJmapQuery{AccountID: accountID, MailboxId: mailboxId, Limit: limit}
}

type SortOption struct {
	Property    string `json:"property"`
	IsAscending bool   `json:"isAscending"`
}

func (eq EmailQueryJmapQuery) MarshalJSON() ([]byte, error) {
	sortByReceivedAt := SortOption{
		Property: "receivedAt", IsAscending: false,
	}
	payload := map[string]interface{}{
		"accountId": eq.AccountID,
		"filter": map[string]string{
			"inMailbox": eq.MailboxId,
		},
		"sort": []SortOption{
			sortByReceivedAt,
		},
		"position":        0,
		"limit":           eq.Limit,
		"collapseThreads": false,
		"calculateTotal":  true,
	}
	return json.Marshal(payload)
}

func NewEmailQueryMethodCall(query JmapQuery, identifier string) MethodCall {
	return NewMethodCall(EmailQuery, query, identifier)
}

// Handling responses for Email/query
type MailboxInfo struct {
	Position   int8
	QueryState string
	Ids        []string
	total      int
}

func NewMailboxInfo(info map[string]interface{}) MailboxInfo {
	rawIdsList := info["ids"].([]interface{})
	idsList := make([]string, len(rawIdsList))
	for i, id := range rawIdsList {
		idsList[i] = id.(string)
	}
	return MailboxInfo{
		Position:   int8(info["position"].(float64)),
		QueryState: info["queryState"].(string),
		Ids:        idsList,
		total:      int(info["total"].(float64)),
	}
}
