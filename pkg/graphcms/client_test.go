package graphcms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClientRun(t *testing.T) {
	req := graphcms.NewRequest(`{ course { title } }`)
	testCases := []struct {
		desc               string
		req                *graphcms.Request
		clientExpectations func(client *graphcmsmocks.MockCMSRequester)

		expectedErr error
	}{
		{
			desc: "given the client makes the request correctly, nil returned",
			req:  req,

			clientExpectations: func(client *graphcmsmocks.MockCMSRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any())
			},
		},
		{
			desc: "given the client returns an error, wrapper error returned",
			req:  req,

			clientExpectations: func(client *graphcmsmocks.MockCMSRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred making request to GraphCMS: something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequester(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			client := graphcms.NewClient(clientMock)

			ctx := context.Background()
			err := client.Run(ctx, tt.req, "test")

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
