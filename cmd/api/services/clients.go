package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/services/client"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
)

type ClientService interface {
	GetDetails(c *fiber.Ctx) error
	ListClients(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

func NewClientService(r fiber.Router, usecase domain.ClientUseCase, authMiddleware fiber.Handler) ClientService {
	service := client.New(usecase)

	r.Post("/", authMiddleware, service.Create)
	r.Get("/", service.ListClients)
	r.Get("/:id", service.GetDetails)
	r.Post("/:id", authMiddleware, service.Update)

	return service
}
