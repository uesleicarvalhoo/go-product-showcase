package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/contracts"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/usecase/product"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type (
	Product           = entity.Product
	ProductRepository = product.Repository
)

type ProductUseCase interface {
	Create(ctx context.Context, payload dto.CreateProductPayload) (Product, error)
	Fetch(ctx context.Context, id uuid.UUID) (Product, error)
	FetchAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id uuid.UUID, payload dto.UpdateProductPayload) (entity.Product, error)
}

func NewProductUseCase(r product.Repository, b contracts.Broker, eventTopic string) product.UseCase {
	return product.New(r, b, eventTopic)
}
