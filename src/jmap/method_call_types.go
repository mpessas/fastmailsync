package jmap

type MethodCallType string

const (
	MailboxGet MethodCallType = "Mailbox/get"
	EmailQuery MethodCallType = "Email/query"
	EmailGet   MethodCallType = "Email/get"
)
