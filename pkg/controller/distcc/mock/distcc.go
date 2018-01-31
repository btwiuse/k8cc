// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mbrt/k8cc/pkg/controller/distcc (interfaces: Operator)

// Package mock_distcc is a generated GoMock package.
package mock_distcc

import (
	gomock "github.com/golang/mock/gomock"
	data "github.com/mbrt/k8cc/pkg/data"
	reflect "reflect"
)

// MockOperator is a mock of Operator interface
type MockOperator struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorMockRecorder
}

// MockOperatorMockRecorder is the mock recorder for MockOperator
type MockOperatorMockRecorder struct {
	mock *MockOperator
}

// NewMockOperator creates a new mock instance
func NewMockOperator(ctrl *gomock.Controller) *MockOperator {
	mock := &MockOperator{ctrl: ctrl}
	mock.recorder = &MockOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperator) EXPECT() *MockOperatorMockRecorder {
	return m.recorder
}

// Hostnames mocks base method
func (m *MockOperator) Hostnames(arg0 data.Tag, arg1 []data.HostID) ([]string, error) {
	ret := m.ctrl.Call(m, "Hostnames", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hostnames indicates an expected call of Hostnames
func (mr *MockOperatorMockRecorder) Hostnames(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hostnames", reflect.TypeOf((*MockOperator)(nil).Hostnames), arg0, arg1)
}

// NotifyUpdated mocks base method
func (m *MockOperator) NotifyUpdated(arg0 data.Tag) error {
	ret := m.ctrl.Call(m, "NotifyUpdated", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyUpdated indicates an expected call of NotifyUpdated
func (mr *MockOperatorMockRecorder) NotifyUpdated(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyUpdated", reflect.TypeOf((*MockOperator)(nil).NotifyUpdated), arg0)
}

// Run mocks base method
func (m *MockOperator) Run(arg0 int, arg1 <-chan struct{}) error {
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockOperatorMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockOperator)(nil).Run), arg0, arg1)
}

// ScaleSettings mocks base method
func (m *MockOperator) ScaleSettings(arg0 data.Tag) (data.ScaleSettings, error) {
	ret := m.ctrl.Call(m, "ScaleSettings", arg0)
	ret0, _ := ret[0].(data.ScaleSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleSettings indicates an expected call of ScaleSettings
func (mr *MockOperatorMockRecorder) ScaleSettings(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleSettings", reflect.TypeOf((*MockOperator)(nil).ScaleSettings), arg0)
}