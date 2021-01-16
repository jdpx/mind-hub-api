package request_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/jdpx/mind-hub-api/pkg/request"
	tools "github.com/jdpx/mind-hub-api/tools/testing"
	"github.com/stretchr/testify/assert"
)

func TestContextCorrelationID(t *testing.T) {
	cID, _ := uuid.NewV4()
	req := http.Request{Header: http.Header{}}
	req.Header.Set(correlationIDHeader, cID.String())

	ctx := context.Background()

	testCases := []struct {
		desc string
		ctx  context.Context

		expectedCID string
	}{
		{
			desc: "given a context that has a cID in the request, CID returned",
			ctx:  tools.GenerateTestGinContextWithRequest(ctx, req),

			expectedCID: cID.String(),
		},
		{
			desc: "given there is no cID in the request, random ID returned",
			ctx:  tools.GenerateTestGinContextWithRequest(ctx, http.Request{Header: http.Header{}}),

			expectedCID: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			id, _ := request.ContextCorrelationID(tt.ctx)

			if tt.expectedCID == "" {
				assert.NotEmpty(t, id)
				assert.NotEqual(t, cID, id)
			} else {
				assert.Equal(t, tt.expectedCID, id)
			}
		})
	}
}
