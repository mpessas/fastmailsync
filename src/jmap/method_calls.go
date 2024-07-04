package jmap

import "encoding/json"

type JmapQuery interface{}

type MethodCall struct {
	QueryType  MethodCallType
	Query      JmapQuery
	Identifier string
}

func NewMethodCall(mcType MethodCallType, query JmapQuery, identifier string) MethodCall {
	return MethodCall{mcType, query, identifier}
}

func (mc MethodCall) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{mc.QueryType, mc.Query, mc.Identifier})
}
