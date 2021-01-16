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

// GenerateTestGinContextWithToken generates a gin context for use in tests that has an authorisation token
// This replicates the Gin Middleware which wraps the Requests context into the Resolver context
func GenerateTestGinContextWithToken(ctx context.Context, tokenClaims jwt.MapClaims) context.Context {
	ts := GenerateTestTokenString(tokenClaims)

	req := http.Request{Header: http.Header{}}
	req.Header.Set(authorizationHeader, fmt.Sprintf("Bearer %s", ts))

	return GenerateTestGinContextWithRequest(ctx, req)
}

// GenerateTestGinContextWithRequest generates a gin context for use in tests
// This replicates the Gin Middleware which wraps the Requests context into the Resolver context
func GenerateTestGinContextWithRequest(ctx context.Context, req http.Request) context.Context {
	gctx := gin.Context{
		Request: &req,
	}

	rContext := context.WithValue(ctx, request.ContextKeyGinContext, req)
	return context.WithValue(rContext, request.ContextKeyGinContext, &gctx)
}
