package client_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"gorm.io/gorm"
)

func TestFetchShouldReturnErrorWhenRecordDoesNotExist(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewRepositorySut()

	// Action
	_, err := sut.Fetch(context.Background(), uuid.Nil)

	// Assert
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestFetch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
	}{}

	// Arrange
	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			// Arrange
			ctx := context.Background()
			sut := NewRepositorySut()

			existingClient, err := entity.NewClient(
				gofakeit.BeerName(),
				gofakeit.Email(),
				gofakeit.Phone(),
				gofakeit.Zip(),
				gofakeit.Street(),
				gofakeit.City(),
			)
			assert.NoError(t, err)
			if err != nil {
				t.FailNow()
			}

			err = sut.Create(context.Background(), existingClient)
			assert.NoError(t, err)
			// Action
			dbProduct, err := sut.Fetch(ctx, existingClient.ID)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, existingClient.ID, dbProduct.ID)
			assert.Equal(t, existingClient.Name, dbProduct.Name)
			assert.Equal(t, existingClient.Email, dbProduct.Email)
			assert.Equal(t, existingClient.Phone, dbProduct.Phone)
			assert.Equal(t, existingClient.Address.City, dbProduct.Address.City)
			assert.Equal(t, existingClient.Address.Street, dbProduct.Address.Street)
			assert.Equal(t, existingClient.Address.ZipCode, dbProduct.Address.ZipCode)
		})
	}
}
