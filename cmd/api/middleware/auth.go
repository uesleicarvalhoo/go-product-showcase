package middleware

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/trace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewAuth(authServer string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(fiberOtel.FromCtx(c), "authorize")

		h := string(c.Request().Header.Peek("Authorization"))

		data, err := json.Encode(map[string]string{"token": h, "path": c.Path()})
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(http.StatusUnauthorized).JSON(map[string]string{"message": "auth token not found"})
		}

		res, err := otelhttp.Post(ctx, authServer, "application/json", bytes.NewBuffer(data))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(http.StatusUnauthorized).JSON(map[string]string{"message": err.Error()})
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return c.Status(res.StatusCode).JSON(map[string]string{"message": "not authorized"})
		}

		return c.Next()
	}
}
