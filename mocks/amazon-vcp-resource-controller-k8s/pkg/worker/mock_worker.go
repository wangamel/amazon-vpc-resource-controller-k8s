// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/worker (interfaces: Worker)

// Package mock_worker is a generated GoMock package.
package mock_worker

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWorker is a mock of Worker interface
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// StartWorkerPool mocks base method
func (m *MockWorker) StartWorkerPool() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkerPool")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWorkerPool indicates an expected call of StartWorkerPool
func (mr *MockWorkerMockRecorder) StartWorkerPool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkerPool", reflect.TypeOf((*MockWorker)(nil).StartWorkerPool))
}

// SubmitJob mocks base method
func (m *MockWorker) SubmitJob(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitJob indicates an expected call of SubmitJob
func (mr *MockWorkerMockRecorder) SubmitJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitJob", reflect.TypeOf((*MockWorker)(nil).SubmitJob), arg0)
}