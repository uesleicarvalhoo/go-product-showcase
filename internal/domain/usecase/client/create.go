package client

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

func (uc UseCase) Create(ctx context.Context, payload dto.CreateClientPayload) (entity.Client, error) {
	c, err := entity.NewClient(
		payload.Name,
		payload.Email,
		payload.Phone,
		payload.Address.ZipCode,
		payload.Address.Street,
		payload.Address.City,
	)
	if err != nil {
		return entity.Client{}, err
	}

	if err := uc.repository.Create(ctx, c); err != nil {
		return entity.Client{}, err
	}

	go uc.broker.SendEvent(ctx, dto.Event{Topic: uc.eventTopic, Key: "created", Data: c})

	return c, nil
}
