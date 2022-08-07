package client

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/contracts"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

type (
	Repository = contracts.Repository[entity.Client]
	Broker     = contracts.Broker
)

type UseCase struct {
	eventTopic string
	repository contracts.Repository[entity.Client]
	broker     contracts.Broker
}

func New(r contracts.Repository[entity.Client], b contracts.Broker, eventTopic string) UseCase {
	return UseCase{
		eventTopic: eventTopic,
		broker:     b,
		repository: r,
	}
}
