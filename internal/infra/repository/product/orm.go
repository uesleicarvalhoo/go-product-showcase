package product

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

type Model struct {
	ID          uuid.UUID
	Name        string
	Code        string
	Description string
	Category    string
	ImageURL    string
	Price       float32
}

func (Model) TableName() string { return "products" }

func toDomain(p Model) domain.Product {
	return domain.Product(p)
}

func fromDomain(p domain.Product) Model {
	return Model(p)
}
