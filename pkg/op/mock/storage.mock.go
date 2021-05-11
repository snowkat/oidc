// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/oidc/pkg/op (interfaces: Storage)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	oidc "github.com/caos/oidc/pkg/oidc"
	op "github.com/caos/oidc/pkg/op"
	gomock "github.com/golang/mock/gomock"
	jose "gopkg.in/square/go-jose.v2"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AuthRequestByCode mocks base method.
func (m *MockStorage) AuthRequestByCode(arg0 context.Context, arg1 string) (op.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthRequestByCode", arg0, arg1)
	ret0, _ := ret[0].(op.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthRequestByCode indicates an expected call of AuthRequestByCode.
func (mr *MockStorageMockRecorder) AuthRequestByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthRequestByCode", reflect.TypeOf((*MockStorage)(nil).AuthRequestByCode), arg0, arg1)
}

// AuthRequestByID mocks base method.
func (m *MockStorage) AuthRequestByID(arg0 context.Context, arg1 string) (op.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthRequestByID", arg0, arg1)
	ret0, _ := ret[0].(op.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthRequestByID indicates an expected call of AuthRequestByID.
func (mr *MockStorageMockRecorder) AuthRequestByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthRequestByID", reflect.TypeOf((*MockStorage)(nil).AuthRequestByID), arg0, arg1)
}

// AuthorizeClientIDSecret mocks base method.
func (m *MockStorage) AuthorizeClientIDSecret(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeClientIDSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthorizeClientIDSecret indicates an expected call of AuthorizeClientIDSecret.
func (mr *MockStorageMockRecorder) AuthorizeClientIDSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeClientIDSecret", reflect.TypeOf((*MockStorage)(nil).AuthorizeClientIDSecret), arg0, arg1, arg2)
}

// CreateAccessAndRefreshTokens mocks base method.
func (m *MockStorage) CreateAccessAndRefreshTokens(arg0 context.Context, arg1 op.TokenRequest, arg2 string) (string, string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessAndRefreshTokens", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAccessAndRefreshTokens indicates an expected call of CreateAccessAndRefreshTokens.
func (mr *MockStorageMockRecorder) CreateAccessAndRefreshTokens(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessAndRefreshTokens", reflect.TypeOf((*MockStorage)(nil).CreateAccessAndRefreshTokens), arg0, arg1, arg2)
}

// CreateAccessToken mocks base method.
func (m *MockStorage) CreateAccessToken(arg0 context.Context, arg1 op.TokenRequest) (string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockStorageMockRecorder) CreateAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockStorage)(nil).CreateAccessToken), arg0, arg1)
}

// CreateAuthRequest mocks base method.
func (m *MockStorage) CreateAuthRequest(arg0 context.Context, arg1 *oidc.AuthRequest, arg2 string) (op.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(op.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthRequest indicates an expected call of CreateAuthRequest.
func (mr *MockStorageMockRecorder) CreateAuthRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthRequest", reflect.TypeOf((*MockStorage)(nil).CreateAuthRequest), arg0, arg1, arg2)
}

// DeleteAuthRequest mocks base method.
func (m *MockStorage) DeleteAuthRequest(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthRequest indicates an expected call of DeleteAuthRequest.
func (mr *MockStorageMockRecorder) DeleteAuthRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthRequest", reflect.TypeOf((*MockStorage)(nil).DeleteAuthRequest), arg0, arg1)
}

// GetClientByClientID mocks base method.
func (m *MockStorage) GetClientByClientID(arg0 context.Context, arg1 string) (op.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientByClientID", arg0, arg1)
	ret0, _ := ret[0].(op.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientByClientID indicates an expected call of GetClientByClientID.
func (mr *MockStorageMockRecorder) GetClientByClientID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientByClientID", reflect.TypeOf((*MockStorage)(nil).GetClientByClientID), arg0, arg1)
}

// GetKeyByIDAndUserID mocks base method.
func (m *MockStorage) GetKeyByIDAndUserID(arg0 context.Context, arg1, arg2 string) (*jose.JSONWebKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyByIDAndUserID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jose.JSONWebKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyByIDAndUserID indicates an expected call of GetKeyByIDAndUserID.
func (mr *MockStorageMockRecorder) GetKeyByIDAndUserID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyByIDAndUserID", reflect.TypeOf((*MockStorage)(nil).GetKeyByIDAndUserID), arg0, arg1, arg2)
}

// GetKeySet mocks base method.
func (m *MockStorage) GetKeySet(arg0 context.Context) (*jose.JSONWebKeySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeySet", arg0)
	ret0, _ := ret[0].(*jose.JSONWebKeySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeySet indicates an expected call of GetKeySet.
func (mr *MockStorageMockRecorder) GetKeySet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeySet", reflect.TypeOf((*MockStorage)(nil).GetKeySet), arg0)
}

// GetPrivateClaimsFromScopes mocks base method.
func (m *MockStorage) GetPrivateClaimsFromScopes(arg0 context.Context, arg1, arg2 string, arg3 []string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateClaimsFromScopes", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateClaimsFromScopes indicates an expected call of GetPrivateClaimsFromScopes.
func (mr *MockStorageMockRecorder) GetPrivateClaimsFromScopes(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateClaimsFromScopes", reflect.TypeOf((*MockStorage)(nil).GetPrivateClaimsFromScopes), arg0, arg1, arg2, arg3)
}

// GetSigningKey mocks base method.
func (m *MockStorage) GetSigningKey(arg0 context.Context, arg1 chan<- jose.SigningKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetSigningKey", arg0, arg1)
}

// GetSigningKey indicates an expected call of GetSigningKey.
func (mr *MockStorageMockRecorder) GetSigningKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigningKey", reflect.TypeOf((*MockStorage)(nil).GetSigningKey), arg0, arg1)
}

// Health mocks base method.
func (m *MockStorage) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockStorageMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockStorage)(nil).Health), arg0)
}

