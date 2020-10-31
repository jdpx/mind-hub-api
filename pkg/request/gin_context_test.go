package request_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/stretchr/testify/assert"
)

func TestGinContext(t *testing.T) {
	ctx := context.Background()
	ginContext := gin.Context{}

	testCases := []struct {
		desc string
		ctx  context.Context

		expectedErr error
	}{
		{
			desc: "given a valid Gin Context, context returned",
			ctx:  context.WithValue(ctx, request.ContextKeyGinContext, &ginContext),
		},
		{
			desc: "given there is no value stored for the Gin Key",
			ctx:  ctx,

			expectedErr: fmt.Errorf("could not retrieve gin.Context from context"),
		},
		{
			desc: "given the context isnt the gin.Context type",
			ctx:  context.WithValue(ctx, request.ContextKeyGinContext, "not a gin type"),

			expectedErr: fmt.Errorf("gin.Context has wrong type"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			gCtx, err := request.GinContext(tt.ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &ginContext, gCtx)
			}
		})
	}
}
