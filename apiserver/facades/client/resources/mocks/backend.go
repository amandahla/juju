// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/resources (interfaces: Backend,NewCharmRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	resource "github.com/juju/charm/v11/resource"
	resources "github.com/juju/juju/apiserver/facades/client/resources"
	resources0 "github.com/juju/juju/core/resources"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AddPendingResource mocks base method.
func (m *MockBackend) AddPendingResource(arg0, arg1 string, arg2 resource.Resource) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPendingResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPendingResource indicates an expected call of AddPendingResource.
func (mr *MockBackendMockRecorder) AddPendingResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPendingResource", reflect.TypeOf((*MockBackend)(nil).AddPendingResource), arg0, arg1, arg2)
}

// ListResources mocks base method.
func (m *MockBackend) ListResources(arg0 string) (resources0.ApplicationResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0)
	ret0, _ := ret[0].(resources0.ApplicationResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockBackendMockRecorder) ListResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockBackend)(nil).ListResources), arg0)
}

// MockNewCharmRepository is a mock of NewCharmRepository interface.
type MockNewCharmRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNewCharmRepositoryMockRecorder
}

// MockNewCharmRepositoryMockRecorder is the mock recorder for MockNewCharmRepository.
type MockNewCharmRepositoryMockRecorder struct {
	mock *MockNewCharmRepository
}

// NewMockNewCharmRepository creates a new mock instance.
func NewMockNewCharmRepository(ctrl *gomock.Controller) *MockNewCharmRepository {
	mock := &MockNewCharmRepository{ctrl: ctrl}
	mock.recorder = &MockNewCharmRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewCharmRepository) EXPECT() *MockNewCharmRepositoryMockRecorder {
	return m.recorder
}

// ResolveResources mocks base method.
func (m *MockNewCharmRepository) ResolveResources(arg0 []resource.Resource, arg1 resources.CharmID) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveResources", arg0, arg1)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveResources indicates an expected call of ResolveResources.
func (mr *MockNewCharmRepositoryMockRecorder) ResolveResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveResources", reflect.TypeOf((*MockNewCharmRepository)(nil).ResolveResources), arg0, arg1)
}
