package product

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

type Service struct {
	usecase domain.ProductUseCase
}

func New(usecase domain.ProductUseCase) Service {
	return Service{usecase: usecase}
}
