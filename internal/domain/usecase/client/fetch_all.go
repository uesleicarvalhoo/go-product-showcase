package client

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func (uc UseCase) FetchAll(ctx context.Context, page, limit int) ([]entity.Client, error) {
	return uc.repository.FetchAll(ctx, page, limit)
}
