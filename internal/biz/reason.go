package biz

import (
	"context"

	"github.com/go-cinch/common/constant"
	"github.com/go-cinch/common/middleware/i18n"
	"github.com/go-cinch/layout/api/reason"
)

var (
	ErrIdempotentTokenExpired = func(ctx context.Context) error {
		return i18n.NewError(ctx, constant.IdempotentTokenExpired, reason.ErrorIllegalParameter)
	}

	ErrTooManyRequests = func(ctx context.Context) error {
		return reason.ErrorTooManyRequests(i18n.FromContext(ctx).T(constant.TooManyRequests))
	}
	ErrDataNotChange = func(ctx context.Context, args ...string) error {
		return i18n.NewError(ctx, constant.DataNotChange, reason.ErrorIllegalParameter, args...)
	}
	ErrDuplicateField = func(ctx context.Context, args ...string) error {
		return i18n.NewError(ctx, constant.DuplicateField, reason.ErrorIllegalParameter, args...)
	}
	ErrRecordNotFound = func(ctx context.Context, args ...string) error {
		return i18n.NewError(ctx, constant.RecordNotFound, reason.ErrorNotFound, args...)
	}
	ErrInternal = func(ctx context.Context, args ...string) error {
		return i18n.NewError(ctx, constant.InternalError, reason.ErrorInternal, args...)
	}
	ErrIllegalParameter = func(ctx context.Context, args ...string) error {
		return i18n.NewError(ctx, constant.IllegalParameter, reason.ErrorIllegalParameter, args...)
	}
)
