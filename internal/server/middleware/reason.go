package middleware

import v1 "github.com/go-cinch/layout/api/helloworld/v1"

var (
	MissingIdempotentToken = v1.ErrorIllegalParameter("idempotent token is missing")
	IdempotentTokenExpired = v1.ErrorIllegalParameter("idempotent token has expired")
)
