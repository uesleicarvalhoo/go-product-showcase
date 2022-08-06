package product_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestFetchShouldReturnErrorWhenRecordDoesNotExist(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewUseCaseSut()

	// Action
	_, err := sut.uc.Fetch(context.Background(), uuid.Nil)

	// Assert
	assert.EqualError(t, err, "record not found")
}

func TestFetch(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewUseCaseSut()

	existingProduct := entity.Product{
		ID:          uuid.New(),
		Name:        gofakeit.BeerName(),
		Description: gofakeit.SentenceSimple(),
		Price:       gofakeit.Float32(),
	}

	err := sut.repo.Create(context.Background(), existingProduct)
	assert.NoError(t, err)

	// Action
	product, err := sut.uc.Fetch(context.Background(), existingProduct.ID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, existingProduct.ID, product.ID)
	assert.Equal(t, existingProduct.Name, product.Name)
	assert.Equal(t, existingProduct.Description, product.Description)
	assert.Equal(t, existingProduct.Price, product.Price)
}
