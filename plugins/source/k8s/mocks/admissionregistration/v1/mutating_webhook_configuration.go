// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/admissionregistration/v1 (interfaces: MutatingWebhookConfigurationsGetter,MutatingWebhookConfigurationInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/admissionregistration/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	v12 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

// MockMutatingWebhookConfigurationsGetter is a mock of MutatingWebhookConfigurationsGetter interface.
type MockMutatingWebhookConfigurationsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockMutatingWebhookConfigurationsGetterMockRecorder
}

// MockMutatingWebhookConfigurationsGetterMockRecorder is the mock recorder for MockMutatingWebhookConfigurationsGetter.
type MockMutatingWebhookConfigurationsGetterMockRecorder struct {
	mock *MockMutatingWebhookConfigurationsGetter
}

// NewMockMutatingWebhookConfigurationsGetter creates a new mock instance.
func NewMockMutatingWebhookConfigurationsGetter(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationsGetter {
	mock := &MockMutatingWebhookConfigurationsGetter{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutatingWebhookConfigurationsGetter) EXPECT() *MockMutatingWebhookConfigurationsGetterMockRecorder {
	return m.recorder
}

// MutatingWebhookConfigurations mocks base method.
func (m *MockMutatingWebhookConfigurationsGetter) MutatingWebhookConfigurations() v12.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutatingWebhookConfigurations")
	ret0, _ := ret[0].(v12.MutatingWebhookConfigurationInterface)
	return ret0
}

// MutatingWebhookConfigurations indicates an expected call of MutatingWebhookConfigurations.
func (mr *MockMutatingWebhookConfigurationsGetterMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutatingWebhookConfigurations", reflect.TypeOf((*MockMutatingWebhookConfigurationsGetter)(nil).MutatingWebhookConfigurations))
}

// MockMutatingWebhookConfigurationInterface is a mock of MutatingWebhookConfigurationInterface interface.
type MockMutatingWebhookConfigurationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMutatingWebhookConfigurationInterfaceMockRecorder
}

// MockMutatingWebhookConfigurationInterfaceMockRecorder is the mock recorder for MockMutatingWebhookConfigurationInterface.
type MockMutatingWebhookConfigurationInterfaceMockRecorder struct {
	mock *MockMutatingWebhookConfigurationInterface
}

// NewMockMutatingWebhookConfigurationInterface creates a new mock instance.
func NewMockMutatingWebhookConfigurationInterface(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationInterface {
	mock := &MockMutatingWebhookConfigurationInterface{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutatingWebhookConfigurationInterface) EXPECT() *MockMutatingWebhookConfigurationInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Apply(arg0 context.Context, arg1 *v11.MutatingWebhookConfigurationApplyConfiguration, arg2 v10.ApplyOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Create(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.CreateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.MutatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Update(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.UpdateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockMutatingWebhookConfigurationInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Watch), arg0, arg1)
}
