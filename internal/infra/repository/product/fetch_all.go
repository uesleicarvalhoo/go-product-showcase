package product

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

func (r Repository) FetchAll(ctx context.Context) ([]domain.Product, error) {
	return r.fetchAll(ctx)
}

func (r Repository) fetchAll(ctx context.Context) ([]domain.Product, error) {
	var models []Model

	if tx := r.db.WithContext(ctx).Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	products := make([]domain.Product, len(models))
	for i, p := range models {
		products[i] = toDomain(p)
	}

	return products, nil
}
