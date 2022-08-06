package product

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

func (r Repository) Update(ctx context.Context, product *domain.Product) error {
	p := fromDomain(*product)

	return r.db.WithContext(ctx).Save(p).Error
}
