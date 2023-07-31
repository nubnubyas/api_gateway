// Create middleware using the following template
package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func MyMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		// ...
		// if there is no 'post-handle' logic, the 'c.Next(ctx)' can be omitted.
		c.Next(ctx)
	}
}
