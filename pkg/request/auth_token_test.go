package request_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/jdpx/mind-hub-api/pkg/request"
	tools "github.com/jdpx/mind-hub-api/tools/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserID(t *testing.T) {
	testUserID := tools.GenerateTestUserID()
	ctx := context.Background()

	testCases := []struct {
		desc string
		ctx  context.Context

		expectedUserID string
		expectedErr    error
	}{
		{
			desc: "given a valid User ID in the context, id returned",
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"sub": testUserID,
			}),

			expectedUserID: testUserID,
		},
		{
			desc: "given there is no gin context, error returned",
			ctx:  context.Background(),

			expectedErr: fmt.Errorf("no auth token in context could not retrieve gin.Context from context"),
		},
		{
			desc: "given the user ID in the token claim, isnt a valid Auth0 user id, error returned",
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"sub": "invalid user ID",
			}),

			expectedErr: fmt.Errorf("token user ID is an invalid Auth0 user ID"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			_, err := request.GetUserID(tt.ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedUserID, testUserID)
			}
		})
	}
}
