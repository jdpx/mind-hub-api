package graphcms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphcms/builder"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	tools "github.com/jdpx/mind-hub-api/tools/testing"
	"github.com/stretchr/testify/assert"
)

func TestClientRun(t *testing.T) {
	orgID := fake.CharactersN(10)
	orgIDScope := tools.GenerateTestOrganisationIDScope(orgID)

	unknownOrgID := fake.CharactersN(10)
	unknownOrgIDScope := tools.GenerateTestOrganisationIDScope(unknownOrgID)

	ctx := context.Background()
	req := graphcms.NewRequest(ctx, `{ course { title } }`)
	course := builder.NewCourseBuilder().Build()

	testCases := []struct {
		desc               string
		req                *graphcms.Request
		ctx                context.Context
		orgID              string
		clientExpectations func(client *graphcmsmocks.MockRequester)

		expectedRes graphcms.Course
		expectedErr error
	}{
		{
			desc: "given the client makes the request correctly, nil returned",
			req:  req,
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"scope": orgIDScope,
			}),

			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, course)
			},

			expectedRes: course,
		},
		{
			desc: "given there is no organisation client registered for the ID, error returned",
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"scope": unknownOrgIDScope,
			}),
			req: req,

			expectedErr: fmt.Errorf("no client registered for organisation %s", unknownOrgID),
		},
		{
			desc: "given there is no organisation ID in context, error returned",
			ctx:  context.Background(),
			req:  req,

			expectedErr: fmt.Errorf("no organisation ID in context no auth token in context could not retrieve gin.Context from context"),
		},
		{
			desc: "given the client makes the request correctly, nil returned",
			req:  req,
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"scope": orgIDScope,
			}),

			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred making request to GraphCMS: something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			clientMock := graphcmsmocks.NewMockRequester(ctrl)
			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			client := graphcms.NewClient(
				graphcms.WithOrganisationClient(orgID, clientMock),
			)

			var res graphcms.Course
			err := client.Run(tt.ctx, tt.req, &res)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedRes, res)
			}
		})
	}
}
