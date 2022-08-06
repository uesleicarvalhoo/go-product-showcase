package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func (uc UseCase) Fetch(ctx context.Context, id uuid.UUID) (entity.Client, error) {
	return uc.repository.Fetch(ctx, id)
}
