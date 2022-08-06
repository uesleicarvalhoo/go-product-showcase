package entity

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/validator"
)

type Client struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" validate:"required,min=4"`
	Email   string    `json:"email" validate:"required,email"`
	Phone   string    `json:"phone" validate:"required"`
	Address Address   `json:"address" validate:"required"`
}

type Address struct {
	ZipCode string `json:"zip_code" validate:"required"`
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
}

func NewClient(name, email, phone, zipCode, street, city string) (Client, error) {
	client := Client{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
		Phone: phone,
		Address: Address{
			ZipCode: zipCode,
			Street:  street,
			City:    city,
		},
	}

	if err := validator.Validate(client); err != nil {
		return Client{}, err
	}

	return client, nil
}
