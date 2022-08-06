package product_test

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

			existingProduct, err := entity.NewProduct(
				gofakeit.BeerName(),
				gofakeit.SentenceSimple(),
				gofakeit.Word(),
				gofakeit.BeerStyle(),
				gofakeit.URL(),
				gofakeit.Float32(),
			)
			assert.NoError(t, err)
			if err != nil {
				t.FailNow()
			}

			err = sut.Create(context.Background(), existingProduct)
			assert.NoError(t, err)
			// Action
			dbProduct, err := sut.Fetch(ctx, existingProduct.ID)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, existingProduct.ID, dbProduct.ID)
			assert.Equal(t, existingProduct.Name, dbProduct.Name)
			assert.Equal(t, existingProduct.Description, dbProduct.Description)
		})
	}
}
