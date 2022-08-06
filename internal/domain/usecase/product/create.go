package product

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

func (uc UseCase) Create(ctx context.Context, payload dto.CreateProductPayload) (entity.Product, error) {
	p, err := entity.NewProduct(
		payload.Name, payload.Description, payload.Code, payload.Category, payload.ImageURL, payload.Price,
	)
	if err != nil {
		return entity.Product{}, err
	}

	if err := uc.repository.Create(ctx, p); err != nil {
		return entity.Product{}, err
	}

	go uc.broker.SendEvent(ctx, dto.Event{Topic: uc.eventTopic, Key: "created", Data: p})

	return p, nil
}
