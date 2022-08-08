package client_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
)

func TestCreateEndpointError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		path               string
		expectedStatusCode int
		payload            dto.CreateClientPayload
	}{
		{
			description:        "when payload is invalid",
			path:               "/clients",
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}

	sut := NewSut()

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			req := httptest.NewRequest(http.MethodPost, tc.path, nil)

			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
			res.Body.Close()
		})
	}
}

func TestCreateEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		expectedStatusCode int
		payload            dto.CreateClientPayload
	}{
		{
			description:        "should create a new client and save in to repository",
			expectedStatusCode: http.StatusCreated,
			payload: dto.CreateClientPayload{
				Name:  gofakeit.Name(),
				Phone: gofakeit.Phone(),
				Email: gofakeit.Email(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
		},
	}

	sut := NewSut()

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			body, err := json.Encode(tc.payload)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer(body))
			req.Header.Add("Content-Type", "application/json")
			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
			defer res.Body.Close()

			// Assert fields
			var client domain.Client
			err = json.NewDecoder(res.Body).Decode(&client)
			assert.NoError(t, err)

			assert.NotEqual(t, uuid.Nil, client.ID)
			assert.Equal(t, tc.payload.Name, client.Name)
			assert.Equal(t, tc.payload.Phone, client.Phone)
			assert.Equal(t, tc.payload.Email, client.Email)
			assert.Equal(t, tc.payload.Address.City, client.Address.City)
			assert.Equal(t, tc.payload.Address.ZipCode, client.Address.ZipCode)
			assert.Equal(t, tc.payload.Address.Street, client.Address.Street)

			// Assert created on repository
			repoClient, err := sut.repo.Fetch(context.Background(), client.ID)
			assert.NoError(t, err)

			assert.Equal(t, client.ID, repoClient.ID)
			assert.Equal(t, client.Name, repoClient.Name)
			assert.Equal(t, client.Phone, repoClient.Phone)
			assert.Equal(t, client.Email, repoClient.Email)
			assert.Equal(t, client.Address.Street, repoClient.Address.Street)
			assert.Equal(t, client.Address.City, repoClient.Address.City)
			assert.Equal(t, client.Address.ZipCode, repoClient.Address.ZipCode)
		})
	}
}
