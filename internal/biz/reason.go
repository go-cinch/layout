package biz

import "github.com/go-cinch/layout/api/reason"

var (
	IllegalParameter = reason.ErrorIllegalParameter
	NotFound         = reason.ErrorNotFound
	TooManyRequests  = reason.ErrorTooManyRequests("too many requests, please try again later")
	DataNotChange    = reason.ErrorIllegalParameter("data has not changed")
	DuplicateField   = reason.ErrorIllegalParameter("duplicate field")
	RecordNotFound   = reason.ErrorNotFound("not found")
)
