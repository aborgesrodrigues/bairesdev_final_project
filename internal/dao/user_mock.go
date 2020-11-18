// Code generated by MockGen. DO NOT EDIT.
// Source: bairesdev_final_project/internal/dao (interfaces: UserDAOInterface)

// Package dao is a generated GoMock package.
package dao

import (
	domain "bairesdev_final_project/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDAOInterface is a mock of UserDAOInterface interface
type MockUserDAOInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDAOInterfaceMockRecorder
}

// MockUserDAOInterfaceMockRecorder is the mock recorder for MockUserDAOInterface
type MockUserDAOInterfaceMockRecorder struct {
	mock *MockUserDAOInterface
}

// NewMockUserDAOInterface creates a new mock instance
func NewMockUserDAOInterface(ctrl *gomock.Controller) *MockUserDAOInterface {
	mock := &MockUserDAOInterface{ctrl: ctrl}
	mock.recorder = &MockUserDAOInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDAOInterface) EXPECT() *MockUserDAOInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserDAOInterface) Create(arg0 domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserDAOInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDAOInterface)(nil).Create), arg0)
}

// FindByID mocks base method
func (m *MockUserDAOInterface) FindByID(arg0 int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserDAOInterfaceMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserDAOInterface)(nil).FindByID), arg0)
}

// GetAll mocks base method
func (m *MockUserDAOInterface) GetAll() (*[]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockUserDAOInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserDAOInterface)(nil).GetAll))
}
