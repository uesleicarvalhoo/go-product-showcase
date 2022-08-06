package product_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
)

func TestCreateError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario    string
		payload     dto.CreateProductPayload
		expectedErr string
	}{
		{
			scenario: "when payload is invalid",
			payload: dto.CreateProductPayload{
				Name:  gofakeit.BeerName(),
				Price: gofakeit.Float32(),
				Code:  gofakeit.Word(),
			},
			expectedErr: "Description is required",
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
	payload := dto.CreateProductPayload{
		Name:        gofakeit.BeerName(),
		Description: gofakeit.SentenceSimple(),
		Code:        gofakeit.Word(),
		Price:       gofakeit.Float32(),
		Category:    gofakeit.BeerStyle(),
		ImageURL:    gofakeit.URL(),
	}

	sut := NewUseCaseSut()

	// Action
	product, err := sut.uc.Create(context.Background(), payload)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, payload.Name, product.Name)
	assert.Equal(t, payload.Description, product.Description)
	assert.Equal(t, payload.Price, product.Price)

	// Assert event
	broker, ok := sut.broker.(*broker.MemoryBroker)
	if assert.True(t, ok) {
		time.Sleep(time.Second) // Wait until broker receive event

		events := broker.GetEvents(sut.eventTopic)

		expectedData := product

		assert.Len(t, events, 1)
		assert.Equal(t, sut.eventTopic, events[0].Topic)
		assert.Equal(t, "created", events[0].Key)
		assert.Equal(t, expectedData, events[0].Data)
	}
}