// SaveAuthCode mocks base method.
func (m *MockStorage) SaveAuthCode(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAuthCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAuthCode indicates an expected call of SaveAuthCode.
func (mr *MockStorageMockRecorder) SaveAuthCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAuthCode", reflect.TypeOf((*MockStorage)(nil).SaveAuthCode), arg0, arg1, arg2)
}

// SetIntrospectionFromToken mocks base method.
func (m *MockStorage) SetIntrospectionFromToken(arg0 context.Context, arg1 oidc.IntrospectionResponse, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIntrospectionFromToken", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIntrospectionFromToken indicates an expected call of SetIntrospectionFromToken.
func (mr *MockStorageMockRecorder) SetIntrospectionFromToken(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIntrospectionFromToken", reflect.TypeOf((*MockStorage)(nil).SetIntrospectionFromToken), arg0, arg1, arg2, arg3, arg4)
}

// SetUserinfoFromScopes mocks base method.
func (m *MockStorage) SetUserinfoFromScopes(arg0 context.Context, arg1 oidc.UserInfoSetter, arg2, arg3 string, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserinfoFromScopes", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserinfoFromScopes indicates an expected call of SetUserinfoFromScopes.
func (mr *MockStorageMockRecorder) SetUserinfoFromScopes(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserinfoFromScopes", reflect.TypeOf((*MockStorage)(nil).SetUserinfoFromScopes), arg0, arg1, arg2, arg3, arg4)
}

// SetUserinfoFromToken mocks base method.
func (m *MockStorage) SetUserinfoFromToken(arg0 context.Context, arg1 oidc.UserInfoSetter, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserinfoFromToken", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserinfoFromToken indicates an expected call of SetUserinfoFromToken.
func (mr *MockStorageMockRecorder) SetUserinfoFromToken(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserinfoFromToken", reflect.TypeOf((*MockStorage)(nil).SetUserinfoFromToken), arg0, arg1, arg2, arg3, arg4)
}

// TerminateSession mocks base method.
func (m *MockStorage) TerminateSession(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateSession indicates an expected call of TerminateSession.
func (mr *MockStorageMockRecorder) TerminateSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateSession", reflect.TypeOf((*MockStorage)(nil).TerminateSession), arg0, arg1, arg2)
}

// TokenRequestByRefreshToken mocks base method.
func (m *MockStorage) TokenRequestByRefreshToken(arg0 context.Context, arg1 string) (op.RefreshTokenRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenRequestByRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(op.RefreshTokenRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenRequestByRefreshToken indicates an expected call of TokenRequestByRefreshToken.
func (mr *MockStorageMockRecorder) TokenRequestByRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenRequestByRefreshToken", reflect.TypeOf((*MockStorage)(nil).TokenRequestByRefreshToken), arg0, arg1)
}

// ValidateJWTProfileScopes mocks base method.
func (m *MockStorage) ValidateJWTProfileScopes(arg0 context.Context, arg1 string, arg2 oidc.Scopes) (oidc.Scopes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateJWTProfileScopes", arg0, arg1, arg2)
	ret0, _ := ret[0].(oidc.Scopes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWTProfileScopes indicates an expected call of ValidateJWTProfileScopes.
func (mr *MockStorageMockRecorder) ValidateJWTProfileScopes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWTProfileScopes", reflect.TypeOf((*MockStorage)(nil).ValidateJWTProfileScopes), arg0, arg1, arg2)
}
