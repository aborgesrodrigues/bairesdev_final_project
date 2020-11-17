// Code generated by MockGen. DO NOT EDIT.
// Source: question.go

// Package service is a generated GoMock package.
package service

import (
	domain "bairesdev_final_project/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockQuestionServiceInterface is a mock of QuestionServiceInterface interface
type MockQuestionServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionServiceInterfaceMockRecorder
}

// MockQuestionServiceInterfaceMockRecorder is the mock recorder for MockQuestionServiceInterface
type MockQuestionServiceInterfaceMockRecorder struct {
	mock *MockQuestionServiceInterface
}

// NewMockQuestionServiceInterface creates a new mock instance
func NewMockQuestionServiceInterface(ctrl *gomock.Controller) *MockQuestionServiceInterface {
	mock := &MockQuestionServiceInterface{ctrl: ctrl}
	mock.recorder = &MockQuestionServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuestionServiceInterface) EXPECT() *MockQuestionServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockQuestionServiceInterface) Create(arg0 domain.Question) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockQuestionServiceInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuestionServiceInterface)(nil).Create), arg0)
}

// Update mocks base method
func (m *MockQuestionServiceInterface) Update(arg0 domain.Question) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockQuestionServiceInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockQuestionServiceInterface)(nil).Update), arg0)
}

// Delete mocks base method
func (m *MockQuestionServiceInterface) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockQuestionServiceInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQuestionServiceInterface)(nil).Delete), arg0)
}

// FindByID mocks base method
func (m *MockQuestionServiceInterface) FindByID(arg0 int) (*domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockQuestionServiceInterfaceMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockQuestionServiceInterface)(nil).FindByID), arg0)
}

// FindByUser mocks base method
func (m *MockQuestionServiceInterface) FindByUser(arg0 int) (*[]domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0)
	ret0, _ := ret[0].(*[]domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockQuestionServiceInterfaceMockRecorder) FindByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockQuestionServiceInterface)(nil).FindByUser), arg0)
}

// GetAll mocks base method
func (m *MockQuestionServiceInterface) GetAll() (*[]domain.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]domain.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockQuestionServiceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockQuestionServiceInterface)(nil).GetAll))
}