// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/domain/domain.go
//
// Generated by this command:
//
//	mockgen -source=internal/user/domain/domain.go -destination=internal/user/domain/mock_repo.go -package=domain
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

// CheckUserExists mocks base method.
func (m *MockRepo) CheckUserExists(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExists indicates an expected call of CheckUserExists.
func (mr *MockRepoMockRecorder) CheckUserExists(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockRepo)(nil).CheckUserExists), email)
}

// FindProfileByID mocks base method.
func (m *MockRepo) FindProfileByID(profileID int) (Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfileByID", profileID)
	ret0, _ := ret[0].(Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProfileByID indicates an expected call of FindProfileByID.
func (mr *MockRepoMockRecorder) FindProfileByID(profileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfileByID", reflect.TypeOf((*MockRepo)(nil).FindProfileByID), profileID)
}

// FindUserByID mocks base method.
func (m *MockRepo) FindUserByID(userID int) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", userID)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockRepoMockRecorder) FindUserByID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockRepo)(nil).FindUserByID), userID)
}

// InsertUser mocks base method.
func (m *MockRepo) InsertUser(u User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", u)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockRepoMockRecorder) InsertUser(u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockRepo)(nil).InsertUser), u)
}

// UpdateProfile mocks base method.
func (m *MockRepo) UpdateProfile(profileId int, profile Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", profileId, profile)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockRepoMockRecorder) UpdateProfile(profileId, profile any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockRepo)(nil).UpdateProfile), profileId, profile)
}

// UpdateUser mocks base method.
func (m *MockRepo) UpdateUser(userId int, user User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepoMockRecorder) UpdateUser(userId, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepo)(nil).UpdateUser), userId, user)
}
