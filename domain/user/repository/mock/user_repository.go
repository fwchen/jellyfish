// Code generated by MockGen. DO NOT EDIT.
// Source: domain/user/repository/user_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	user "github.com/fwchen/jellyfish/domain/user"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockRepository) Save(user *user.AppUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockRepositoryMockRecorder) Save(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), user)
}

// Has mocks base method
func (m *MockRepository) Has(username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockRepositoryMockRecorder) Has(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRepository)(nil).Has), username)
}

// FindByID mocks base method
func (m *MockRepository) FindByID(userID string) (*user.AppUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", userID)
	ret0, _ := ret[0].(*user.AppUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockRepositoryMockRecorder) FindByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), userID)
}