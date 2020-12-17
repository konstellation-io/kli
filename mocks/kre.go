// Code generated by MockGen. DO NOT EDIT.
// Source: kre.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	runtime "github.com/konstellation-io/kli/api/kre/runtime"
	version "github.com/konstellation-io/kli/api/kre/version"
	reflect "reflect"
)

// MockKreInterface is a mock of KreInterface interface
type MockKreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKreInterfaceMockRecorder
}

// MockKreInterfaceMockRecorder is the mock recorder for MockKreInterface
type MockKreInterfaceMockRecorder struct {
	mock *MockKreInterface
}

// NewMockKreInterface creates a new mock instance
func NewMockKreInterface(ctrl *gomock.Controller) *MockKreInterface {
	mock := &MockKreInterface{ctrl: ctrl}
	mock.recorder = &MockKreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKreInterface) EXPECT() *MockKreInterfaceMockRecorder {
	return m.recorder
}

// Runtime mocks base method
func (m *MockKreInterface) Runtime() runtime.RuntimeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Runtime")
	ret0, _ := ret[0].(runtime.RuntimeInterface)
	return ret0
}

// Runtime indicates an expected call of Runtime
func (mr *MockKreInterfaceMockRecorder) Runtime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Runtime", reflect.TypeOf((*MockKreInterface)(nil).Runtime))
}

// Version mocks base method
func (m *MockKreInterface) Version() version.VersionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(version.VersionInterface)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockKreInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockKreInterface)(nil).Version))
}
