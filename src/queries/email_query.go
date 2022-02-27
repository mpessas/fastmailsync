package queries

func NewEmailQueryMethodCall(query JmapQuery) MethodCall {
	return MethodCall{"Email/query", query}
}

func NewEmailQueryPayload(methodCalls []MethodCall) *JmapPayload {
	using := []string{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"}
	return newJmapPayload(using, methodCalls)
}
