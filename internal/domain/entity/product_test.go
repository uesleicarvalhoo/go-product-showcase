package entity_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestNewProductError(t *testing.T) {
	t.Parallel()

	// Arrange
	tests := []struct {
		scenario    string
		name        string
		description string
		code        string
		price       float32
		category    string
		imageURL    string
		expectedErr string
	}{
		{
			scenario:    "when name is empty",
			description: gofakeit.SentenceSimple(),
			code:        gofakeit.Word(),
			price:       gofakeit.Float32(),
			category:    gofakeit.BeerStyle(),
			expectedErr: "Name is required",
		},
		{
			scenario:    "when description is empty",
			name:        gofakeit.BeerName(),
			code:        gofakeit.Word(),
			price:       gofakeit.Float32(),
			category:    gofakeit.BeerStyle(),
			expectedErr: "Description is required",
		},
		{
			scenario:    "when code is empty",
			name:        gofakeit.BeerName(),
			description: gofakeit.SentenceSimple(),
			category:    gofakeit.BeerStyle(),
			price:       gofakeit.Float32(),
			expectedErr: "Code is required",
		},
		{
			scenario:    "when category is empty",
			name:        gofakeit.BeerName(),
			description: gofakeit.SentenceSimple(),
			code:        gofakeit.BeerStyle(),
			price:       gofakeit.Float32(),
			expectedErr: "Category is required",
		},
		{
			scenario:    "when Price is empty",
			name:        gofakeit.BeerName(),
			description: gofakeit.SentenceSimple(),
			category:    gofakeit.BeerStyle(),
			code:        gofakeit.Word(),
			expectedErr: "Price is required",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			// Action
			_, err := entity.NewProduct(tc.name, tc.description, tc.code, tc.category, tc.imageURL, tc.price)

			// Assert
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestNewProduct(t *testing.T) {
	t.Parallel()

	// Arrange
	name := gofakeit.BeerName()
	description := gofakeit.SentenceSimple()
	code := gofakeit.Word()
	price := gofakeit.Float32()
	category := gofakeit.BeerStyle()
	imageURL := gofakeit.URL()

	// Action
	p, err := entity.NewProduct(name, description, code, category, imageURL, price)

	// Assert
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, p.ID)
	assert.Equal(t, code, p.Code)
	assert.Equal(t, price, p.Price)
}
