// Code generated by MockGen. DO NOT EDIT.
// Source: gomock-showcase/internal/store (interfaces: Client,Villain)

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDarkMagicClient is a mock of Client interface.
type MockDarkMagicClient struct {
	ctrl     *gomock.Controller
	recorder *MockDarkMagicClientMockRecorder
}

// MockDarkMagicClientMockRecorder is the mock recorder for MockDarkMagicClient.
type MockDarkMagicClientMockRecorder struct {
	mock *MockDarkMagicClient
}

// NewMockDarkMagicClient creates a new mock instance.
func NewMockDarkMagicClient(ctrl *gomock.Controller) *MockDarkMagicClient {
	mock := &MockDarkMagicClient{ctrl: ctrl}
	mock.recorder = &MockDarkMagicClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDarkMagicClient) EXPECT() *MockDarkMagicClientMockRecorder {
	return m.recorder
}

// DarkMagic mocks base method.
func (m *MockDarkMagicClient) DarkMagic(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DarkMagic", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DarkMagic indicates an expected call of DarkMagic.
func (mr *MockDarkMagicClientMockRecorder) DarkMagic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DarkMagic", reflect.TypeOf((*MockDarkMagicClient)(nil).DarkMagic), arg0)
}

// MockDarkVillain is a mock of Villain interface.
type MockDarkVillain struct {
	ctrl     *gomock.Controller
	recorder *MockDarkVillainMockRecorder
}

// MockDarkVillainMockRecorder is the mock recorder for MockDarkVillain.
type MockDarkVillainMockRecorder struct {
	mock *MockDarkVillain
}

// NewMockDarkVillain creates a new mock instance.
func NewMockDarkVillain(ctrl *gomock.Controller) *MockDarkVillain {
	mock := &MockDarkVillain{ctrl: ctrl}
	mock.recorder = &MockDarkVillainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDarkVillain) EXPECT() *MockDarkVillainMockRecorder {
	return m.recorder
}

// Lose mocks base method.
func (m *MockDarkVillain) Lose() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lose")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Lose indicates an expected call of Lose.
func (mr *MockDarkVillainMockRecorder) Lose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lose", reflect.TypeOf((*MockDarkVillain)(nil).Lose))
}
