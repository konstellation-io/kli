// Code generated by MockGen. DO NOT EDIT.
// Source: runtime.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	runtime "github.com/konstellation-io/kli/api/kre/runtime"
	reflect "reflect"
)

// MockRuntimeInterface is a mock of RuntimeInterface interface
type MockRuntimeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeInterfaceMockRecorder
}

// MockRuntimeInterfaceMockRecorder is the mock recorder for MockRuntimeInterface
type MockRuntimeInterfaceMockRecorder struct {
	mock *MockRuntimeInterface
}

// NewMockRuntimeInterface creates a new mock instance
func NewMockRuntimeInterface(ctrl *gomock.Controller) *MockRuntimeInterface {
	mock := &MockRuntimeInterface{ctrl: ctrl}
	mock.recorder = &MockRuntimeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuntimeInterface) EXPECT() *MockRuntimeInterfaceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockRuntimeInterface) List() (runtime.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(runtime.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRuntimeInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRuntimeInterface)(nil).List))
}
