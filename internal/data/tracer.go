package data

import (
	"context"
	"github.com/go-cinch/layout/internal/conf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc"
)

func NewTracer(c *conf.Bootstrap) (tp *trace.TracerProvider, err error) {
	ctx := context.Background()

	var exporter trace.SpanExporter
	if c.Tracer.Otlp.Endpoint != "" {
		// rpc driver
		driverOpts := []otlptracegrpc.Option{
			otlptracegrpc.WithEndpoint(c.Tracer.Otlp.Endpoint),
			otlptracegrpc.WithDialOption(grpc.WithBlock()),
		}
		if c.Tracer.Otlp.Insecure {
			driverOpts = append(driverOpts, otlptracegrpc.WithInsecure())
		}
		driver := otlptracegrpc.NewClient(driverOpts...)
		exporter, err = otlptrace.New(ctx, driver)
	} else {
		// stdout driver
		if c.Tracer.Stdout.PrettyPrint {
			exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
		} else {
			exporter, err = stdouttrace.New()
		}
	}

	if err != nil {
		return
	}
	tp = trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSyncer(exporter),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(c.System.Name),
			attribute.String("exporter", "otlp"),
		)),
	)
	otel.SetTracerProvider(tp)
	return
}
