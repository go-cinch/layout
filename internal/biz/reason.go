package biz

import (
	v1 "github.com/go-cinch/layout/api/helloworld/v1"
)

var (
	IllegalParameter = v1.ErrorIllegalParameter
	NotFound         = v1.ErrorNotFound
	TooManyRequests  = v1.ErrorTooManyRequests("too many requests, please try again later")
	DataNotChange    = v1.ErrorIllegalParameter("data has not changed")
	DuplicateField   = v1.ErrorIllegalParameter("duplicate field")
	RecordNotFound   = v1.ErrorNotFound("not found")
)
