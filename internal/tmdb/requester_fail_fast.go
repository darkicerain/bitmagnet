package tmdb

import (
	"bitmagnet-io/bitmagnet/internal/concurrency"
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

// requesterFailFast is a Requester that fails fast on subsequent requests having received an unauthorized response.
type requesterFailFast struct {
	requester      Requester
	isUnauthorized *concurrency.AtomicValue[bool]
}

func (r requesterFailFast) Request(ctx context.Context, path string, queryParams map[string]string, result interface{}) (*resty.Response, error) {
	if r.isUnauthorized.Get() {
		return nil, ErrUnauthorized
	}
	res, err := r.requester.Request(ctx, path, queryParams, result)
	if err != nil && errors.Is(err, ErrUnauthorized) {
		r.isUnauthorized.Set(true)
	}
	return res, err
}
