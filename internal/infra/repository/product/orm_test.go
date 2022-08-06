package product //nolint:testpackage

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestToDomain(t *testing.T) {
	t.Parallel()

	// Arrange
	product, err := entity.NewProduct(
		gofakeit.BeerName(),
		gofakeit.SentenceSimple(),
		gofakeit.Word(),
		gofakeit.BeerStyle(),
		gofakeit.URL(),
		gofakeit.Float32(),
	)

	assert.NoError(t, err)
	// Action
	model := fromDomain(product)

	// Assert
	assert.Equal(t, product.ID, model.ID)
	assert.Equal(t, product.Name, model.Name)
	assert.Equal(t, product.Description, model.Description)
}

func TestFromDomain(t *testing.T) {
	t.Parallel()

	// Arrange
	productID := uuid.New()
	model := Model{
		ID:          productID,
		Name:        gofakeit.BeerName(),
		Description: gofakeit.SentenceSimple(),
	}

	// Action
	product := toDomain(model)

	// Assert
	assert.Equal(t, model.ID, product.ID)
	assert.Equal(t, model.Name, product.Name)
	assert.Equal(t, model.Description, product.Description)
}
