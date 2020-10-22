// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/studyzy/fabric-sdk-go/pkg/common/providers/context (interfaces: Providers,Client)

// Package mockcontext is a generated GoMock package.
package mockcontext

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/studyzy/fabric-sdk-go/pkg/common/providers/core"
	fab "github.com/studyzy/fabric-sdk-go/pkg/common/providers/fab"
	msp "github.com/studyzy/fabric-sdk-go/pkg/common/providers/msp"
	metrics "github.com/studyzy/fabric-sdk-go/pkg/fabsdk/metrics"
)

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockProviders) ChannelProvider() fab.ChannelProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(fab.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockProvidersMockRecorder) ChannelProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockProviders)(nil).ChannelProvider))
}

// CryptoSuite mocks base method
func (m *MockProviders) CryptoSuite() core.CryptoSuite {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockProvidersMockRecorder) CryptoSuite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockProviders)(nil).CryptoSuite))
}

// EndpointConfig mocks base method
func (m *MockProviders) EndpointConfig() fab.EndpointConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointConfig")
	ret0, _ := ret[0].(fab.EndpointConfig)
	return ret0
}

// EndpointConfig indicates an expected call of EndpointConfig
func (mr *MockProvidersMockRecorder) EndpointConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointConfig", reflect.TypeOf((*MockProviders)(nil).EndpointConfig))
}

// GetMetrics mocks base method
func (m *MockProviders) GetMetrics() *metrics.ClientMetrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics")
	ret0, _ := ret[0].(*metrics.ClientMetrics)
	return ret0
}

// GetMetrics indicates an expected call of GetMetrics
func (mr *MockProvidersMockRecorder) GetMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockProviders)(nil).GetMetrics))
}

// IdentityConfig mocks base method
func (m *MockProviders) IdentityConfig() msp.IdentityConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityConfig")
	ret0, _ := ret[0].(msp.IdentityConfig)
	return ret0
}

// IdentityConfig indicates an expected call of IdentityConfig
func (mr *MockProvidersMockRecorder) IdentityConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityConfig", reflect.TypeOf((*MockProviders)(nil).IdentityConfig))
}

// IdentityManager mocks base method
func (m *MockProviders) IdentityManager(arg0 string) (msp.IdentityManager, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityManager", arg0)
	ret0, _ := ret[0].(msp.IdentityManager)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IdentityManager indicates an expected call of IdentityManager
func (mr *MockProvidersMockRecorder) IdentityManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityManager", reflect.TypeOf((*MockProviders)(nil).IdentityManager), arg0)
}

// InfraProvider mocks base method
func (m *MockProviders) InfraProvider() fab.InfraProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfraProvider")
	ret0, _ := ret[0].(fab.InfraProvider)
	return ret0
}

// InfraProvider indicates an expected call of InfraProvider
func (mr *MockProvidersMockRecorder) InfraProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfraProvider", reflect.TypeOf((*MockProviders)(nil).InfraProvider))
}

// LocalDiscoveryProvider mocks base method
func (m *MockProviders) LocalDiscoveryProvider() fab.LocalDiscoveryProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalDiscoveryProvider")
	ret0, _ := ret[0].(fab.LocalDiscoveryProvider)
	return ret0
}

// LocalDiscoveryProvider indicates an expected call of LocalDiscoveryProvider
func (mr *MockProvidersMockRecorder) LocalDiscoveryProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalDiscoveryProvider", reflect.TypeOf((*MockProviders)(nil).LocalDiscoveryProvider))
}

// SigningManager mocks base method
func (m *MockProviders) SigningManager() core.SigningManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(core.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockProvidersMockRecorder) SigningManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockProviders)(nil).SigningManager))
}

// UserStore mocks base method
func (m *MockProviders) UserStore() msp.UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserStore")
	ret0, _ := ret[0].(msp.UserStore)
	return ret0
}

// UserStore indicates an expected call of UserStore
func (mr *MockProvidersMockRecorder) UserStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserStore", reflect.TypeOf((*MockProviders)(nil).UserStore))
}

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockClient) ChannelProvider() fab.ChannelProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(fab.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockClientMockRecorder) ChannelProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockClient)(nil).ChannelProvider))
}

// CryptoSuite mocks base method
func (m *MockClient) CryptoSuite() core.CryptoSuite {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockClientMockRecorder) CryptoSuite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockClient)(nil).CryptoSuite))
}

// EndpointConfig mocks base method
func (m *MockClient) EndpointConfig() fab.EndpointConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointConfig")
	ret0, _ := ret[0].(fab.EndpointConfig)
	return ret0
}

