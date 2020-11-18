// Code generated by MockGen. DO NOT EDIT.
// Source: question.go

// Package dao is a generated GoMock package.
package dao

import (
	domain "bairesdev_final_project/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockQuestionDAOInterface is a mock of QuestionDAOInterface interface
type MockQuestionDAOInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionDAOInterfaceMockRecorder
}

// MockQuestionDAOInterfaceMockRecorder is the mock recorder for MockQuestionDAOInterface
type MockQuestionDAOInterfaceMockRecorder struct {
	mock *MockQuestionDAOInterface
}

// NewMockQuestionDAOInterface creates a new mock instance
func NewMockQuestionDAOInterface(ctrl *gomock.Controller) *MockQuestionDAOInterface {
	mock := &MockQuestionDAOInterface{ctrl: ctrl}
	mock.recorder = &MockQuestionDAOInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuestionDAOInterface) EXPECT() *MockQuestionDAOInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockQuestionDAOInterface) Create(arg0 domain.Question) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockQuestionDAOInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuestionDAOInterface)(nil).Create), arg0)
}

// Update mocks base method
func (m *MockQuestionDAOInterface) Update(arg0 domain.Question) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockQuestionDAOInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockQuestionDAOInterface)(nil).Update), arg0)
}

// Delete mocks base method
func (m *MockQuestionDAOInterface) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockQuestionDAOInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQuestionDAOInterface)(nil).Delete), arg0)
}

// FindByID mocks base method
func (m *MockQuestionDAOInterface) FindByID(arg0 int) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockQuestionDAOInterfaceMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockQuestionDAOInterface)(nil).FindByID), arg0)
}

// FindByUser mocks base method
func (m *MockQuestionDAOInterface) FindByUser(arg0 int) (*[]domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0)
	ret0, _ := ret[0].(*[]domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockQuestionDAOInterfaceMockRecorder) FindByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockQuestionDAOInterface)(nil).FindByUser), arg0)
}

// GetAll mocks base method
func (m *MockQuestionDAOInterface) GetAll() (*[]domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockQuestionDAOInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockQuestionDAOInterface)(nil).GetAll))
}
