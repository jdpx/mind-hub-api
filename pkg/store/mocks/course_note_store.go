// Code generated by MockGen. DO NOT EDIT.
// Source: course_note_store.go

// Package storemocks is a generated GoMock package.
package storemocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	store "github.com/jdpx/mind-hub-api/pkg/store"
	reflect "reflect"
)

// MockCourseNoteRepositor is a mock of CourseNoteRepositor interface
type MockCourseNoteRepositor struct {
	ctrl     *gomock.Controller
	recorder *MockCourseNoteRepositorMockRecorder
}

// MockCourseNoteRepositorMockRecorder is the mock recorder for MockCourseNoteRepositor
type MockCourseNoteRepositorMockRecorder struct {
	mock *MockCourseNoteRepositor
}

// NewMockCourseNoteRepositor creates a new mock instance
func NewMockCourseNoteRepositor(ctrl *gomock.Controller) *MockCourseNoteRepositor {
	mock := &MockCourseNoteRepositor{ctrl: ctrl}
	mock.recorder = &MockCourseNoteRepositorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCourseNoteRepositor) EXPECT() *MockCourseNoteRepositorMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCourseNoteRepositor) Get(ctx context.Context, cID, uID string) (*store.CourseNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, cID, uID)
	ret0, _ := ret[0].(*store.CourseNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCourseNoteRepositorMockRecorder) Get(ctx, cID, uID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCourseNoteRepositor)(nil).Get), ctx, cID, uID)
}

// Create mocks base method
func (m *MockCourseNoteRepositor) Create(ctx context.Context, note store.CourseNote) (*store.CourseNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, note)
	ret0, _ := ret[0].(*store.CourseNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCourseNoteRepositorMockRecorder) Create(ctx, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCourseNoteRepositor)(nil).Create), ctx, note)
}

// Update mocks base method
func (m *MockCourseNoteRepositor) Update(ctx context.Context, note store.CourseNote) (*store.CourseNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, note)
	ret0, _ := ret[0].(*store.CourseNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCourseNoteRepositorMockRecorder) Update(ctx, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCourseNoteRepositor)(nil).Update), ctx, note)
}
