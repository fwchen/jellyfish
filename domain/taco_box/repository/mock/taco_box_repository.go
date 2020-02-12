// FileName generated by MockGen. DO NOT EDIT.
// Source: domain/taco_box/repository/taco_box_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	taco_box "github.com/fwchen/jellyfish/domain/taco_box"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTacoBoxRepository is a mock of TacoBoxRepository interface
type MockTacoBoxRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTacoBoxRepositoryMockRecorder
}

// MockTacoBoxRepositoryMockRecorder is the mock recorder for MockTacoBoxRepository
type MockTacoBoxRepositoryMockRecorder struct {
	mock *MockTacoBoxRepository
}

// NewMockTacoBoxRepository creates a new mock instance
func NewMockTacoBoxRepository(ctrl *gomock.Controller) *MockTacoBoxRepository {
	mock := &MockTacoBoxRepository{ctrl: ctrl}
	mock.recorder = &MockTacoBoxRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTacoBoxRepository) EXPECT() *MockTacoBoxRepositoryMockRecorder {
	return m.recorder
}

// SaveTacoBox mocks base method
func (m *MockTacoBoxRepository) SaveTacoBox(box *taco_box.TacoBox) (*taco_box.TacoBoxID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTacoBox", box)
	ret0, _ := ret[0].(*taco_box.TacoBoxID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTacoBox indicates an expected call of SaveTacoBox
func (mr *MockTacoBoxRepositoryMockRecorder) SaveTacoBox(box interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTacoBox", reflect.TypeOf((*MockTacoBoxRepository)(nil).SaveTacoBox), box)
}

// ListTacoBoxes mocks base method
func (m *MockTacoBoxRepository) ListTacoBoxes(userID string) ([]taco_box.TacoBox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTacoBoxes", userID)
	ret0, _ := ret[0].([]taco_box.TacoBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTacoBoxes indicates an expected call of ListTacoBoxes
func (mr *MockTacoBoxRepositoryMockRecorder) ListTacoBoxes(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTacoBoxes", reflect.TypeOf((*MockTacoBoxRepository)(nil).ListTacoBoxes), userID)
}

// FindTacoBox mocks base method
func (m *MockTacoBoxRepository) FindTacoBox(boxID string) (*taco_box.TacoBox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTacoBox", boxID)
	ret0, _ := ret[0].(*taco_box.TacoBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTacoBox indicates an expected call of FindTacoBox
func (mr *MockTacoBoxRepositoryMockRecorder) FindTacoBox(boxID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTacoBox", reflect.TypeOf((*MockTacoBoxRepository)(nil).FindTacoBox), boxID)
}
