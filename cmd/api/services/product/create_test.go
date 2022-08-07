package product_test

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
		payload            dto.CreateProductPayload
	}{
		{
			description:        "when payload is invalid",
			path:               "/products",
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
		})
	}
}

func TestCreateEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description        string
		expectedStatusCode int
		payload            dto.CreateProductPayload
	}{
		{
			description:        "should create a new product and save into repository",
			expectedStatusCode: http.StatusCreated,
			payload: dto.CreateProductPayload{
				Name:        gofakeit.BeerName(),
				Code:        gofakeit.Word(),
				Description: gofakeit.SentenceSimple(),
				Price:       gofakeit.Float32(),
				Category:    gofakeit.BeerStyle(),
				ImageURL:    gofakeit.URL(),
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

			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
			req.Header.Add("Content-Type", "application/json")
			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			// Assert fields
			var product domain.Product
			err = json.NewDecoder(res.Body).Decode(&product)
			assert.NoError(t, err)

			assert.NotEqual(t, uuid.Nil, product.ID)
			assert.Equal(t, tc.payload.Name, product.Name)
			assert.Equal(t, tc.payload.Code, product.Code)
			assert.Equal(t, tc.payload.Price, product.Price)
			assert.Equal(t, tc.payload.Description, product.Description)

			// Assert created on repository
			repoProduct, err := sut.repo.Fetch(context.Background(), product.ID)
			assert.NoError(t, err)

			assert.Equal(t, product.ID, repoProduct.ID)
			assert.Equal(t, product.Name, repoProduct.Name)
			assert.Equal(t, product.Description, repoProduct.Description)
			assert.Equal(t, product.Price, repoProduct.Price)
		})
	}
}
