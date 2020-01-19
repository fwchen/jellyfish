// Code generated by MockGen. DO NOT EDIT.
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

// InsertTacoBox mocks base method
func (m *MockTacoBoxRepository) InsertTacoBox(box *taco_box.TacoBox) (*taco_box.TacoBoxID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTacoBox", box)
	ret0, _ := ret[0].(*taco_box.TacoBoxID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTacoBox indicates an expected call of InsertTacoBox
func (mr *MockTacoBoxRepositoryMockRecorder) InsertTacoBox(box interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTacoBox", reflect.TypeOf((*MockTacoBoxRepository)(nil).InsertTacoBox), box)
}
