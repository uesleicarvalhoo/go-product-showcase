package client

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

func (r Repository) Update(ctx context.Context, client *domain.Client) error {
	p := fromDomain(*client)

	return r.db.WithContext(ctx).Save(p).Error
}
