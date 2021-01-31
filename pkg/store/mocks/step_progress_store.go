// Code generated by MockGen. DO NOT EDIT.
// Source: step_progress_store.go

// Package storemocks is a generated GoMock package.
package storemocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	store "github.com/jdpx/mind-hub-api/pkg/store"
	reflect "reflect"
)

// MockStepProgressRepositor is a mock of StepProgressRepositor interface
type MockStepProgressRepositor struct {
	ctrl     *gomock.Controller
	recorder *MockStepProgressRepositorMockRecorder
}

// MockStepProgressRepositorMockRecorder is the mock recorder for MockStepProgressRepositor
type MockStepProgressRepositorMockRecorder struct {
	mock *MockStepProgressRepositor
}

// NewMockStepProgressRepositor creates a new mock instance
func NewMockStepProgressRepositor(ctrl *gomock.Controller) *MockStepProgressRepositor {
	mock := &MockStepProgressRepositor{ctrl: ctrl}
	mock.recorder = &MockStepProgressRepositorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStepProgressRepositor) EXPECT() *MockStepProgressRepositorMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStepProgressRepositor) Get(ctx context.Context, sID, uID string) (*store.StepProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, sID, uID)
	ret0, _ := ret[0].(*store.StepProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStepProgressRepositorMockRecorder) Get(ctx, sID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStepProgressRepositor)(nil).Get), ctx, sID, uID)
}

// GetCompletedByStepID mocks base method
func (m *MockStepProgressRepositor) GetCompletedByStepID(ctx context.Context, uID string, ids ...string) ([]*store.StepProgress, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, uID}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCompletedByStepID", varargs...)
	ret0, _ := ret[0].([]*store.StepProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompletedByStepID indicates an expected call of GetCompletedByStepID
func (mr *MockStepProgressRepositorMockRecorder) GetCompletedByStepID(ctx, uID interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, uID}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletedByStepID", reflect.TypeOf((*MockStepProgressRepositor)(nil).GetCompletedByStepID), varargs...)
}

// Start mocks base method
func (m *MockStepProgressRepositor) Start(ctx context.Context, sID, uID string) (*store.StepProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx, sID, uID)
	ret0, _ := ret[0].(*store.StepProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockStepProgressRepositorMockRecorder) Start(ctx, sID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStepProgressRepositor)(nil).Start), ctx, sID, uID)
}

// Complete mocks base method
func (m *MockStepProgressRepositor) Complete(ctx context.Context, sID, uID string) (*store.StepProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Complete", ctx, sID, uID)
	ret0, _ := ret[0].(*store.StepProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Complete indicates an expected call of Complete
func (mr *MockStepProgressRepositorMockRecorder) Complete(ctx, sID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockStepProgressRepositor)(nil).Complete), ctx, sID, uID)
}
