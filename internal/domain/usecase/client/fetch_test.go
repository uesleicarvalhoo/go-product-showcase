package client_test

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

	existingClient, err := entity.NewClient(
		gofakeit.BeerName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Zip(),
		gofakeit.Street(),
		gofakeit.City(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), existingClient)
	assert.NoError(t, err)

	// Action
	client, err := sut.uc.Fetch(context.Background(), existingClient.ID)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, existingClient.ID, client.ID)
	assert.Equal(t, existingClient.Name, client.Name)
	assert.Equal(t, existingClient.Email, client.Email)
	assert.Equal(t, existingClient.Phone, client.Phone)
	assert.Equal(t, existingClient.Address.ZipCode, client.Address.ZipCode)
	assert.Equal(t, existingClient.Address.Street, client.Address.Street)
	assert.Equal(t, existingClient.Address.City, client.Address.City)
}
