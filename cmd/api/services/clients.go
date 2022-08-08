package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/services/client"
)

type ClienttService interface {
	GetDetails(c *fiber.Ctx) error
	ListProducts(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

func NewClientService(r fiber.Router, usecase domain.ClientUseCase, authMiddleware fiber.Handler) ProductService {
	service := client.New(usecase)

	r.Post("/", authMiddleware, service.Create)
	r.Get("/", service.ListProducts)
	r.Get("/:id", service.GetDetails)
	r.Post("/:id", authMiddleware, service.Update)

	return service
}
