package client_test

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/usecase/client"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
)

type Sut struct {
	uc         client.UseCase
	repo       client.Repository
	broker     client.Broker
	eventTopic string
}

func NewUseCaseSut() Sut {
	db, err := database.NewSQLiteMemory(database.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewClientRepository(db)
	broker := broker.NewMemory(broker.Config{})
	eventTopic := "clients"
	usecase := client.New(repo, broker, eventTopic)

	return Sut{
		repo:       repo,
		uc:         usecase,
		broker:     broker,
		eventTopic: eventTopic,
	}
}
