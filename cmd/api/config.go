package main

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
)

type ServerConfig struct {
	ServiceName    string `env:"SERVICE_NAME,default=go-product-showcase"`
	ServiceVersion string `env:"SERVICE_VERSION,default=0.0.0"`
	Environment    string `env:"ENV,default=dev"`
	Port           int    `env:"SERVER_PORT,default=8080"`
	Debug          bool   `env:"DEBUG,default=false"`
	OtelURL        string `env:"OTEL_URL,default=http://localhost:14268/api/traces"`
	// Topic to send all product events
	BrokerProductTopic string `env:"BROKER_PRODUCT_TOPIC,default=products"`
	// Topic to send all client events
	BrokerClientTopic string `env:"BROKER_CLIENT_TOPIC,default=clients"`
	// Authentication endpoint
	AuthEndpoint string `env:"AUTHENTICATION_ENDPOINT,default=http://localhost:8000/v1/auth/authorize"`

	Database database.Config
	Broker   broker.Config
}
