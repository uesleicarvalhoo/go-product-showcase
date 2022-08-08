package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
)

func TestGetEndpointError(t *testing.T) {
	t.Parallel()

	// Prepare
	sut := NewSut()

	tests := []struct {
		description        string
		path               string
		query              string
		expectedStatusCode int
		clientID           string
	}{
		{
			description:        "when id is invalid",
			expectedStatusCode: http.StatusUnprocessableEntity,
			clientID:           "invalid-id",
		},
		{
			description:        "when client doesn't exist",
			expectedStatusCode: http.StatusBadRequest,
			clientID:           uuid.NewString(),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/clients/%s%s", tc.clientID, tc.query), nil)

			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
			res.Body.Close()
		})
	}
}

func TestGetEndpoint(t *testing.T) {
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
	endpoint, err := url.Parse(fmt.Sprintf("/clients/%s", existingClient.ID.String()))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, endpoint.String(), nil)
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()

	// Assert fields
	var client domain.Client

	err = json.NewDecoder(res.Body).Decode(&client)
	assert.NoError(t, err)

	assert.Equal(t, existingClient.ID, client.ID)
	assert.Equal(t, existingClient.Name, client.Name)
	assert.Equal(t, existingClient.Email, client.Email)
	assert.Equal(t, existingClient.Phone, client.Phone)
	assert.Equal(t, existingClient.Address.City, client.Address.City)
	assert.Equal(t, existingClient.Address.Street, client.Address.Street)
	assert.Equal(t, existingClient.Address.ZipCode, client.Address.ZipCode)
}
