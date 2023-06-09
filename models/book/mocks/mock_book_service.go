// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/snehil-sinha/goBookStore/models/book (interfaces: BookService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/snehil-sinha/goBookStore/common"
	book "github.com/snehil-sinha/goBookStore/models/book"
)

// MockBookService is a mock of BookService interface.
type MockBookService struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceMockRecorder
}

// MockBookServiceMockRecorder is the mock recorder for MockBookService.
type MockBookServiceMockRecorder struct {
	mock *MockBookService
}

// NewMockBookService creates a new mock instance.
func NewMockBookService(ctrl *gomock.Controller) *MockBookService {
	mock := &MockBookService{ctrl: ctrl}
	mock.recorder = &MockBookServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookService) EXPECT() *MockBookServiceMockRecorder {
	return m.recorder
}

// ReadAll mocks base method.
func (m *MockBookService) ReadAll(arg0 *common.Logger) ([]*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0)
	ret0, _ := ret[0].([]*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockBookServiceMockRecorder) ReadAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockBookService)(nil).ReadAll), arg0)
}

// ReadById mocks base method.
func (m *MockBookService) ReadById(arg0 *common.Logger, arg1 string) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadById", arg0, arg1)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadById indicates an expected call of ReadById.
func (mr *MockBookServiceMockRecorder) ReadById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadById", reflect.TypeOf((*MockBookService)(nil).ReadById), arg0, arg1)
}
