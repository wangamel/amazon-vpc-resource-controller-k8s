// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/node (interfaces: Manager)

// Package mock_node is a generated GoMock package.
package mock_node

import (
	reflect "reflect"

	node "github.com/aws/amazon-vpc-resource-controller-k8s/pkg/node"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddOrUpdateNode mocks base method
func (m *MockManager) AddOrUpdateNode(arg0 *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateNode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateNode indicates an expected call of AddOrUpdateNode
func (mr *MockManagerMockRecorder) AddOrUpdateNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateNode", reflect.TypeOf((*MockManager)(nil).AddOrUpdateNode), arg0)
}

// DeleteNode mocks base method
func (m *MockManager) DeleteNode(arg0 *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode
func (mr *MockManagerMockRecorder) DeleteNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockManager)(nil).DeleteNode), arg0)
}

// GetNode mocks base method
func (m *MockManager) GetNode(arg0 string) (node.Node, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", arg0)
	ret0, _ := ret[0].(node.Node)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode
func (mr *MockManagerMockRecorder) GetNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockManager)(nil).GetNode), arg0)
}