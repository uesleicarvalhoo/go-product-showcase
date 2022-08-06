package client

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

type Model struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	ZipCode string    `json:"zip_code"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
}

func (Model) TableName() string { return "clients" }

func toDomain(c Model) domain.Client {
	return domain.Client{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
		Address: domain.ClientAdress{
			ZipCode: c.ZipCode,
			Street:  c.Street,
			City:    c.City,
		},
	}
}

func fromDomain(c domain.Client) Model {
	return Model{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		ZipCode: c.Address.ZipCode,
		Street:  c.Address.Street,
		City:    c.Address.City,
	}
}
