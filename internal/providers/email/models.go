package email

import "github.com/hse-telescope/emailer/pkg/wrapper"

type Message struct {
	Subject string
	Body    string
}

func WrapperMessageToProviderMessage(in wrapper.Message) (out Message) {
	return Message{
		Subject: in.Title,
		Body:    in.Message,
	}
}
