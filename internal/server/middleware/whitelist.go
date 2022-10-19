package middleware

import (
	"context"
	v1 "github.com/go-cinch/layout/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func idempotentBlacklist() selector.MatchFunc {
	blacklist := make(map[string]struct{})
	blacklist[v1.OperationHelloworldCreateGreeter] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := blacklist[operation]; ok {
			return true
		}
		return false
	}
}
