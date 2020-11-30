// Code generated by MockGen. DO NOT EDIT.
// Source: factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/konstellation-io/kli/api"
	config "github.com/konstellation-io/kli/internal/config"
	logger "github.com/konstellation-io/kli/internal/logger"
	iostreams "github.com/konstellation-io/kli/pkg/iostreams"
	reflect "reflect"
)

// MockCmdFactory is a mock of CmdFactory interface
type MockCmdFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCmdFactoryMockRecorder
}

// MockCmdFactoryMockRecorder is the mock recorder for MockCmdFactory
type MockCmdFactoryMockRecorder struct {
	mock *MockCmdFactory
}

// NewMockCmdFactory creates a new mock instance
func NewMockCmdFactory(ctrl *gomock.Controller) *MockCmdFactory {
	mock := &MockCmdFactory{ctrl: ctrl}
	mock.recorder = &MockCmdFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCmdFactory) EXPECT() *MockCmdFactoryMockRecorder {
	return m.recorder
}

// IOStreams mocks base method
func (m *MockCmdFactory) IOStreams() *iostreams.IOStreams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IOStreams")
	ret0, _ := ret[0].(*iostreams.IOStreams)
	return ret0
}

// IOStreams indicates an expected call of IOStreams
func (mr *MockCmdFactoryMockRecorder) IOStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IOStreams", reflect.TypeOf((*MockCmdFactory)(nil).IOStreams))
}

// Config mocks base method
func (m *MockCmdFactory) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockCmdFactoryMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCmdFactory)(nil).Config))
}

// Logger mocks base method
func (m *MockCmdFactory) Logger() logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockCmdFactoryMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockCmdFactory)(nil).Logger))
}

// ServerClient mocks base method
func (m *MockCmdFactory) ServerClient(arg0 string) (api.ServerClienter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerClient", arg0)
	ret0, _ := ret[0].(api.ServerClienter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerClient indicates an expected call of ServerClient
func (mr *MockCmdFactoryMockRecorder) ServerClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerClient", reflect.TypeOf((*MockCmdFactory)(nil).ServerClient), arg0)
}