package client

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/models"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

// @Summary 	List Clients
// @Description List all clients clients
// @Tags 		Client
// @Param   	page  	query     int     false  "page number" 1
// @Param   	limit  	query     int     false  "items per page"
// @Accepts 	json
// @Produce 	json
// @Success 	200 		{object} 	domain.Pagination[domain.Client]
// @Failure 	400 		{object} 	models.MessageJSON
// @Router 		/clients/ [get].
func (s Service) ListClients(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "get-all-clients")
	defer span.End()

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	clients, err := s.usecase.FetchAll(ctx, page, limit)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusOK).JSON(domain.Pagination[domain.Client]{
		Page:  page,
		Items: clients,
	})
}
