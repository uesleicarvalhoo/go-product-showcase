package entity

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=4"`
	Code        string    `json:"code" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	ImageURL    string    `json:"image_url"`
	Price       float32   `json:"price" validate:"required"`
}

func NewProduct(name, description, code, category, imageURL string, price float32) (Product, error) {
	p := Product{
		ID:          uuid.New(),
		Name:        name,
		Code:        code,
		Category:    category,
		ImageURL:    imageURL,
		Description: description,
		Price:       price,
	}

	if err := p.Validate(); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (p Product) Validate() error {
	return Validate(p)
}
