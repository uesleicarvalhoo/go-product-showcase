package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/models"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

// @Summary 	List Clients
// @Description List all clients clients
// @Tags 		Client
// @Accepts 	json
// @Produce 	json
// @Success 	200 		{object} 	domain.Pagination[domain.Client]
// @Failure 	400 		{object} 	models.MessageJSON
// @Router 		/clients/ [get].
func (s Service) ListProducts(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "get-all-clients")
	defer span.End()

	clients, err := s.usecase.FetchAll(ctx)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusOK).JSON(domain.Pagination[domain.Client]{
		Items: clients,
		Page:  1,
		Total: len(clients),
	})
}
