package jmap

import "encoding/json"

const ApiUrl = "https://jmap.fastmail.com/api/"

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

func NewPayload(methodCalls []MethodCall) *JmapPayload {
	using := []string{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"}
	return newJmapPayload(using, methodCalls)
}
