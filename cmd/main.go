package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hse-telescope/emailer/internal/app"
	"github.com/hse-telescope/emailer/internal/config"
)

func main() {
	stopCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)

	configPath := os.Args[1]
	fmt.Println(configPath)
	conf, err := config.Parse(configPath)
	if err != nil {
		log.Fatalf("Error during config parse: %+v", err)
	}

	app, err := app.New(stopCtx, conf)
	if err != nil {
		panic(err)
	}

	defer cancel()
	go func() {
		err = app.Run(stopCtx)
		if err != nil {
			log.Fatalf("Error during app run: %+v", err)
		}
	}()
	<-stopCtx.Done()

	app.Shutdown()
}
