package crud

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"gorm.io/gorm"
)

type Crud[Model any, Entity domain.Entity] struct {
	db *gorm.DB
	// Contert DB Model to domain Entity
	toDomain func(m Model) Entity
	// Convert domain Entity to DB Model
	fromDomain func(e Entity) Model
}

func New[M any, E domain.Entity](db *gorm.DB, toDomain func(M) E, fromDomain func(e E) M) Crud[M, E] {
	return Crud[M, E]{
		db:         db,
		toDomain:   toDomain,
		fromDomain: fromDomain,
	}
}
