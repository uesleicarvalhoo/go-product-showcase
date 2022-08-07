package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/models"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

// Get product details
// @Summary 	Get Client
// @Description Get clients details
// @Tags 		Client
// @Accepts 	json
// @Produce 	json
// @Param		id			path		string				true	"the uuid of client"
// @Success		200 		{object}	domain.Client
// @Failure		400 		{object}	models.MessageJSON
// @Router			/clients/{id} [get].
func (s Service) GetDetails(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "get-client-by-id")
	defer span.End()

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusUnprocessableEntity).JSON(models.NewErrorMsg(err))
	}

	trace.AddSpanTags(span, map[string]string{"client_id": id.String()})

	product, err := s.usecase.Fetch(ctx, id)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusOK).JSON(product)
}
