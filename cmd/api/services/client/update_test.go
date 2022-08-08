package client_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
)

func TestUpdateEndpointError(t *testing.T) {
	t.Parallel()

	// Prepare
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

	tests := []struct {
		description        string
		path               string
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
			expectedStatusCode: http.StatusUnprocessableEntity,
			clientID:           uuid.NewString(),
		},
		{
			description:        "when payload is invalid",
			expectedStatusCode: http.StatusUnprocessableEntity,
			clientID:           existingClient.ID.String(),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/clients/%s", tc.clientID), nil)

			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			res.Body.Close()
		})
	}
}

func TestUpdateEndpoint(t *testing.T) {
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

	payload := dto.UpdateClientPayload{
		Name:  gofakeit.Name(),
		Phone: gofakeit.Phone(),
		Email: gofakeit.Email(),
		Address: dto.AddressPayload{
			ZipCode: gofakeit.Zip(),
			Street:  gofakeit.Street(),
			City:    gofakeit.City(),
		},
	}

	// Action
	body, err := json.Encode(payload)
	assert.NoError(t, err)

	endpoint := fmt.Sprintf("/clients/%s", existingClient.ID.String())

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()

	// Assert fields
	var client domain.Client

	err = json.NewDecoder(res.Body).Decode(&client)
	assert.NoError(t, err)

	assert.Equal(t, client.ID, existingClient.ID)
	assert.Equal(t, client.Name, payload.Name)
	assert.Equal(t, client.Email, payload.Email)
	assert.Equal(t, client.Phone, payload.Phone)
	assert.Equal(t, client.Address.City, payload.Address.City)
	assert.Equal(t, client.Address.Street, payload.Address.Street)
	assert.Equal(t, client.Address.ZipCode, payload.Address.ZipCode)
}
