package jmap

import (
	"encoding/json"
)

type Response struct {
	SessionState    string           `json:"sessionState"`
	MethodResponses []MethodResponse `json:methodResponses`
}

type MethodResponse struct {
	MethodType MethodCallType
	Result     map[string]interface{}
	QueryKey   string
}

func (mr *MethodResponse) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[0], &mr.MethodType); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[1], &mr.Result); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[2], &mr.QueryKey); err != nil {
		return err
	}
	return nil
}
