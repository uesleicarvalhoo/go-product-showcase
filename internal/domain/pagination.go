package domain

type Pagination[T Entity] struct {
	Items []T `json:"items"`
	Page  int `json:"page"`
}
