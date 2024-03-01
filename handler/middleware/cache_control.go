package middleware

import (
	"context"
	"strconv"
	"time"

	hzapp "github.com/cloudwego/hertz/pkg/app"
)

type CacheControlMiddleware struct {
	MaxAge time.Duration
	Public bool
}

type CacheControlOption func(*CacheControlMiddleware)

func NewCacheControlMiddleware(opts ...CacheControlOption) *CacheControlMiddleware {
	c := &CacheControlMiddleware{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *CacheControlMiddleware) CacheControl() hzapp.HandlerFunc {
	value := ""
	if c.Public {
		value = "public,"
	}
	if c.MaxAge > 0 {
		value = "max-age=" + strconv.FormatInt(int64(c.MaxAge.Seconds()), 10)
	}
	return func(_ctx context.Context, ctx *hzapp.RequestContext) {
		ctx.Header("Cache-Control", value)
	}
}

func WithMaxAge(maxAge time.Duration) CacheControlOption {
	return func(c *CacheControlMiddleware) {
		c.MaxAge = maxAge
	}
}

func WithPublic(public bool) CacheControlOption {
	return func(c *CacheControlMiddleware) {
		c.Public = public
	}
}
