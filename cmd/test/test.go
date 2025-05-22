package main

import (
	"context"

	"github.com/hse-telescope/emailer/pkg/wrapper"
	"github.com/hse-telescope/utils/queues/kafka"
)

func main() {
	w, err := wrapper.New(kafka.QueueCredentials{
		URLs:  []string{"localhost:9092"},
		Topic: "send-email-events",
	})
	if err != nil {
		panic(err)
	}
	w.SendEmail(context.Background(), wrapper.Message{
		EMail:   "o_sidorenkov@mail.ru",
		Title:   "lol",
		Message: "kek",
	})
}
