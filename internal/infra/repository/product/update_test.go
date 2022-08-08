package product_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewRepositorySut()

	product, err := entity.NewProduct(
		gofakeit.BeerName(),
		gofakeit.SentenceSimple(),
		gofakeit.Word(),
		gofakeit.BeerStyle(),
		gofakeit.URL(),
		gofakeit.Float32(),
	)
	assert.NoError(t, err)

	newName := gofakeit.BeerName()
	newDescription := gofakeit.SentenceSimple()
	newValue := gofakeit.Float32()

	// Action
	product.Name = newName
	product.Description = newDescription
	product.Price = newValue

	err = sut.Update(context.Background(), product)
	assert.NoError(t, err)

	// Assert
	dbProduct, err := sut.Fetch(context.Background(), product.ID)
	assert.NoError(t, err)

	assert.Equal(t, product.ID, dbProduct.ID)
	assert.Equal(t, newName, dbProduct.Name)
	assert.Equal(t, newDescription, dbProduct.Description)
	assert.Equal(t, dbProduct.Price, newValue)
}
