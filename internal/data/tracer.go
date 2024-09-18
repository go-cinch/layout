package data

import (
	"context"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
)

func NewTracer(c *conf.Bootstrap) (tp *trace.TracerProvider, err error) {
	if !c.Tracer.Enable {
		log.Info("skip initialize tracer")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var exporter trace.SpanExporter
	var resourcer *resource.Resource
	attrs := []attribute.KeyValue{semconv.ServiceNameKey.String(c.Name)}
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
		resourcer = resource.NewSchemaless(attrs...)
	} else {
		// stdout driver
		if c.Tracer.Stdout.PrettyPrint {
			exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
		} else {
			exporter, err = stdouttrace.New()
		}
		resourcer = resource.NewSchemaless(attrs...)
	}

	if err != nil {
		log.Error(err)
		err = errors.New("initialize tracer failed")
		return
	}
	providerOpts := []trace.TracerProviderOption{
		trace.WithBatcher(exporter),
		trace.WithResource(resourcer),
		trace.WithSampler(trace.TraceIDRatioBased(float64(c.Tracer.Ratio))),
	}
	tp = trace.NewTracerProvider(providerOpts...)
	otel.SetTracerProvider(tp)
	log.Info("initialize tracer success, ratio: %v", c.Tracer.Ratio)
	return
}
