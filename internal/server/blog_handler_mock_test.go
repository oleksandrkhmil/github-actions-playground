// Code generated by MockGen. DO NOT EDIT.
// Source: blog_handler.go
//
// Generated by this command:
//
//	mockgen -source=blog_handler.go -destination=blog_handler_mock_test.go -package=server_test -typed=true
//

// Package server_test is a generated GoMock package.
package server_test

import (
	context "context"
	reflect "reflect"

	blog "github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
	gomock "go.uber.org/mock/gomock"
)

// MockblogService is a mock of blogService interface.
type MockblogService struct {
	ctrl     *gomock.Controller
	recorder *MockblogServiceMockRecorder
}

// MockblogServiceMockRecorder is the mock recorder for MockblogService.
type MockblogServiceMockRecorder struct {
	mock *MockblogService
}

// NewMockblogService creates a new mock instance.
func NewMockblogService(ctrl *gomock.Controller) *MockblogService {
	mock := &MockblogService{ctrl: ctrl}
	mock.recorder = &MockblogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockblogService) EXPECT() *MockblogServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockblogService) Create(arg0 context.Context, arg1 blog.Post) (blog.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(blog.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockblogServiceMockRecorder) Create(arg0, arg1 any) *MockblogServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockblogService)(nil).Create), arg0, arg1)
	return &MockblogServiceCreateCall{Call: call}
}

// MockblogServiceCreateCall wrap *gomock.Call
type MockblogServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockblogServiceCreateCall) Return(arg0 blog.Post, arg1 error) *MockblogServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockblogServiceCreateCall) Do(f func(context.Context, blog.Post) (blog.Post, error)) *MockblogServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockblogServiceCreateCall) DoAndReturn(f func(context.Context, blog.Post) (blog.Post, error)) *MockblogServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAll mocks base method.
func (m *MockblogService) GetAll(arg0 context.Context) ([]blog.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]blog.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockblogServiceMockRecorder) GetAll(arg0 any) *MockblogServiceGetAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockblogService)(nil).GetAll), arg0)
	return &MockblogServiceGetAllCall{Call: call}
}

// MockblogServiceGetAllCall wrap *gomock.Call
type MockblogServiceGetAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockblogServiceGetAllCall) Return(arg0 []blog.Post, arg1 error) *MockblogServiceGetAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockblogServiceGetAllCall) Do(f func(context.Context) ([]blog.Post, error)) *MockblogServiceGetAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockblogServiceGetAllCall) DoAndReturn(f func(context.Context) ([]blog.Post, error)) *MockblogServiceGetAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
