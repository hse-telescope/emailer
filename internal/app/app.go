package app

import (
	"github.com/hse-telescope/emailer/internal/consumer"
	"github.com/hse-telescope/emailer/internal/providers/email"
)

type App struct {
	emailProvider email.Provider
	consumer      consumer.Consumer
}
