package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/factory"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/config"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/logger"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

func main() {
	cfg := ServerConfig{}
	if err := config.LoadFromEnv(&cfg); err != nil {
		panic(err)
	}

	logger.InitLogger(cfg.Debug)

	// Database
	db, err := database.NewPostgreSQL(cfg.Database)
	if err != nil {
		logger.Fatalf("couldn't connect to database: %s", err)
	}

	// Broker
	broker, err := broker.NewRabbitMQ(cfg.Broker)
	if err != nil {
		logger.Fatalf("couldn't connect to broker: %s", err)
	}

	// Open Telemetry
	provider, err := trace.NewProvider(trace.ProviderConfig{
		Endpoint:       cfg.OtelURL,
		ServiceName:    cfg.ServiceName,
		ServiceVersion: cfg.ServiceVersion,
		Environment:    cfg.Environment,
		Disabled:       cfg.OtelURL == "",
	})
	if err != nil {
		logger.Fatalf("couldn't connect to otel provider: %s", err)
	}
	defer provider.Close(context.Background())

	// Configure and run server
	productUc := factory.NewProductUseCase(db, broker, cfg.BrokerProductTopic)
	clientUc := factory.NewClientUseCase(db, broker, cfg.BrokerProductTopic)

	server := NewServer(cfg, productUc, clientUc)

	go logger.Fatal(server.Listen(fmt.Sprintf(":%d", cfg.Port)))
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit
	logger.Info("shutting down server..")

	if err := server.Shutdown(); err != nil {
		logger.Errorf("graceful shutdown failed: %s\n", err)
	}
}
