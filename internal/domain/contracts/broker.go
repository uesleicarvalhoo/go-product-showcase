package contracts

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type Broker interface {
	SendEvent(ctx context.Context, event dto.Event)
}
