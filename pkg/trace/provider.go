package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.8.0"
	"go.opentelemetry.io/otel/trace"
)

var ServiceName string //nolint: gochecknoglobals

type ProviderConfig struct {
	Endpoint       string
	ServiceName    string
	ServiceVersion string
	Environment    string
	Disabled       bool
}

type Provider struct {
	Provider trace.TracerProvider
}

func NewProvider(config ProviderConfig) (Provider, error) {
	ServiceName = config.ServiceName

	if config.Disabled {
		return Provider{Provider: trace.NewNoopTracerProvider()}, nil
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.Endpoint)))
	if err != nil {
		return Provider{}, err
	}

	prv := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(config.ServiceName),
				semconv.ServiceVersionKey.String(config.ServiceVersion),
				attribute.String("environment", config.Environment),
			),
		),
	)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return Provider{Provider: prv}, nil
}

func (p *Provider) Close(ctx context.Context) error {
	if prov, ok := p.Provider.(*tracesdk.TracerProvider); ok {
		return prov.Shutdown(ctx)
	}

	return nil
}
