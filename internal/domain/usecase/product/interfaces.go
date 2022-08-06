package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type Repository interface {
	Create(ctx context.Context, p entity.Product) error
	Update(ctx context.Context, p *entity.Product) error
	Fetch(ctx context.Context, id uuid.UUID) (entity.Product, error)
	FetchAll(ctx context.Context) ([]entity.Product, error)
}

type Broker interface {
	SendEvent(ctx context.Context, event dto.Event)
}
