// Code generated by MockGen. DO NOT EDIT.
// Source: auth/domain.go
//
// Generated by this command:
//
//	mockgen -source=auth/domain.go -destination=auth/mock_repo.go -package=auth
//

// Package auth is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
	isgomock struct{}
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// FindProfileByUserID mocks base method.
func (m *MockRepo) FindProfileByUserID(userID int) (Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfileByUserID", userID)
	ret0, _ := ret[0].(Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProfileByUserID indicates an expected call of FindProfileByUserID.
func (mr *MockRepoMockRecorder) FindProfileByUserID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfileByUserID", reflect.TypeOf((*MockRepo)(nil).FindProfileByUserID), userID)
}

// FindUserByEmail mocks base method.
func (m *MockRepo) FindUserByEmail(email string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockRepoMockRecorder) FindUserByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockRepo)(nil).FindUserByEmail), email)
}