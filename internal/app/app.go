package app

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/hse-telescope/emailer/internal/config"
	"github.com/hse-telescope/emailer/internal/consumer"
	"github.com/hse-telescope/emailer/internal/providers/email"
	"github.com/hse-telescope/emailer/internal/server"
)

type App struct {
	emailProvider email.Provider
	consumer      consumer.Consumer
	server        *server.Server
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

	server := server.New(conf)

	return App{
		server:        server,
		emailProvider: emailProvider,
		consumer:      consumer,
	}, nil
}

func (a App) Run(ctx context.Context) error {
	go func() {
		a.server.Start()
	}()
	return a.consumer.Consume(ctx)
}

func (a App) Shutdown() {
}
