package biz

import (
	"context"

	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/layout/api/reason"
)

var (
	ErrIdempotentMissingToken = func(ctx context.Context) error {
		return reason.ErrorIllegalParameter(i18n.FromContext(ctx).T(constant.IdempotentMissingToken))
	}
	ErrTooManyRequests = func(ctx context.Context) error {
		return reason.ErrorTooManyRequests(i18n.FromContext(ctx).T(constant.TooManyRequests))
	}
	ErrNoPermission = func(ctx context.Context) error {
		return reason.ErrorForbidden(i18n.FromContext(ctx).T(constant.NoPermission))
	}
)
