package app

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/hse-telescope/emailer/internal/config"
	"github.com/hse-telescope/emailer/internal/consumer"
	"github.com/hse-telescope/emailer/internal/providers/email"
)

type App struct {
	emailProvider email.Provider
	consumer      consumer.Consumer
}

func New(ctx context.Context, conf config.Config) (App, error) {
	fmt.Printf("%+v\n", conf)

	addr := fmt.Sprintf("%s:%d", conf.EmailCredentials.Host, conf.EmailCredentials.Port)
	auth := smtp.PlainAuth(
		"",
		conf.EmailCredentials.Email,
		conf.EmailCredentials.Password,
		conf.EmailCredentials.Host,
	)

	emailProvider := email.NewEmailProvider(addr, auth, conf.EmailCredentials.Email)
	consumer, err := consumer.New(emailProvider, conf.QueueCredentials)
	if err != nil {
		return App{}, err
	}

	return App{
		emailProvider: emailProvider,
		consumer:      consumer,
	}, nil
}

func (a App) Run(ctx context.Context) error {
	return a.consumer.Consume(ctx)
}

func (a App) Shutdown() {
}
