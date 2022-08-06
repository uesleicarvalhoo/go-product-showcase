package client_test

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

	client, err := entity.NewClient(
		gofakeit.BeerName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City(),
	)
	assert.NoError(t, err)

	newName := gofakeit.BeerName()
	newEmail := gofakeit.Email()
	newPhone := gofakeit.Phone()
	newZipCode := gofakeit.Zip()
	newStreet := gofakeit.Street()
	newCity := gofakeit.City()

	// Action
	client.Name = newName
	client.Email = newEmail
	client.Phone = newPhone
	client.Address.ZipCode = newZipCode
	client.Address.Street = newStreet
	client.Address.City = newCity

	err = sut.Update(context.Background(), &client)
	assert.NoError(t, err)

	// Assert
	dbClient, err := sut.Fetch(context.Background(), client.ID)
	assert.NoError(t, err)

	assert.Equal(t, client.ID, dbClient.ID)
	assert.Equal(t, newName, dbClient.Name)
	assert.Equal(t, newEmail, dbClient.Email)
	assert.Equal(t, newPhone, dbClient.Phone)
	assert.Equal(t, newZipCode, dbClient.Address.ZipCode)
	assert.Equal(t, newStreet, dbClient.Address.Street)
	assert.Equal(t, newCity, dbClient.Address.City)
}
