// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/oidc/pkg/op (interfaces: Configuration)

// Package mock is a generated GoMock package.
package mock

import (
	op "github.com/caos/oidc/pkg/op"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConfiguration is a mock of Configuration interface
type MockConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationMockRecorder
}

// MockConfigurationMockRecorder is the mock recorder for MockConfiguration
type MockConfigurationMockRecorder struct {
	mock *MockConfiguration
}

// NewMockConfiguration creates a new mock instance
func NewMockConfiguration(ctrl *gomock.Controller) *MockConfiguration {
	mock := &MockConfiguration{ctrl: ctrl}
	mock.recorder = &MockConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfiguration) EXPECT() *MockConfigurationMockRecorder {
	return m.recorder
}

// AuthMethodPostSupported mocks base method
func (m *MockConfiguration) AuthMethodPostSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMethodPostSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMethodPostSupported indicates an expected call of AuthMethodPostSupported
func (mr *MockConfigurationMockRecorder) AuthMethodPostSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMethodPostSupported", reflect.TypeOf((*MockConfiguration)(nil).AuthMethodPostSupported))
}

// AuthorizationEndpoint mocks base method
func (m *MockConfiguration) AuthorizationEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// AuthorizationEndpoint indicates an expected call of AuthorizationEndpoint
func (mr *MockConfigurationMockRecorder) AuthorizationEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationEndpoint", reflect.TypeOf((*MockConfiguration)(nil).AuthorizationEndpoint))
}

// CodeMethodS256Supported mocks base method
func (m *MockConfiguration) CodeMethodS256Supported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeMethodS256Supported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CodeMethodS256Supported indicates an expected call of CodeMethodS256Supported
func (mr *MockConfigurationMockRecorder) CodeMethodS256Supported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeMethodS256Supported", reflect.TypeOf((*MockConfiguration)(nil).CodeMethodS256Supported))
}

// EndSessionEndpoint mocks base method
func (m *MockConfiguration) EndSessionEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndSessionEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// EndSessionEndpoint indicates an expected call of EndSessionEndpoint
func (mr *MockConfigurationMockRecorder) EndSessionEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndSessionEndpoint", reflect.TypeOf((*MockConfiguration)(nil).EndSessionEndpoint))
}

// Issuer mocks base method
func (m *MockConfiguration) Issuer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Issuer")
	ret0, _ := ret[0].(string)
	return ret0
}

// Issuer indicates an expected call of Issuer
func (mr *MockConfigurationMockRecorder) Issuer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Issuer", reflect.TypeOf((*MockConfiguration)(nil).Issuer))
}

// KeysEndpoint mocks base method
func (m *MockConfiguration) KeysEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeysEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// KeysEndpoint indicates an expected call of KeysEndpoint
func (mr *MockConfigurationMockRecorder) KeysEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysEndpoint", reflect.TypeOf((*MockConfiguration)(nil).KeysEndpoint))
}

// TokenEndpoint mocks base method
func (m *MockConfiguration) TokenEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// TokenEndpoint indicates an expected call of TokenEndpoint
func (mr *MockConfigurationMockRecorder) TokenEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenEndpoint", reflect.TypeOf((*MockConfiguration)(nil).TokenEndpoint))
}

// UserinfoEndpoint mocks base method
func (m *MockConfiguration) UserinfoEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserinfoEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// UserinfoEndpoint indicates an expected call of UserinfoEndpoint
func (mr *MockConfigurationMockRecorder) UserinfoEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserinfoEndpoint", reflect.TypeOf((*MockConfiguration)(nil).UserinfoEndpoint))
}
