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

		excpected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			excpected: "USER#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.UserPK(tt.id)

			assert.Equal(t, tt.excpected, pk)
		})
	}
}

func TestUserProgressSK(t *testing.T) {
	testCases := []struct {
		desc string
		id   string

		excpected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			excpected: "PROGRESS#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.ProgressSK(tt.id)

			assert.Equal(t, tt.excpected, pk)
		})
	}
}

func TestUserNoteSK(t *testing.T) {
	testCases := []struct {
		desc string
		id   string

		excpected string
	}{
		{
			desc: "given id, it returns corret key",
			id:   "Foo",

			excpected: "NOTE#Foo",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.NoteSK(tt.id)

			assert.Equal(t, tt.excpected, pk)
		})
	}
}

func TestUserTimemapSK(t *testing.T) {
	testCases := []struct {
		desc string

		excpected string
	}{
		{
			desc: "given id, it returns corret key",

			excpected: "TIMEMAP",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			pk := store.TimemapSK()

			assert.Equal(t, tt.excpected, pk)
		})
	}
}
