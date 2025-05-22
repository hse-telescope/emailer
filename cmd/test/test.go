package main

import (
	"context"

	"github.com/hse-telescope/emailer/pkg/wrapper"
	"github.com/hse-telescope/utils/queues/kafka"
)

func main() {
	w, err := wrapper.New(kafka.QueueCredentials{
		URLs:  []string{"localhost:9092"}, // From config
		Topic: "send-email-events",        // From config
	})
	if err != nil {
		panic(err)
	}
	w.SendEmail(context.Background(), wrapper.Message{
		EMail:   "ol-sidorenkov@mail.ru",
		Title:   "lol",
		Message: "kek",
	})
}
