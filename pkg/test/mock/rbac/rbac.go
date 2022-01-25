// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/rbac/rbac.go

// Package mock_rbac is a generated GoMock package.
package mock_rbac

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/kralicky/opni-monitoring/pkg/core"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// SubjectAccess mocks base method.
func (m *MockProvider) SubjectAccess(arg0 context.Context, arg1 *core.SubjectAccessRequest) (*core.ReferenceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectAccess", arg0, arg1)
	ret0, _ := ret[0].(*core.ReferenceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubjectAccess indicates an expected call of SubjectAccess.
func (mr *MockProviderMockRecorder) SubjectAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectAccess", reflect.TypeOf((*MockProvider)(nil).SubjectAccess), arg0, arg1)
}
