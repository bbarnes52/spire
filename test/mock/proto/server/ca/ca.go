// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_ca is a generated GoMock package.
package mock_ca

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	ca "github.com/spiffe/spire/proto/server/ca"
)

// MockServerCa is a mock of ServerCa interface
type MockServerCa struct {
	ctrl     *gomock.Controller
	recorder *MockServerCaMockRecorder
}

// MockServerCaMockRecorder is the mock recorder for MockServerCa
type MockServerCaMockRecorder struct {
	mock *MockServerCa
}

// NewMockServerCa creates a new mock instance
func NewMockServerCa(ctrl *gomock.Controller) *MockServerCa {
	mock := &MockServerCa{ctrl: ctrl}
	mock.recorder = &MockServerCaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerCa) EXPECT() *MockServerCaMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockServerCa) Configure(request *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", request)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockServerCaMockRecorder) Configure(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockServerCa)(nil).Configure), request)
}

// GetPluginInfo mocks base method
func (m *MockServerCa) GetPluginInfo(arg0 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockServerCaMockRecorder) GetPluginInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockServerCa)(nil).GetPluginInfo), arg0)
}

// SignCsr mocks base method
func (m *MockServerCa) SignCsr(arg0 *ca.SignCsrRequest) (*ca.SignCsrResponse, error) {
	ret := m.ctrl.Call(m, "SignCsr", arg0)
	ret0, _ := ret[0].(*ca.SignCsrResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignCsr indicates an expected call of SignCsr
func (mr *MockServerCaMockRecorder) SignCsr(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignCsr", reflect.TypeOf((*MockServerCa)(nil).SignCsr), arg0)
}

// GenerateCsr mocks base method
func (m *MockServerCa) GenerateCsr(arg0 *ca.GenerateCsrRequest) (*ca.GenerateCsrResponse, error) {
	ret := m.ctrl.Call(m, "GenerateCsr", arg0)
	ret0, _ := ret[0].(*ca.GenerateCsrResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCsr indicates an expected call of GenerateCsr
func (mr *MockServerCaMockRecorder) GenerateCsr(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCsr", reflect.TypeOf((*MockServerCa)(nil).GenerateCsr), arg0)
}

// FetchCertificate mocks base method
func (m *MockServerCa) FetchCertificate(request *ca.FetchCertificateRequest) (*ca.FetchCertificateResponse, error) {
	ret := m.ctrl.Call(m, "FetchCertificate", request)
	ret0, _ := ret[0].(*ca.FetchCertificateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCertificate indicates an expected call of FetchCertificate
func (mr *MockServerCaMockRecorder) FetchCertificate(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCertificate", reflect.TypeOf((*MockServerCa)(nil).FetchCertificate), request)
}

// LoadCertificate mocks base method
func (m *MockServerCa) LoadCertificate(arg0 *ca.LoadCertificateRequest) (*ca.LoadCertificateResponse, error) {
	ret := m.ctrl.Call(m, "LoadCertificate", arg0)
	ret0, _ := ret[0].(*ca.LoadCertificateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadCertificate indicates an expected call of LoadCertificate
func (mr *MockServerCaMockRecorder) LoadCertificate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadCertificate", reflect.TypeOf((*MockServerCa)(nil).LoadCertificate), arg0)
}
