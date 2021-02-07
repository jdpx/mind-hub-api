// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package storemocks is a generated GoMock package.
package storemocks

import (
	context "context"
	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorer is a mock of Storer interface
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStorer) Get(ctx context.Context, tableName, pk, sk string, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, tableName, pk, sk, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockStorerMockRecorder) Get(ctx, tableName, pk, sk, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorer)(nil).Get), ctx, tableName, pk, sk, i)
}

// BatchGet mocks base method
func (m *MockStorer) BatchGet(ctx context.Context, tableName, pk string, sk []string, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGet", ctx, tableName, pk, sk, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchGet indicates an expected call of BatchGet
func (mr *MockStorerMockRecorder) BatchGet(ctx, tableName, pk, sk, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGet", reflect.TypeOf((*MockStorer)(nil).BatchGet), ctx, tableName, pk, sk, i)
}

// Query mocks base method
func (m *MockStorer) Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, tableName, ex, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockStorerMockRecorder) Query(ctx, tableName, ex, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockStorer)(nil).Query), ctx, tableName, ex, i)
}

// Put mocks base method
func (m *MockStorer) Put(ctx context.Context, tableName string, body interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, tableName, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockStorerMockRecorder) Put(ctx, tableName, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStorer)(nil).Put), ctx, tableName, body)
}

// Update mocks base method
func (m *MockStorer) Update(ctx context.Context, tableName, pk, sk string, ex expression.Expression, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tableName, pk, sk, ex, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStorerMockRecorder) Update(ctx, tableName, pk, sk, ex, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStorer)(nil).Update), ctx, tableName, pk, sk, ex, i)
}
