package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
	"go.opentelemetry.io/otel/trace"
)

func NewOtel() fiber.Handler {
	return fiberOtel.New(fiberOtel.Config{
		SpanName: "http/request",
		TracerStartAttributes: []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindConsumer),
		},
	})
}
