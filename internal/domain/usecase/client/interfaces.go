package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type Repository interface {
	Create(ctx context.Context, p entity.Client) error
	Update(ctx context.Context, p *entity.Client) error
	Fetch(ctx context.Context, id uuid.UUID) (entity.Client, error)
	FetchAll(ctx context.Context) ([]entity.Client, error)
}

type Broker interface {
	SendEvent(ctx context.Context, event dto.Event)
}
