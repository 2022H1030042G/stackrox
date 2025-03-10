// Code generated by MockGen. DO NOT EDIT.
// Source: unsafe_searcher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

// MockUnsafeSearcher is a mock of UnsafeSearcher interface.
type MockUnsafeSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSearcherMockRecorder
}

// MockUnsafeSearcherMockRecorder is the mock recorder for MockUnsafeSearcher.
type MockUnsafeSearcherMockRecorder struct {
	mock *MockUnsafeSearcher
}

// NewMockUnsafeSearcher creates a new mock instance.
func NewMockUnsafeSearcher(ctrl *gomock.Controller) *MockUnsafeSearcher {
	mock := &MockUnsafeSearcher{ctrl: ctrl}
	mock.recorder = &MockUnsafeSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSearcher) EXPECT() *MockUnsafeSearcherMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockUnsafeSearcher) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUnsafeSearcherMockRecorder) Count(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUnsafeSearcher)(nil).Count), varargs...)
}

// Search mocks base method.
func (m *MockUnsafeSearcher) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockUnsafeSearcherMockRecorder) Search(ctx, q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUnsafeSearcher)(nil).Search), varargs...)
}
