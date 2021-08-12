// Code generated by MockGen. DO NOT EDIT.
// Source: villain.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVillain is a mock of Villain interface.
type MockVillain struct {
	ctrl     *gomock.Controller
	recorder *MockVillainMockRecorder
}

// MockVillainMockRecorder is the mock recorder for MockVillain.
type MockVillainMockRecorder struct {
	mock *MockVillain
}

// NewMockVillain creates a new mock instance.
func NewMockVillain(ctrl *gomock.Controller) *MockVillain {
	mock := &MockVillain{ctrl: ctrl}
	mock.recorder = &MockVillainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVillain) EXPECT() *MockVillainMockRecorder {
	return m.recorder
}

// Lose mocks base method.
func (m *MockVillain) Lose() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lose")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Lose indicates an expected call of Lose.
func (mr *MockVillainMockRecorder) Lose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lose", reflect.TypeOf((*MockVillain)(nil).Lose))
}
