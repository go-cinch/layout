package data

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// get default timeout ctx
func getDefaultTimeoutCtx(ctx context.Context) context.Context {
	// get span and create new ctx since current ctx maybe timeout
	span := trace.SpanFromContext(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	return trace.ContextWithSpan(ctx, span)
}
