package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
)

func TestUpdateError(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := NewUseCaseSut()

	existingClient, err := entity.NewClient(
		gofakeit.BeerName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), existingClient)
	assert.NoError(t, err)

	tests := []struct {
		scenario    string
		clientID    uuid.UUID
		payload     dto.UpdateClientPayload
		expectedErr string
	}{
		{
			scenario:    "when payload is valid but product doesn't exist",
			clientID:    uuid.New(),
			expectedErr: "record not found",
			payload:     dto.UpdateClientPayload{Name: gofakeit.Name()},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			// Action
			_, err := sut.uc.Update(context.Background(), tc.clientID, tc.payload)

			// Assert
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewUseCaseSut()

	client, err := entity.NewClient(
		gofakeit.BeerName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), client)
	assert.NoError(t, err)

	payload := dto.UpdateClientPayload{
		Name:  gofakeit.BeerName(),
		Email: gofakeit.Email(),
		Phone: gofakeit.Phone(),
		Address: dto.AddressPayload{
			Street:  gofakeit.Street(),
			ZipCode: gofakeit.Zip(),
			City:    gofakeit.City(),
		},
	}

	// Action
	updatedClient, err := sut.uc.Update(context.Background(), client.ID, payload)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, client.ID, updatedClient.ID)
	assert.Equal(t, updatedClient.Name, payload.Name)
	assert.Equal(t, updatedClient.Email, payload.Email)
	assert.Equal(t, updatedClient.Phone, payload.Phone)
	assert.Equal(t, updatedClient.Address.City, payload.Address.City)
	assert.Equal(t, updatedClient.Address.Street, payload.Address.Street)
	assert.Equal(t, updatedClient.Address.ZipCode, payload.Address.ZipCode)

	// Assert db
	dbClient, err := sut.repo.Fetch(context.Background(), updatedClient.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedClient, dbClient)

	// Assert event
	broker, ok := sut.broker.(*broker.MemoryBroker)
	if assert.True(t, ok) {
		time.Sleep(time.Second) // Wait until broker receive event

		expectedEventData := map[string]any{
			"id": client.ID.String(),
			"updated_fields": map[string]any{
				"name":  payload.Name,
				"email": payload.Email,
				"phone": payload.Phone,
				"address": map[string]string{
					"zip_code": payload.Address.ZipCode,
					"city":     payload.Address.City,
					"street":   payload.Address.Street,
				},
			},
		}

		events := broker.GetEvents("clients")

		assert.Equal(t, "clients", events[0].Topic)
		assert.Equal(t, "updated", events[0].Key)
		assert.Equal(t, expectedEventData, events[0].Data)
	}
}
