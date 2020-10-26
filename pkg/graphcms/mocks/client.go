// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package graphcmsmocks is a generated GoMock package.
package graphcmsmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	graphcms "github.com/jdpx/mind-hub-api/pkg/graphcms"
	reflect "reflect"
)

// MockCMSRequester is a mock of CMSRequester interface
type MockCMSRequester struct {
	ctrl     *gomock.Controller
	recorder *MockCMSRequesterMockRecorder
}

// MockCMSRequesterMockRecorder is the mock recorder for MockCMSRequester
type MockCMSRequesterMockRecorder struct {
	mock *MockCMSRequester
}

// NewMockCMSRequester creates a new mock instance
func NewMockCMSRequester(ctrl *gomock.Controller) *MockCMSRequester {
	mock := &MockCMSRequester{ctrl: ctrl}
	mock.recorder = &MockCMSRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCMSRequester) EXPECT() *MockCMSRequesterMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockCMSRequester) Run(ctx context.Context, req *graphcms.Request, resp interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, req, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockCMSRequesterMockRecorder) Run(ctx, req, resp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCMSRequester)(nil).Run), ctx, req, resp)
}
