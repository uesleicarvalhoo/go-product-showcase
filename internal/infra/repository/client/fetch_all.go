package client

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

func (r Repository) FetchAll(ctx context.Context) ([]domain.Client, error) {
	return r.fetchAll(ctx)
}

func (r Repository) fetchAll(ctx context.Context) ([]domain.Client, error) {
	var models []Model

	if tx := r.db.WithContext(ctx).Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	clients := make([]domain.Client, len(models))
	for i, p := range models {
		clients[i] = toDomain(p)
	}

	return clients, nil
}
