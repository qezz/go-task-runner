package future

import (
	"context"
)

type Future interface {
	// poll() method gets the context, and returns the (value, error)
	// `value` represents the value returned by future
	// `error` represents the error happened while evaluating the future
	Poll(ctx *context.Context) (interface{}, error)
}
