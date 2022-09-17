package biz

import (
	v1 "github.com/go-cinch/layout/api/helloworld/v1"
)

var (
	// TooManyRequests is too many requests in a short time
	TooManyRequests = v1.ErrorTooManyRequests(v1.ErrorReason_TOO_MANY_REQUESTS.String())

	// ErrGreeterNotFound is greeter not found.
	ErrGreeterNotFound = v1.ErrorGreeterNotFound(v1.ErrorReason_GREETER_NOT_FOUND.String())
)
