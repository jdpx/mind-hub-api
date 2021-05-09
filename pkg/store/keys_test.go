package store_test

import (
	"testing"

	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUserPK(t *testing.T) {
	testCases := []struct {
		desc string
		id   string

		expected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			expected: "USER#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.UserPK(tt.id)

			assert.Equal(t, tt.expected, pk)
		})
	}
}

func TestUserProgressSK(t *testing.T) {
	testCases := []struct {
		desc string
		id   string

		expected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			expected: "PROGRESS#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.ProgressSK(tt.id)

			assert.Equal(t, tt.expected, pk)
		})
	}
}

func TestUserNoteSK(t *testing.T) {
	testCases := []struct {
		desc string
		id   string

		expected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			expected: "NOTE#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.NoteSK(tt.id)

			assert.Equal(t, tt.expected, pk)
		})
	}
}

func TestUserTimemapSK(t *testing.T) {
	testCases := []struct {
		desc      string
		courseID  string
		timemapID string

		expected string
	}{
		{
			desc:      "given id, it returns corret key",
			courseID:  "Foo",
			timemapID: "Bar",

			expected: "COURSE#FooTIMEMAP#Bar",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.TimemapSK(tt.courseID, tt.timemapID)

			assert.Equal(t, tt.expected, pk)
		})
	}
}

func TestUserCourseTimemapSK(t *testing.T) {
	testCases := []struct {
		desc      string
		courseID  string
		timemapID string

		expected string
	}{
		{
			desc:     "given id, it returns corret key",
			courseID: "Foo",

			expected: "COURSE#FooTIMEMAP",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.CourseTimemapsSK(tt.courseID)

			assert.Equal(t, tt.expected, pk)
		})
	}
}
