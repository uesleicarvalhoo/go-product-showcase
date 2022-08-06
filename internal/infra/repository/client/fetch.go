package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

func (r Repository) Fetch(ctx context.Context, id uuid.UUID) (domain.Client, error) {
	return r.fetch(ctx, id)
}

func (r Repository) fetch(ctx context.Context, id uuid.UUID) (domain.Client, error) {
	var p Model

	if tx := r.db.WithContext(ctx).First(&p, "id = ?", id); tx.Error != nil {
		return domain.Client{}, tx.Error
	}

	return toDomain(p), nil
}
