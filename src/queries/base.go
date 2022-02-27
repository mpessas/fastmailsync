package queries

import "encoding/json"

type JmapQuery struct {
	AccountID string `json:"accountId"`
	Limit     int8   `json:"limit"`
}

func NewJmapQuery(accountID string, limit int8) JmapQuery {
	return JmapQuery{AccountID: accountID, Limit: limit}
}

type MethodCall struct {
	QueryType string
	Query     JmapQuery
}

func (mc MethodCall) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{mc.QueryType, mc.Query, "a"})
}

type JmapPayload struct {
	Using       []string     `json:"using"`
	MethodCalls []MethodCall `json:"methodCalls"`
}

func newJmapPayload(using []string, methodCalls []MethodCall) *JmapPayload {
	p := JmapPayload{Using: using, MethodCalls: methodCalls}
	return &p
}

func (p *JmapPayload) ToJson() ([]byte, error) {
	return json.Marshal(p)
}
