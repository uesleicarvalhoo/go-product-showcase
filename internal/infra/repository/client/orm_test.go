package client //nolint:testpackage

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestFromDomain(t *testing.T) {
	t.Parallel()

	// Arrange
	client, err := entity.NewClient(
		gofakeit.BeerName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City(),
	)
	assert.NoError(t, err)
	// Action
	model := fromDomain(client)

	// Assert
	assert.Equal(t, client.ID, model.ID)
	assert.Equal(t, client.Name, model.Name)
	assert.Equal(t, client.Email, model.Email)
	assert.Equal(t, client.Address.City, model.City)
	assert.Equal(t, client.Address.Street, model.Street)
	assert.Equal(t, client.Address.ZipCode, model.ZipCode)
}

func TestFromModel(t *testing.T) {
	t.Parallel()

	// Arrange
	model := Model{
		ID:      uuid.New(),
		Name:    gofakeit.Name(),
		Email:   gofakeit.Email(),
		City:    gofakeit.City(),
		Street:  gofakeit.Street(),
		ZipCode: gofakeit.Zip(),
	}
	// Action
	product := toDomain(model)

	// Assert
	assert.Equal(t, model.ID, product.ID)
	assert.Equal(t, model.Name, product.Name)
}
