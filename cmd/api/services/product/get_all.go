package product

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/cmd/api/models"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
)

// @Summary 	List Products
// @Description List all products products
// @Tags 		Product
// @Accepts 	json
// @Produce 	json
// @Success 	200 		{object} 	[]domain.Product
// @Failure 	400 		{object} 	models.MessageJSON
// @Router 		/products/ [get].
func (s Service) ListProducts(c *fiber.Ctx) error {
	ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "get-all-products")
	defer span.End()

	products, err := s.usecase.FetchAll(ctx)
	if err != nil {
		trace.AddSpanError(span, err)

		return c.Status(http.StatusBadRequest).JSON(models.NewErrorMsg(err))
	}

	return c.Status(http.StatusOK).JSON(products)
}
