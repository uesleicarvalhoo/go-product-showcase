package product_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/usecase/product"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
)

func TestUpdateError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario    string
		productID   uuid.UUID
		payload     dto.UpdateProductPayload
		expectedErr string
	}{
		{
			scenario:    "when payload is valid but product doesn't exist",
			expectedErr: "record not found",
			payload:     dto.UpdateProductPayload{Name: gofakeit.BeerName()},
		},
		{
			scenario:    "when payload is empty",
			expectedErr: product.ErrNoDataForUpdate.Error(),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			// Arrange
			sut := NewUseCaseSut()

			// Action
			_, err := sut.uc.Update(context.Background(), tc.productID, tc.payload)

			// Assert
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewUseCaseSut()

	product, err := entity.NewProduct(
		gofakeit.BeerName(),
		gofakeit.SentenceSimple(),
		gofakeit.Word(),
		gofakeit.BeerStyle(),
		gofakeit.URL(),
		gofakeit.Float32(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), product)
	assert.NoError(t, err)

	payload := dto.UpdateProductPayload{
		Name:        gofakeit.BeerName(),
		Description: gofakeit.SentenceSimple(),
		Code:        gofakeit.Word(),
		Price:       gofakeit.Float32(),
		Category:    gofakeit.BeerStyle(),
		ImageURL:    gofakeit.URL(),
	}

	// Action
	updatedProduct, err := sut.uc.Update(context.Background(), product.ID, payload)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, product.ID, updatedProduct.ID)
	assert.Equal(t, updatedProduct.Name, payload.Name)
	assert.Equal(t, updatedProduct.Description, payload.Description)

	// Assert db
	dbProduct, err := sut.repo.Fetch(context.Background(), updatedProduct.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedProduct, dbProduct)

	// Assert event
	broker, ok := sut.broker.(*broker.MemoryBroker)
	if assert.True(t, ok) {
		time.Sleep(time.Second) // Wait until broker receive event

		expectedEventData := map[string]any{
			"id": product.ID.String(),
			"updated_fields": map[string]any{
				"name":        payload.Name,
				"description": payload.Description,
				"code":        payload.Code,
				"price":       payload.Price,
				"category":    payload.Category,
				"image_url":   payload.ImageURL,
			},
		}

		events := broker.GetEvents("products")

		assert.Equal(t, "products", events[0].Topic)
		assert.Equal(t, "updated", events[0].Key)
		assert.Equal(t, expectedEventData, events[0].Data)
	}
}
