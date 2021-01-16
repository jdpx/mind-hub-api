package testing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/request"
)

const (
	authorizationHeader = "Authorization"
)

// GenerateTestGinContext generates a gin context for use in tests
// This replicates the Gin Middleware which wraps the Requests context into the Resolver context
func GenerateTestGinContext(ctx context.Context, tokenClaims jwt.MapClaims) context.Context {
	ts := GenerateTestTokenString(tokenClaims)

	req := http.Request{Header: http.Header{}}
	req.Header.Set(authorizationHeader, fmt.Sprintf("Bearer %s", ts))

	gctx := gin.Context{
		Request: &req,
	}

	rContext := context.WithValue(ctx, request.ContextKeyGinContext, req)
	return context.WithValue(rContext, request.ContextKeyGinContext, &gctx)
}
