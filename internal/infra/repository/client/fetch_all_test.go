package client_test

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
		scenario       string
		dbClientsCount int
		page           int
		limit          int
	}{
		{
			scenario:       "should return all db clients",
			dbClientsCount: 3,
			page:           1,
			limit:          3,
		},
	}

	// Arrange
	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()
			assert.Positive(t, tc.dbClientsCount)

			// Arrange
			sut := NewRepositorySut()

			for i := 0; i < tc.dbClientsCount; i++ {
				c, err := entity.NewClient(
					gofakeit.BeerName(),
					gofakeit.Email(),
					gofakeit.Phone(),
					gofakeit.Zip(),
					gofakeit.Street(),
					gofakeit.City(),
				)
				assert.NoError(t, err)

				if err != nil {
					t.FailNow()
				}

				err = sut.Create(context.Background(), c)
				assert.NoError(t, err)
			}

			// Action
			clients, err := sut.FetchAll(context.Background(), tc.page, tc.limit)

			// Assert
			assert.NoError(t, err)
			assert.Len(t, clients, (tc.dbClientsCount))
		})
	}
}
