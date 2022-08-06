package sqlite

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&repository.ProductModel{}, &repository.ClientModel{})
}
