package repository

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository/client"
	"gorm.io/gorm"
)

type (
	ClientModel = client.Model
)

func NewClientRepository(db *gorm.DB) domain.ClientRepository {
	return client.New(db)
}
