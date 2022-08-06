package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

func (uc UseCase) Update(ctx context.Context, id uuid.UUID, payload dto.UpdateClientPayload) (entity.Client, error) {
	client, err := uc.repository.Fetch(ctx, id)
	if err != nil {
		return entity.Client{}, err
	}

	updatedFields := map[string]any{}

	if payload.Name != "" {
		client.Name = payload.Name
		updatedFields["name"] = payload.Name
	}

	if payload.Email != "" {
		client.Email = payload.Email
		updatedFields["email"] = payload.Email
	}

	if payload.Phone != "" {
		client.Phone = payload.Phone
		updatedFields["phone"] = payload.Phone
	}

	if updatedAdress := updateAdress(&client, payload.Address); len(updatedAdress) > 0 {
		updatedFields["address"] = updatedAdress
	}

	if err := uc.repository.Update(ctx, &client); err != nil {
		return entity.Client{}, err
	}

	if len(updatedFields) == 0 {
		return entity.Client{}, ErrNoDataForUpdate
	}

	go uc.broker.SendEvent(ctx, dto.Event{
		Topic: uc.eventTopic,
		Key:   "updated",
		Data: map[string]any{
			"id":             client.ID.String(),
			"updated_fields": updatedFields,
		},
	})

	return client, nil
}

func updateAdress(c *entity.Client, payload dto.AddressPayload) map[string]string {
	updatedFields := map[string]string{}

	if payload.City != "" {
		c.Address.City = payload.City
		updatedFields["city"] = payload.City
	}

	if payload.Street != "" {
		c.Address.Street = payload.Street
		updatedFields["street"] = payload.Street
	}

	if payload.ZipCode != "" {
		c.Address.ZipCode = payload.ZipCode
		updatedFields["zip_code"] = payload.ZipCode
	}

	return updatedFields
}
