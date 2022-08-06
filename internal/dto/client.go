package dto

type AddressPayload struct {
	ZipCode string `json:"zip_code" validate:"omitempty"`
	Street  string `json:"street" validate:"omitempty"`
	City    string `json:"city" validate:"omitempty"`
}

type CreateClientPayload struct {
	Name    string         `json:"name" validate:"required,min=4"`
	Email   string         `json:"email" validate:"required,email"`
	Phone   string         `json:"phone" validate:"required"`
	Address AddressPayload `json:"address" validate:"required"`
}

type UpdateClientPayload struct {
	Name    string         `json:"name" validate:"omitempty,min=4"`
	Email   string         `json:"email" validate:"omitempty,email"`
	Phone   string         `json:"phone" validate:"omitempty"`
	Address AddressPayload `json:"address" validate:"omitempty"`
}
