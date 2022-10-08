package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/models"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

// @Summary		Create client
// @Description	Create a new client and return client details
// @Tags		Client
// @Accepts		json
// @Produce		json
// @Param		payload body dto.CreateClientPayload true "data to create a new client"
// @Success		201	{object} domain.Client
// @Failure		422	{object} models.MessageJSON
// @Failure		400	{object} models.MessageJSON
// @Router		/clients/ [post].
func (s Service) Create(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "create-client")
	defer span.End()

	var payload dto.CreateClientPayload
	if err := c.BodyParser(&payload); err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusUnprocessableEntity).JSON(models.NewErrorMsg(err))
	}

	p, err := s.usecase.Create(ctx, payload)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusCreated).JSON(p)
}
