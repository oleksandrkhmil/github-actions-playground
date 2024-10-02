package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/oleksandrkhmil/github-actions-playground/internal/app/service"
	"github.com/oleksandrkhmil/github-actions-playground/internal/config"

	"github.com/caarlos0/env/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Service run error", "error_message", err.Error())
		os.Exit(1)
	}
}

func run() error {
	var c config.Config
	if err := env.Parse(&c); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	shutdown, err := service.Run(context.Background(), c)
	if err != nil {
		return fmt.Errorf("run service: %w", err)
	}

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-shutdownChannel

	if err := shutdown(); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

// to trigger change
