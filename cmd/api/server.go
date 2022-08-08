package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	_ "github.com/uesleicarvalhoo/go-product-showcase/docs"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/middleware"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/services"
)

func NewServer(
	cfg ServerConfig, productSvc domain.ProductUseCase, clientSvc domain.ClientUseCase,
) *fiber.App {
	server := fiber.New(fiber.Config{
		AppName:               cfg.ServiceName,
		DisableStartupMessage: !cfg.Debug,
		JSONEncoder:           json.Encode,
		JSONDecoder:           json.Decode,
	})

	authMiddleware := middleware.NewAuth(cfg.AuthEndpoint)

	server.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		middleware.NewOtel(),
		middleware.NewLogger(cfg.ServiceName, cfg.ServiceVersion),
	)

	server.Get("/swagger/*", swagger.HandlerDefault)

	services.NewProductService(server.Group("/products"), productSvc, authMiddleware)
	services.NewClientService(server.Group("/clients"), clientSvc, authMiddleware)

	return server
}