// EndpointConfig indicates an expected call of EndpointConfig
func (mr *MockClientMockRecorder) EndpointConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointConfig", reflect.TypeOf((*MockClient)(nil).EndpointConfig))
}

// EnrollmentCertificate mocks base method
func (m *MockClient) EnrollmentCertificate() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrollmentCertificate")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// EnrollmentCertificate indicates an expected call of EnrollmentCertificate
func (mr *MockClientMockRecorder) EnrollmentCertificate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrollmentCertificate", reflect.TypeOf((*MockClient)(nil).EnrollmentCertificate))
}

// GetMetrics mocks base method
func (m *MockClient) GetMetrics() *metrics.ClientMetrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics")
	ret0, _ := ret[0].(*metrics.ClientMetrics)
	return ret0
}

// GetMetrics indicates an expected call of GetMetrics
func (mr *MockClientMockRecorder) GetMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockClient)(nil).GetMetrics))
}

// Identifier mocks base method
func (m *MockClient) Identifier() *msp.IdentityIdentifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(*msp.IdentityIdentifier)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockClientMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockClient)(nil).Identifier))
}

// IdentityConfig mocks base method
func (m *MockClient) IdentityConfig() msp.IdentityConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityConfig")
	ret0, _ := ret[0].(msp.IdentityConfig)
	return ret0
}

// IdentityConfig indicates an expected call of IdentityConfig
func (mr *MockClientMockRecorder) IdentityConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityConfig", reflect.TypeOf((*MockClient)(nil).IdentityConfig))
}

// IdentityManager mocks base method
func (m *MockClient) IdentityManager(arg0 string) (msp.IdentityManager, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentityManager", arg0)
	ret0, _ := ret[0].(msp.IdentityManager)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IdentityManager indicates an expected call of IdentityManager
func (mr *MockClientMockRecorder) IdentityManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentityManager", reflect.TypeOf((*MockClient)(nil).IdentityManager), arg0)
}

// InfraProvider mocks base method
func (m *MockClient) InfraProvider() fab.InfraProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfraProvider")
	ret0, _ := ret[0].(fab.InfraProvider)
	return ret0
}

// InfraProvider indicates an expected call of InfraProvider
func (mr *MockClientMockRecorder) InfraProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfraProvider", reflect.TypeOf((*MockClient)(nil).InfraProvider))
}

// LocalDiscoveryProvider mocks base method
func (m *MockClient) LocalDiscoveryProvider() fab.LocalDiscoveryProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalDiscoveryProvider")
	ret0, _ := ret[0].(fab.LocalDiscoveryProvider)
	return ret0
}

// LocalDiscoveryProvider indicates an expected call of LocalDiscoveryProvider
func (mr *MockClientMockRecorder) LocalDiscoveryProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalDiscoveryProvider", reflect.TypeOf((*MockClient)(nil).LocalDiscoveryProvider))
}

// PrivateKey mocks base method
func (m *MockClient) PrivateKey() core.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateKey")
	ret0, _ := ret[0].(core.Key)
	return ret0
}

// PrivateKey indicates an expected call of PrivateKey
func (mr *MockClientMockRecorder) PrivateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateKey", reflect.TypeOf((*MockClient)(nil).PrivateKey))
}

// PublicVersion mocks base method
func (m *MockClient) PublicVersion() msp.Identity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicVersion")
	ret0, _ := ret[0].(msp.Identity)
	return ret0
}

// PublicVersion indicates an expected call of PublicVersion
func (mr *MockClientMockRecorder) PublicVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicVersion", reflect.TypeOf((*MockClient)(nil).PublicVersion))
}

// Serialize mocks base method
func (m *MockClient) Serialize() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Serialize indicates an expected call of Serialize
func (mr *MockClientMockRecorder) Serialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockClient)(nil).Serialize))
}

// Sign mocks base method
func (m *MockClient) Sign(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockClientMockRecorder) Sign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockClient)(nil).Sign), arg0)
}

// SigningManager mocks base method
func (m *MockClient) SigningManager() core.SigningManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(core.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockClientMockRecorder) SigningManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockClient)(nil).SigningManager))
}

// UserStore mocks base method
func (m *MockClient) UserStore() msp.UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserStore")
	ret0, _ := ret[0].(msp.UserStore)
	return ret0
}

// UserStore indicates an expected call of UserStore
func (mr *MockClientMockRecorder) UserStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserStore", reflect.TypeOf((*MockClient)(nil).UserStore))
}

// Verify mocks base method
func (m *MockClient) Verify(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockClientMockRecorder) Verify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockClient)(nil).Verify), arg0, arg1)
}
