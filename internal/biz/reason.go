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
	ErrDataNotChange = func(ctx context.Context) error {
		return reason.ErrorIllegalParameter(i18n.FromContext(ctx).T(constant.DataNotChange))
	}
	ErrDuplicateField = func(ctx context.Context, k, v string) error {
		return reason.ErrorIllegalParameter("%s `%s`: %s", i18n.FromContext(ctx).T(constant.DuplicateField), k, v)
	}
	ErrRecordNotFound = func(ctx context.Context) error {
		return reason.ErrorNotFound(i18n.FromContext(ctx).T(constant.RecordNotFound))
	}
)
