package product

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

func (uc UseCase) Update(ctx context.Context, id uuid.UUID, payload dto.UpdateProductPayload) (entity.Product, error) {
	if reflect.DeepEqual(payload, dto.UpdateProductPayload{}) {
		return entity.Product{}, ErrNoDataForUpdate
	}

	product, err := uc.repository.Fetch(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}

	updatedFields := map[string]any{}

	if payload.Name != "" {
		product.Name = payload.Name
		updatedFields["name"] = payload.Name
	}

	if payload.Description != "" {
		product.Description = payload.Description
		updatedFields["description"] = payload.Description
	}

	if payload.Code != "" {
		product.Code = payload.Code
		updatedFields["code"] = payload.Code
	}

	if payload.Price > 0 {
		product.Price = payload.Price
		updatedFields["price"] = payload.Price
	}

	if payload.Category != "" {
		product.Category = payload.Category
		updatedFields["category"] = payload.Category
	}

	if payload.Category != "" {
		product.ImageURL = payload.ImageURL
		updatedFields["image_url"] = payload.ImageURL
	}

	if err := uc.repository.Update(ctx, product); err != nil {
		return entity.Product{}, err
	}

	go uc.broker.SendEvent(ctx, dto.Event{
		Topic: uc.eventTopic,
		Key:   "updated",
		Data: map[string]any{
			"id":             product.ID.String(),
			"updated_fields": updatedFields,
		},
	})

	return product, nil
}
