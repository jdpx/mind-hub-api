// Code generated by MockGen. DO NOT EDIT.
// Source: coure_progress_repository.go

// Package storemocks is a generated GoMock package.
package storemocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	store "github.com/jdpx/mind-hub-api/pkg/store"
	reflect "reflect"
)

// MockCourseRepositor is a mock of CourseRepositor interface
type MockCourseRepositor struct {
	ctrl     *gomock.Controller
	recorder *MockCourseRepositorMockRecorder
}

// MockCourseRepositorMockRecorder is the mock recorder for MockCourseRepositor
type MockCourseRepositorMockRecorder struct {
	mock *MockCourseRepositor
}

// NewMockCourseRepositor creates a new mock instance
func NewMockCourseRepositor(ctrl *gomock.Controller) *MockCourseRepositor {
	mock := &MockCourseRepositor{ctrl: ctrl}
	mock.recorder = &MockCourseRepositorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCourseRepositor) EXPECT() *MockCourseRepositorMockRecorder {
	return m.recorder
}

// GetProgressForCourse mocks base method
func (m *MockCourseRepositor) GetProgressForCourse(ctx context.Context, cID, uID string) (*store.Progress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgressForCourse", ctx, cID, uID)
	ret0, _ := ret[0].(*store.Progress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgressForCourse indicates an expected call of GetProgressForCourse
func (mr *MockCourseRepositorMockRecorder) GetProgressForCourse(ctx, cID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgressForCourse", reflect.TypeOf((*MockCourseRepositor)(nil).GetProgressForCourse), ctx, cID, uID)
}

// StartCourse mocks base method
func (m *MockCourseRepositor) StartCourse(ctx context.Context, cID, uID string) (*store.Progress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCourse", ctx, cID, uID)
	ret0, _ := ret[0].(*store.Progress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCourse indicates an expected call of StartCourse
func (mr *MockCourseRepositorMockRecorder) StartCourse(ctx, cID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCourse", reflect.TypeOf((*MockCourseRepositor)(nil).StartCourse), ctx, cID, uID)
}
