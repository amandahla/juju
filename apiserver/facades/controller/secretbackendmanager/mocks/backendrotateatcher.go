// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: SecretBackendRotateWatcher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	watcher "github.com/juju/juju/core/watcher"
)

// MockSecretBackendRotateWatcher is a mock of SecretBackendRotateWatcher interface.
type MockSecretBackendRotateWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendRotateWatcherMockRecorder
}

// MockSecretBackendRotateWatcherMockRecorder is the mock recorder for MockSecretBackendRotateWatcher.
type MockSecretBackendRotateWatcherMockRecorder struct {
	mock *MockSecretBackendRotateWatcher
}

// NewMockSecretBackendRotateWatcher creates a new mock instance.
func NewMockSecretBackendRotateWatcher(ctrl *gomock.Controller) *MockSecretBackendRotateWatcher {
	mock := &MockSecretBackendRotateWatcher{ctrl: ctrl}
	mock.recorder = &MockSecretBackendRotateWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendRotateWatcher) EXPECT() *MockSecretBackendRotateWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockSecretBackendRotateWatcher) Changes() watcher.SecretBackendRotateChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(watcher.SecretBackendRotateChannel)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockSecretBackendRotateWatcherMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockSecretBackendRotateWatcher)(nil).Changes))
}

// Err mocks base method.
func (m *MockSecretBackendRotateWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSecretBackendRotateWatcherMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSecretBackendRotateWatcher)(nil).Err))
}

// Kill mocks base method.
func (m *MockSecretBackendRotateWatcher) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockSecretBackendRotateWatcherMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockSecretBackendRotateWatcher)(nil).Kill))
}

// Stop mocks base method.
func (m *MockSecretBackendRotateWatcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockSecretBackendRotateWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSecretBackendRotateWatcher)(nil).Stop))
}

// Wait mocks base method.
func (m *MockSecretBackendRotateWatcher) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockSecretBackendRotateWatcherMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockSecretBackendRotateWatcher)(nil).Wait))
}
