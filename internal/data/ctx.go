package data

import (
	"context"
	"time"
)

// get default timeout ctx
func getDefaultTimeoutCtx(ctx context.Context) context.Context {
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	return ctx
}
