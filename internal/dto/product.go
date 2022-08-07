package dto

type CreateProductPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Code        string  `json:"code"`
	Price       float32 `json:"price"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}

type UpdateProductPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Code        string  `json:"code"`
	Price       float32 `json:"price"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}
