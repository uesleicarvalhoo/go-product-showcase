package product

import (
	"context"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"gorm.io/gorm/clause"
)

func (r Repository) Create(ctx context.Context, product domain.Product) error {
	p := fromDomain(product)

	return r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&p).Error
}
