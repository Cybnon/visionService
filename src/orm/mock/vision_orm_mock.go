// Code generated by MockGen. DO NOT EDIT.
// Source: vision_orm_interface.go

// Package mock_orm is a generated GoMock package.
package mock_orm

import (
	reflect "reflect"
	model "visionServiceGo/src/model"

	gomock "github.com/golang/mock/gomock"
)

// MockVisionORMModel is a mock of VisionORMModel interface.
type MockVisionORMModel struct {
	ctrl     *gomock.Controller
	recorder *MockVisionORMModelMockRecorder
}

// MockVisionORMModelMockRecorder is the mock recorder for MockVisionORMModel.
type MockVisionORMModelMockRecorder struct {
	mock *MockVisionORMModel
}

// NewMockVisionORMModel creates a new mock instance.
func NewMockVisionORMModel(ctrl *gomock.Controller) *MockVisionORMModel {
	mock := &MockVisionORMModel{ctrl: ctrl}
	mock.recorder = &MockVisionORMModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisionORMModel) EXPECT() *MockVisionORMModelMockRecorder {
	return m.recorder
}

// Activate mocks base method.
func (m *MockVisionORMModel) Activate(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activate", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Activate indicates an expected call of Activate.
func (mr *MockVisionORMModelMockRecorder) Activate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockVisionORMModel)(nil).Activate), id)
}

// Create mocks base method.
func (m *MockVisionORMModel) Create(vision model.Vision) (model.Vision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", vision)
	ret0, _ := ret[0].(model.Vision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVisionORMModelMockRecorder) Create(vision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVisionORMModel)(nil).Create), vision)
}

// Delete mocks base method.
func (m *MockVisionORMModel) Delete(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVisionORMModelMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVisionORMModel)(nil).Delete), id)
}

// GetActive mocks base method.
func (m *MockVisionORMModel) GetActive() (model.Vision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActive")
	ret0, _ := ret[0].(model.Vision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActive indicates an expected call of GetActive.
func (mr *MockVisionORMModelMockRecorder) GetActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActive", reflect.TypeOf((*MockVisionORMModel)(nil).GetActive))
}

// GetAll mocks base method.
func (m *MockVisionORMModel) GetAll() ([]model.Vision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Vision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVisionORMModelMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVisionORMModel)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockVisionORMModel) GetById(id uint64) (model.Vision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Vision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVisionORMModelMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVisionORMModel)(nil).GetById), id)
}

// IsActivated mocks base method.
func (m *MockVisionORMModel) IsActivated(id uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActivated", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActivated indicates an expected call of IsActivated.
func (mr *MockVisionORMModelMockRecorder) IsActivated(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActivated", reflect.TypeOf((*MockVisionORMModel)(nil).IsActivated), id)
}

// Update mocks base method.
func (m *MockVisionORMModel) Update(vision model.Vision) (model.Vision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", vision)
	ret0, _ := ret[0].(model.Vision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVisionORMModelMockRecorder) Update(vision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVisionORMModel)(nil).Update), vision)
}