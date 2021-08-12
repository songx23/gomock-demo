// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockdownstreamClient is a mock of downstreamClient interface.
type MockdownstreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockdownstreamClientMockRecorder
}

// MockdownstreamClientMockRecorder is the mock recorder for MockdownstreamClient.
type MockdownstreamClientMockRecorder struct {
	mock *MockdownstreamClient
}

// NewMockdownstreamClient creates a new mock instance.
func NewMockdownstreamClient(ctrl *gomock.Controller) *MockdownstreamClient {
	mock := &MockdownstreamClient{ctrl: ctrl}
	mock.recorder = &MockdownstreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdownstreamClient) EXPECT() *MockdownstreamClientMockRecorder {
	return m.recorder
}

// Magic mocks base method.
func (m *MockdownstreamClient) Magic(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Magic", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Magic indicates an expected call of Magic.
func (mr *MockdownstreamClientMockRecorder) Magic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Magic", reflect.TypeOf((*MockdownstreamClient)(nil).Magic), arg0)
}