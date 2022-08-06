package product_test

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

	existingProduct, err := entity.NewProduct(
		gofakeit.BeerName(),
		gofakeit.SentenceSimple(),
		gofakeit.Word(),
		gofakeit.BeerStyle(),
		gofakeit.URL(),
		gofakeit.Float32(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), existingProduct)
	assert.NoError(t, err)

	tests := []struct {
		description        string
		path               string
		expectedStatusCode int
		productID          string
	}{
		{
			description:        "when id is invalid",
			expectedStatusCode: http.StatusUnprocessableEntity,
			productID:          "invalid-id",
		},
		{
			description:        "when product doesn't exist",
			expectedStatusCode: http.StatusUnprocessableEntity,
			productID:          uuid.NewString(),
		},
		{
			description:        "when payload is invalid",
			expectedStatusCode: http.StatusUnprocessableEntity,
			productID:          existingProduct.ID.String(),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/products/%s", tc.productID), nil)

			// Action
			res, err := sut.app.Test(req, 30)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
		})
	}
}

func TestUpdateEndpoint(t *testing.T) {
	t.Parallel()

	// Arrange
	sut := NewSut()

	existingProduct, err := entity.NewProduct(
		gofakeit.BeerName(),
		gofakeit.SentenceSimple(),
		gofakeit.Word(),
		gofakeit.BeerStyle(),
		gofakeit.URL(),
		gofakeit.Float32(),
	)
	assert.NoError(t, err)

	err = sut.repo.Create(context.Background(), existingProduct)
	assert.NoError(t, err)

	payload := dto.UpdateProductPayload{
		Name:        gofakeit.BeerName(),
		Description: gofakeit.SentenceSimple(),
		Code:        gofakeit.Word(),
		Price:       gofakeit.Float32(),
	}

	// Action
	body, err := json.Encode(payload)
	assert.NoError(t, err)

	endpoint := fmt.Sprintf("/products/%s", existingProduct.ID.String())

	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Assert fields
	var product domain.Product

	err = json.NewDecoder(res.Body).Decode(&product)
	assert.NoError(t, err)

	assert.Equal(t, product.ID, product.ID)
	assert.Equal(t, payload.Name, product.Name)
	assert.Equal(t, payload.Description, product.Description)
	assert.Equal(t, payload.Price, product.Price)
}
