package factory

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/contracts"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
	"gorm.io/gorm"
)

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return repository.NewProductRepository(db)
}

func NewProductUseCase(db *gorm.DB, b contracts.Broker, topic string) domain.ProductUseCase {
	repo := NewProductRepository(db)

	return domain.NewProductUseCase(repo, b, topic)
}
