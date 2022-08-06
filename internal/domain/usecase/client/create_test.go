package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
)

func TestCreateError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario    string
		payload     dto.CreateClientPayload
		expectedErr string
	}{
		{
			scenario: "when payload is invalid",
			payload: dto.CreateClientPayload{
				Name: gofakeit.Name(),
			},
			expectedErr: "Phone is required",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			// Arrange
			sut := NewUseCaseSut()

			// Action
			_, err := sut.uc.Create(context.Background(), tc.payload)

			// Assert
			assert.ErrorContains(t, err, tc.expectedErr)
		})
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()

	// Arrange
	payload := dto.CreateClientPayload{
		Name:  gofakeit.BeerName(),
		Email: gofakeit.Email(),
		Phone: gofakeit.Phone(),
		Address: dto.AddressPayload{
			ZipCode: gofakeit.Zip(),
			Street:  gofakeit.Street(),
			City:    gofakeit.City(),
		},
	}

	sut := NewUseCaseSut()

	// Action
	client, err := sut.uc.Create(context.Background(), payload)

	// Assert
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, client.Name)
	assert.Equal(t, payload.Name, client.Name)
	assert.Equal(t, payload.Email, client.Email)
	assert.Equal(t, payload.Phone, client.Phone)
	assert.Equal(t, payload.Address.ZipCode, client.Address.ZipCode)
	assert.Equal(t, payload.Address.Street, client.Address.Street)
	assert.Equal(t, payload.Address.City, client.Address.City)

	// Assert event
	broker, ok := sut.broker.(*broker.MemoryBroker)
	if assert.True(t, ok) {
		time.Sleep(time.Second) // Wait until broker receive event

		events := broker.GetEvents(sut.eventTopic)

		expectedData := client

		assert.Len(t, events, 1)
		assert.Equal(t, sut.eventTopic, events[0].Topic)
		assert.Equal(t, "created", events[0].Key)
		assert.Equal(t, expectedData, events[0].Data)
	}
}
