package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
)

func TestGetAllEndpoint(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewSut()

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
	endpoint, err := url.Parse("/clients")
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, endpoint.String(), nil)
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Assert fields
	var clients []domain.Client

	err = json.NewDecoder(res.Body).Decode(&clients)
	assert.NoError(t, err)

	assert.Len(t, clients, 1)
	assert.Equal(t, existingClient.ID, clients[0].ID)
	assert.Equal(t, existingClient.Name, clients[0].Name)
	assert.Equal(t, existingClient.Email, clients[0].Email)
	assert.Equal(t, existingClient.Phone, clients[0].Phone)
	assert.Equal(t, existingClient.Address.City, clients[0].Address.City)
	assert.Equal(t, existingClient.Address.Street, clients[0].Address.Street)
	assert.Equal(t, existingClient.Address.ZipCode, clients[0].Address.ZipCode)
}
