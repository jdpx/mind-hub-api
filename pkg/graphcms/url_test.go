package graphcms_test

import (
	"testing"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/stretchr/testify/assert"
)

func TestNewCMSUrl(t *testing.T) {
	testCases := []struct {
		desc string
		url  string

		expectedURL string
	}{
		{
			desc: "given a URL, GraphCMS URL generated",
			url:  "foo/bar",

			expectedURL: "https://api-eu-central-1.graphcms.com/v2/foo/bar",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			url := graphcms.NewCMSUrl(tt.url)

			assert.Equal(t, tt.expectedURL, url)
		})
	}
}
