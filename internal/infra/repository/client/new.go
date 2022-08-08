package client

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository/crud"
	"gorm.io/gorm"
)

type Repository struct {
	crud.Crud[Model, domain.Client]
}

func New(db *gorm.DB) Repository {
	return Repository{crud.New(db, toDomain, fromDomain)}
}
