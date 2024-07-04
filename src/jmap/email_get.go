package jmap

import "encoding/json"

// Querying for Email/get
type EmailGetJmapQuery struct {
	AccountID string
	ResultOf  string
}

func NewEmailGetJmapQuery(accountID string, resultOf string) JmapQuery {
	return EmailGetJmapQuery{AccountID: accountID, ResultOf: resultOf}
}

func (q EmailGetJmapQuery) MarshalJSON() ([]byte, error) {
	payload := map[string]interface{}{
		"accountId": q.AccountID,
		"#ids": map[string]string{
			"name":     "Email/query",
			"path":     "/ids",
			"resultOf": q.ResultOf,
		},
	}
	return json.Marshal(payload)
}

func NewEmailGetMethodCall(query JmapQuery, identifier string) MethodCall {
	return NewMethodCall(EmailGet, query, identifier)
}

// Handling responses for Email/get
type Email struct {
	Id        string
	MessageId string
	BlobId    string
}

func NewEmail(payload map[string]interface{}) Email {
	return Email{
		Id:        payload["id"].(string),
		MessageId: payload["messageId"].([]interface{})[0].(string),
		BlobId:    payload["blobId"].(string),
	}
}

type EmailList struct {
	State string
	List  []Email
}

func NewEmailList(payload map[string]interface{}) EmailList {
	rawEmailsList := payload["list"].([]interface{})
	emailsList := make([]Email, len(rawEmailsList))
	for i, e := range rawEmailsList {
		emailsList[i] = NewEmail(e.(map[string]interface{}))
	}
	return EmailList{
		State: payload["state"].(string),
		List:  emailsList,
	}
}
