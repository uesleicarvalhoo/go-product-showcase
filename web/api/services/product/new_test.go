package product_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/broker"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/database"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/infra/repository"
	"github.com/uesleicarvalhoo/go-product-showcase/web/api/middleware"
	"github.com/uesleicarvalhoo/go-product-showcase/web/api/services"
)

type Sut struct {
	app  *fiber.App
	repo domain.ProductRepository
}

func NewSut() Sut {
	db, err := database.NewSQLiteMemory(database.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewProductRepository(db)
	broker := broker.NewMemory(broker.Config{})
	usecase := domain.NewProductUseCase(repo, broker, "products")

	app := fiber.New()
	app.Use(middleware.NewOtel())

	services.NewProductService(app.Group("/products"), usecase, func(c *fiber.Ctx) error { return c.Next() })

	return Sut{
		app:  app,
		repo: repo,
	}
}
