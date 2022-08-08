package product_test

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/contracts"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/usecase/product"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
)

type Sut struct {
	uc         product.UseCase
	repo       product.Repository
	broker     contracts.Broker
	eventTopic string
}

func NewUseCaseSut() Sut {
	db, err := database.NewSQLiteMemory(database.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewProductRepository(db)
	broker := broker.NewMemory(broker.Config{})
	eventTopic := "products"
	usecase := product.New(repo, broker, eventTopic)

	return Sut{
		repo:       repo,
		uc:         usecase,
		broker:     broker,
		eventTopic: eventTopic,
	}
}
