package middleware

import "github.com/go-cinch/layout/api/reason"

var (
	MissingIdempotentToken = reason.ErrorIllegalParameter("idempotent token is missing")
	IdempotentTokenExpired = reason.ErrorIllegalParameter("idempotent token has expired")
)
