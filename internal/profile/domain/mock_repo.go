// Code generated by MockGen. DO NOT EDIT.
// Source: internal/profile/domain/domain.go
//
// Generated by this command:
//
//	mockgen -source=internal/profile/domain/domain.go -destination=internal/profile/domain/mock_repo.go -package=domain
//

// Package domain is a generated GoMock package.
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

// FindProfile mocks base method.
func (m *MockRepo) FindProfile(profileID int) (Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfile", profileID)
	ret0, _ := ret[0].(Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProfile indicates an expected call of FindProfile.
func (mr *MockRepoMockRecorder) FindProfile(profileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfile", reflect.TypeOf((*MockRepo)(nil).FindProfile), profileID)
}

// FindProfileByUsername mocks base method.
func (m *MockRepo) FindProfileByUsername(username string) (Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfileByUsername", username)
	ret0, _ := ret[0].(Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProfileByUsername indicates an expected call of FindProfileByUsername.
func (mr *MockRepoMockRecorder) FindProfileByUsername(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfileByUsername", reflect.TypeOf((*MockRepo)(nil).FindProfileByUsername), username)
}

// FindProfileWithFollowingByUsername mocks base method.
func (m *MockRepo) FindProfileWithFollowingByUsername(username string, currentUserId int) (Profile, Following, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfileWithFollowingByUsername", username, currentUserId)
	ret0, _ := ret[0].(Profile)
	ret1, _ := ret[1].(Following)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindProfileWithFollowingByUsername indicates an expected call of FindProfileWithFollowingByUsername.
func (mr *MockRepoMockRecorder) FindProfileWithFollowingByUsername(username, currentUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfileWithFollowingByUsername", reflect.TypeOf((*MockRepo)(nil).FindProfileWithFollowingByUsername), username, currentUserId)
}