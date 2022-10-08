package product_test

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
	endpoint, err := url.Parse("/products")
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, endpoint.String(), nil)
	res, err := sut.app.Test(req, 30)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()

	// Assert fields
	var response domain.Pagination[domain.Product]
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	products := response.Items

	assert.Len(t, products, 1)
	assert.Equal(t, existingProduct.ID, products[0].ID)
	assert.Equal(t, existingProduct.Name, products[0].Name)
	assert.Equal(t, existingProduct.Description, products[0].Description)
	assert.Equal(t, existingProduct.Price, products[0].Price)
}
