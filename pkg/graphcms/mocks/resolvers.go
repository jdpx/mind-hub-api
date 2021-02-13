// Code generated by MockGen. DO NOT EDIT.
// Source: resolvers.go

// Package graphcmsmocks is a generated GoMock package.
package graphcmsmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	graphcms "github.com/jdpx/mind-hub-api/pkg/graphcms"
	reflect "reflect"
)

// MockCMSResolver is a mock of CMSResolver interface
type MockCMSResolver struct {
	ctrl     *gomock.Controller
	recorder *MockCMSResolverMockRecorder
}

// MockCMSResolverMockRecorder is the mock recorder for MockCMSResolver
type MockCMSResolverMockRecorder struct {
	mock *MockCMSResolver
}

// NewMockCMSResolver creates a new mock instance
func NewMockCMSResolver(ctrl *gomock.Controller) *MockCMSResolver {
	mock := &MockCMSResolver{ctrl: ctrl}
	mock.recorder = &MockCMSResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCMSResolver) EXPECT() *MockCMSResolverMockRecorder {
	return m.recorder
}

// GetCourses mocks base method
func (m *MockCMSResolver) GetCourses(ctx context.Context) ([]*graphcms.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourses", ctx)
	ret0, _ := ret[0].([]*graphcms.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourses indicates an expected call of GetCourses
func (mr *MockCMSResolverMockRecorder) GetCourses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourses", reflect.TypeOf((*MockCMSResolver)(nil).GetCourses), ctx)
}

// GetCourseByID mocks base method
func (m *MockCMSResolver) GetCourseByID(ctx context.Context, id string) (*graphcms.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseByID", ctx, id)
	ret0, _ := ret[0].(*graphcms.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseByID indicates an expected call of GetCourseByID
func (mr *MockCMSResolverMockRecorder) GetCourseByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseByID", reflect.TypeOf((*MockCMSResolver)(nil).GetCourseByID), ctx, id)
}

// GetSessionsByCourseID mocks base method
func (m *MockCMSResolver) GetSessionsByCourseID(ctx context.Context, id string) ([]*graphcms.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionsByCourseID", ctx, id)
	ret0, _ := ret[0].([]*graphcms.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionsByCourseID indicates an expected call of GetSessionsByCourseID
func (mr *MockCMSResolverMockRecorder) GetSessionsByCourseID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionsByCourseID", reflect.TypeOf((*MockCMSResolver)(nil).GetSessionsByCourseID), ctx, id)
}

// GetSessionByID mocks base method
func (m *MockCMSResolver) GetSessionByID(ctx context.Context, id string) (*graphcms.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", ctx, id)
	ret0, _ := ret[0].(*graphcms.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID
func (mr *MockCMSResolverMockRecorder) GetSessionByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockCMSResolver)(nil).GetSessionByID), ctx, id)
}

// GetStepIDsByCourseID mocks base method
func (m *MockCMSResolver) GetStepIDsByCourseID(ctx context.Context, id string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepIDsByCourseID", ctx, id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepIDsByCourseID indicates an expected call of GetStepIDsByCourseID
func (mr *MockCMSResolverMockRecorder) GetStepIDsByCourseID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepIDsByCourseID", reflect.TypeOf((*MockCMSResolver)(nil).GetStepIDsByCourseID), ctx, id)
}

// GetStepsByID mocks base method
func (m *MockCMSResolver) GetStepsByID(ctx context.Context, id string) (*graphcms.Step, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepsBySessionID", ctx, id)
	ret0, _ := ret[0].(*graphcms.Step)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepsByID indicates an expected call of GetStepsByID
func (mr *MockCMSResolverMockRecorder) GetStepsByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepsBySessionID", reflect.TypeOf((*MockCMSResolver)(nil).GetStepsByID), ctx, id)
}
