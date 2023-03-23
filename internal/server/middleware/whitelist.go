package middleware

import (
	"context"
	"github.com/go-cinch/layout/api/game"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func permissionWhitelist() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	whitelist["/grpc.health.v1.Health/Check"] = struct{}{}
	whitelist["/grpc.health.v1.Health/Watch"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

func idempotentBlacklist() selector.MatchFunc {
	blacklist := make(map[string]struct{})
	blacklist[game.OperationGameCreateGame] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := blacklist[operation]; ok {
			return true
		}
		return false
	}
}
