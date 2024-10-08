// Code generated by MockGen. DO NOT EDIT.
// Source: authorize.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	client "github.com/GSVillas/pic-pay-desafio/client"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorizationService is a mock of AuthorizationService interface.
type MockAuthorizationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationServiceMockRecorder
}

// MockAuthorizationServiceMockRecorder is the mock recorder for MockAuthorizationService.
type MockAuthorizationServiceMockRecorder struct {
	mock *MockAuthorizationService
}

// NewMockAuthorizationService creates a new mock instance.
func NewMockAuthorizationService(ctrl *gomock.Controller) *MockAuthorizationService {
	mock := &MockAuthorizationService{ctrl: ctrl}
	mock.recorder = &MockAuthorizationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationService) EXPECT() *MockAuthorizationServiceMockRecorder {
	return m.recorder
}

// CheckAuthorization mocks base method.
func (m *MockAuthorizationService) CheckAuthorization(ctx context.Context) (*client.AuthorizationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthorization", ctx)
	ret0, _ := ret[0].(*client.AuthorizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuthorization indicates an expected call of CheckAuthorization.
func (mr *MockAuthorizationServiceMockRecorder) CheckAuthorization(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthorization", reflect.TypeOf((*MockAuthorizationService)(nil).CheckAuthorization), ctx)
}
