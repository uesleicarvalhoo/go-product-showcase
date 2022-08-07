package product

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/contracts"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

type (
	Repository = contracts.Repository[entity.Product]
)

type UseCase struct {
	eventTopic string
	repository Repository
	broker     contracts.Broker
}

func New(r Repository, b contracts.Broker, eventTopic string) UseCase {
	return UseCase{
		eventTopic: eventTopic,
		broker:     b,
		repository: r,
	}
}
