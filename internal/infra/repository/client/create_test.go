package client_test

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

	client, err := entity.NewClient(
		gofakeit.Name(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City())
	assert.NoError(t, err)

	// Action
	err = sut.Create(context.TODO(), client)

	// Assert
	assert.NoError(t, err)
}
