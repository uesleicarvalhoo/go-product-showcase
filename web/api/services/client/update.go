package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
	"github.com/uesleicarvalhoo/go-product-showcase/web/api/models"
)

// @Summary		Update client
// @Description	Update client data
// @Tags		Product
// @Accepts		json
// @Produce		json
// Param		id		path string true	"payload to update client data"
// Param		payload body dto.UpdateProductPayload true	"fields to update, empty fields will be ignored"
// @Success		201	{object} domain.Product
// @Failure		422	{object} models.MessageJSON
// @Failure		400	{object} models.MessageJSON
// @Router		/clients/{id} [post].
func (s Service) Update(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "update-client")
	defer span.End()

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusUnprocessableEntity).JSON(models.NewErrorMsg(err))
	}

	trace.AddSpanTags(span, map[string]string{"client_id": id.String()})

	var payload dto.UpdateClientPayload

	if err := c.BodyParser(&payload); err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusUnprocessableEntity).JSON(models.NewErrorMsg(err))
	}

	p, err := s.usecase.Update(ctx, id, payload)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusOK).JSON(p)
}
