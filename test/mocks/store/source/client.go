// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DarkMagic mocks base method.
func (m *MockClient) DarkMagic(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DarkMagic", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DarkMagic indicates an expected call of DarkMagic.
func (mr *MockClientMockRecorder) DarkMagic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DarkMagic", reflect.TypeOf((*MockClient)(nil).DarkMagic), arg0)
}
