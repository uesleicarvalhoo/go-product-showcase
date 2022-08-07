package client

import (
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

type Service struct {
	usecase domain.ClientUseCase
}

func New(usecase domain.ClientUseCase) Service {
	return Service{usecase: usecase}
}
