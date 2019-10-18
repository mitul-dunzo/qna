// Code generated by MockGen. DO NOT EDIT.
// Source: qna/main/services (interfaces: IJwtService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIJwtService is a mock of IJwtService interface
type MockIJwtService struct {
	ctrl     *gomock.Controller
	recorder *MockIJwtServiceMockRecorder
}

// MockIJwtServiceMockRecorder is the mock recorder for MockIJwtService
type MockIJwtServiceMockRecorder struct {
	mock *MockIJwtService
}

// NewMockIJwtService creates a new mock instance
func NewMockIJwtService(ctrl *gomock.Controller) *MockIJwtService {
	mock := &MockIJwtService{ctrl: ctrl}
	mock.recorder = &MockIJwtServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIJwtService) EXPECT() *MockIJwtServiceMockRecorder {
	return m.recorder
}

// CreateToken mocks base method
func (m *MockIJwtService) CreateToken(arg0 uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken
func (mr *MockIJwtServiceMockRecorder) CreateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockIJwtService)(nil).CreateToken), arg0)
}

// ValidateUser mocks base method
func (m *MockIJwtService) ValidateUser(arg0 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUser", arg0)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUser indicates an expected call of ValidateUser
func (mr *MockIJwtServiceMockRecorder) ValidateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUser", reflect.TypeOf((*MockIJwtService)(nil).ValidateUser), arg0)
}