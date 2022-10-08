package product_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestFetchAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario        string
		dbProductsCount int
		page            int
		limit           int
	}{
		{
			scenario:        "should return all db products",
			dbProductsCount: 3,
			page:            1,
			limit:           3,
		},
	}

	// Arrange
	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()
			assert.Positive(t, tc.dbProductsCount)

			// Arrange
			sut := NewRepositorySut()

			for i := 0; i < tc.dbProductsCount; i++ {
				p, err := entity.NewProduct(
					gofakeit.BeerName(),
					gofakeit.SentenceSimple(),
					gofakeit.Word(),
					gofakeit.BeerStyle(),
					gofakeit.URL(),
					gofakeit.Float32(),
				)
				assert.NoError(t, err)
				if err != nil {
					t.FailNow()
				}

				err = sut.Create(context.Background(), p)
				assert.NoError(t, err)
			}

			// Action
			products, err := sut.FetchAll(context.Background(), tc.page, tc.limit)

			// Assert
			assert.NoError(t, err)
			assert.Len(t, products, (tc.dbProductsCount))
		})
	}
}
