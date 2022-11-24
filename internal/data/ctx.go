package data

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"time"
)

// get default timeout ctx
func getDefaultTimeoutCtx(ctx context.Context) context.Context {
	span := trace.SpanFromContext(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	return trace.ContextWithSpan(ctx, span)
}
