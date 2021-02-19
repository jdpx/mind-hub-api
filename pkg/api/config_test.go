package api_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestConfigIsLocal(t *testing.T) {
	testCases := []struct {
		desc string
		env  string

		expected bool
	}{
		{
			desc: "given local env, true returned",
			env:  "local",

			expected: true,
		},
		{
			desc: "given non local env, false returned",
			env:  fake.CharactersN(10),

			expected: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			c := api.Config{
				Env: tt.env,
			}

			res := c.IsLocal()

			assert.Equal(t, tt.expected, res)
		})
	}
}
