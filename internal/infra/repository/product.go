package repository

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository/product"
	"gorm.io/gorm"
)

type (
	ProductModel = product.Model
)

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return product.New(db)
}
