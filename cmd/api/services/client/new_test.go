package client_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/middleware"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/services"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
)

type Sut struct {
	app  *fiber.App
	repo domain.ClientRepository
}

func NewSut() Sut {
	db, err := database.NewSQLiteMemory(database.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewClientRepository(db)
	broker := broker.NewMemory(broker.Config{})
	usecase := domain.NewClientUseCase(repo, broker, "clients")

	app := fiber.New()
	app.Use(middleware.NewOtel())

	services.NewClientService(app.Group("/clients"), usecase, func(c *fiber.Ctx) error { return c.Next() })

	return Sut{
		app:  app,
		repo: repo,
	}
}
