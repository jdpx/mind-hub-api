package request_test

import (
	"net/http"
	"testing"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/stretchr/testify/assert"
)

const (
	correlationIDHeader = "X-Correlation-Id"
	authorizationHeader = "Authorization"
)

func TestHeadersGetHeader(t *testing.T) {
	key := fake.Word()
	value := fake.Word()

	header := http.Header{}
	header.Add(key, value)

	testCases := []struct {
		desc   string
		key    string
		header http.Header

		expectedValue string
	}{
		{
			desc:   "given a valid header, value returned",
			key:    key,
			header: header,

			expectedValue: value,
		},
		{
			desc:   "given an unknown header, empty string returned",
			key:    fake.Word(),
			header: header,

			expectedValue: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			val := request.GetHeader(tt.key, tt.header)

			assert.Equal(t, tt.expectedValue, val)
		})
	}
}

func TestHeadersGetCorrelationIDHeader(t *testing.T) {
	value := fake.Word()

	header := http.Header{}
	header.Add(correlationIDHeader, value)

	emptyHeader := http.Header{}

	testCases := []struct {
		desc   string
		header http.Header

		expectedValue string
	}{
		{
			desc:   "given a valid header, value returned",
			header: header,

			expectedValue: value,
		},
		{
			desc:   "given header with it not set, empty string returned",
			header: emptyHeader,

			expectedValue: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			val := request.GetCorrelationIDHeader(tt.header)

			assert.Equal(t, tt.expectedValue, val)
		})
	}
}

func TestHeadersGetAuthorizationHeader(t *testing.T) {
	value := fake.Word()

	header := http.Header{}
	header.Add(authorizationHeader, value)

	emptyHeader := http.Header{}

	testCases := []struct {
		desc   string
		header http.Header

		expectedValue string
	}{
		{
			desc:   "given a valid header, value returned",
			header: header,

			expectedValue: value,
		},
		{
			desc:   "given header with it not set, empty string returned",
			header: emptyHeader,

			expectedValue: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			val := request.GetAuthorizationHeader(tt.header)

			assert.Equal(t, tt.expectedValue, val)
		})
	}
}
