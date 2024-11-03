// Code generated by MockGen. DO NOT EDIT.
// Source: auth/logic.go
//
// Generated by this command:
//
//	mockgen -source=auth/logic.go -destination=auth/mock_logic.go -package=auth
//

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLogic is a mock of Logic interface.
type MockLogic struct {
	ctrl     *gomock.Controller
	recorder *MockLogicMockRecorder
	isgomock struct{}
}

// MockLogicMockRecorder is the mock recorder for MockLogic.
type MockLogicMockRecorder struct {
	mock *MockLogic
}

// NewMockLogic creates a new mock instance.
func NewMockLogic(ctrl *gomock.Controller) *MockLogic {
	mock := &MockLogic{ctrl: ctrl}
	mock.recorder = &MockLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogic) EXPECT() *MockLogicMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockLogic) Login(email, password string) (LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockLogicMockRecorder) Login(email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLogic)(nil).Login), email, password)
}