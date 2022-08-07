package contracts

import (
	"context"

	"github.com/google/uuid"
)

type Repository[Entity any] interface {
	Create(ctx context.Context, e Entity) error
	Update(ctx context.Context, e Entity) error
	Fetch(ctx context.Context, id uuid.UUID) (Entity, error)
	FetchAll(ctx context.Context) ([]Entity, error)
}
