package product_test

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
		productID          string
	}{
		{
			description:        "when id is invalid",
			expectedStatusCode: http.StatusUnprocessableEntity,
			productID:          "invalid-id",
		},
		{
			description:        "when product doesn't exist",
			expectedStatusCode: http.StatusBadRequest,
			productID:          uuid.NewString(),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s%s", tc.productID, tc.query), nil)

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

	// Action
	endpoint, err := url.Parse(fmt.Sprintf("/products/%s", existingProduct.ID.String()))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, endpoint.String(), nil)
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()

	// Assert fields
	var product domain.Product

	err = json.NewDecoder(res.Body).Decode(&product)
	assert.NoError(t, err)

	assert.Equal(t, product.ID, product.ID)
	assert.Equal(t, existingProduct.Name, product.Name)
	assert.Equal(t, existingProduct.Description, product.Description)
	assert.Equal(t, existingProduct.Price, product.Price)
}
