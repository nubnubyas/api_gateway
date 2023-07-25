package middleware

import (
	"context"
	"time"

	"github.com/Narasimha1997/ratelimiter"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RateLimiter(limit uint64, interval time.Duration) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		limiter := ratelimiter.NewSyncLimiter(limit, interval)
		allowed, err := limiter.ShouldAllow(1)
		if err != nil {
			hlog.Error(err)
		}
		if allowed {
			c.Next(ctx)
		}
	}
}
