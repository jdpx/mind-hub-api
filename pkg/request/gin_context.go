package request

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	mindContextKey = "mind-hub-api"
	versionHeader  = "x-mind-api-version"
)

type contextKey string

func (c contextKey) String() string {
	return fmt.Sprintf("%s%s", mindContextKey, string(c))
}

const (
	// ContextKeyGinContext defines the key for the Gin Context stored within a request context
	ContextKeyGinContext = contextKey("GinContextKey")
)

// GinContext extracts the Gin Context from a wrapped context
func GinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ContextKeyGinContext)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context from context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// ContextMiddleware extracts the request context and sets it as a value on the gin context
// This allows the request context to be accessed by downstream GraphQL resolvers
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ContextKeyGinContext, c)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// VersionMiddleware sets the api version as a header for all API responses
func VersionMiddleware(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(versionHeader, version)
		c.Next()
	}
}
