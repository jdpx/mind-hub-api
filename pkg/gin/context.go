package gin

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type contextKey string

func (c contextKey) String() string {
	return "mind-hub-api" + string(c)
}

const (
	contextKeyGinContext = contextKey("GinContextKey")
)

// RequestContextFromContext ...
func RequestContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(contextKeyGinContext)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// RequestContextToContextMiddleware ...
func RequestContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKeyGinContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
