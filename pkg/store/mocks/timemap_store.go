// Code generated by MockGen. DO NOT EDIT.
// Source: timemap_store.go

// Package storemocks is a generated GoMock package.
package storemocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	store "github.com/jdpx/mind-hub-api/pkg/store"
	reflect "reflect"
)

// MockTimemapRepositor is a mock of TimemapRepositor interface
type MockTimemapRepositor struct {
	ctrl     *gomock.Controller
	recorder *MockTimemapRepositorMockRecorder
}

// MockTimemapRepositorMockRecorder is the mock recorder for MockTimemapRepositor
type MockTimemapRepositorMockRecorder struct {
	mock *MockTimemapRepositor
}

// NewMockTimemapRepositor creates a new mock instance
func NewMockTimemapRepositor(ctrl *gomock.Controller) *MockTimemapRepositor {
	mock := &MockTimemapRepositor{ctrl: ctrl}
	mock.recorder = &MockTimemapRepositorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimemapRepositor) EXPECT() *MockTimemapRepositorMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTimemapRepositor) Get(ctx context.Context, uID string) (*store.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, uID)
	ret0, _ := ret[0].(*store.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTimemapRepositorMockRecorder) Get(ctx, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTimemapRepositor)(nil).Get), ctx, uID)
}

// Create mocks base method
func (m *MockTimemapRepositor) Create(ctx context.Context, tm store.Timemap) (*store.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tm)
	ret0, _ := ret[0].(*store.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTimemapRepositorMockRecorder) Create(ctx, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTimemapRepositor)(nil).Create), ctx, tm)
}

// Update mocks base method
func (m *MockTimemapRepositor) Update(ctx context.Context, tm *store.Timemap) (*store.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tm)
	ret0, _ := ret[0].(*store.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTimemapRepositorMockRecorder) Update(ctx, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTimemapRepositor)(nil).Update), ctx, tm)
}
