package factory

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
	"gorm.io/gorm"
)

func NewClientRepository(db *gorm.DB) domain.ClientRepository {
	return repository.NewClientRepository(db)
}

func NewClientUseCase(db *gorm.DB, b domain.ClientBroker, topic string) domain.ClientUseCase {
	repo := NewClientRepository(db)

	return domain.NewClientUseCase(repo, b, topic)
}
