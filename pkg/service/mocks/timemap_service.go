// Code generated by MockGen. DO NOT EDIT.
// Source: timemap_service.go

// Package servicemocks is a generated GoMock package.
package servicemocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	service "github.com/jdpx/mind-hub-api/pkg/service"
	reflect "reflect"
)

// MockTimemapServicer is a mock of TimemapServicer interface
type MockTimemapServicer struct {
	ctrl     *gomock.Controller
	recorder *MockTimemapServicerMockRecorder
}

// MockTimemapServicerMockRecorder is the mock recorder for MockTimemapServicer
type MockTimemapServicerMockRecorder struct {
	mock *MockTimemapServicer
}

// NewMockTimemapServicer creates a new mock instance
func NewMockTimemapServicer(ctrl *gomock.Controller) *MockTimemapServicer {
	mock := &MockTimemapServicer{ctrl: ctrl}
	mock.recorder = &MockTimemapServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimemapServicer) EXPECT() *MockTimemapServicerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTimemapServicer) Get(ctx context.Context, uID, cID, tID string) (*service.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, uID, cID, tID)
	ret0, _ := ret[0].(*service.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTimemapServicerMockRecorder) Get(ctx, uID, cID, tID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTimemapServicer)(nil).Get), ctx, uID, cID, tID)
}

// GetByCourseID mocks base method
func (m *MockTimemapServicer) GetByCourseID(ctx context.Context, uID, cID string) ([]service.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCourseID", ctx, uID, cID)
	ret0, _ := ret[0].([]service.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCourseID indicates an expected call of GetByCourseID
func (mr *MockTimemapServicerMockRecorder) GetByCourseID(ctx, uID, cID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCourseID", reflect.TypeOf((*MockTimemapServicer)(nil).GetByCourseID), ctx, uID, cID)
}

// Update mocks base method
func (m *MockTimemapServicer) Update(ctx context.Context, uID, cID, tID, value string) (*service.Timemap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, uID, cID, tID, value)
	ret0, _ := ret[0].(*service.Timemap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTimemapServicerMockRecorder) Update(ctx, uID, cID, tID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTimemapServicer)(nil).Update), ctx, uID, cID, tID, value)
}
