package product_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestCreate(t *testing.T) {
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
	// Action
	err = sut.Create(context.TODO(), product)

	// Assert
	assert.NoError(t, err)
}
