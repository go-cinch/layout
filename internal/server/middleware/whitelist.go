package middleware

import (
	"context"
	"github.com/go-cinch/layout/api/greeter"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func idempotentBlacklist() selector.MatchFunc {
	blacklist := make(map[string]struct{})
	blacklist[greeter.OperationGreeterCreateGreeter] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := blacklist[operation]; ok {
			return true
		}
		return false
	}
}
