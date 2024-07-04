package jmap

// Querying for Mailbox/Get
type MailboxGetQuery struct {
	AccountID string `json:"accountId"`
	ids       []string
}

func NewMailboxGetQuery(accountID string) JmapQuery {
	return MailboxGetQuery{AccountID: accountID, ids: nil}
}

func NewMailboxGetMethodCall(query JmapQuery, identifier string) MethodCall {
	return NewMethodCall(MailboxGet, query, identifier)
}

// Handling responses for Mailbox/Get
type Mailbox struct {
	Name        string
	Role        string
	TotalEmails int16
	Id          string
}

func NewMailbox(f map[string]interface{}) Mailbox {
	var role string
	if f["role"] == nil {
		role = ""
	} else {
		role = f["role"].(string)
	}
	return Mailbox{
		Name:        f["name"].(string),
		Role:        role,
		TotalEmails: int16(f["totalEmails"].(float64)),
		Id:          f["id"].(string),
	}
}

type AccountMailboxesInfo struct {
	AccountId string
	State     string
	List      []Mailbox
}

func NewAccountMailboxesInfo(m map[string]interface{}) AccountMailboxesInfo {
	rawMailboxList := m["list"].([]interface{})
	mailboxList := make([]Mailbox, len(rawMailboxList))
	for i, f := range rawMailboxList {
		mailboxList[i] = NewMailbox(f.(map[string]interface{}))
	}
	return AccountMailboxesInfo{
		AccountId: m["accountId"].(string),
		State:     m["state"].(string),
		List:      mailboxList,
	}
}
