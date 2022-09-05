// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: ResourcesGroupsClient,ResourcesPolicyAssignmentsClient,ResourcesLinksClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	policy "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	links "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	gomock "github.com/golang/mock/gomock"
)

// MockResourcesGroupsClient is a mock of ResourcesGroupsClient interface.
type MockResourcesGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesGroupsClientMockRecorder
}

// MockResourcesGroupsClientMockRecorder is the mock recorder for MockResourcesGroupsClient.
type MockResourcesGroupsClientMockRecorder struct {
	mock *MockResourcesGroupsClient
}

// NewMockResourcesGroupsClient creates a new mock instance.
func NewMockResourcesGroupsClient(ctrl *gomock.Controller) *MockResourcesGroupsClient {
	mock := &MockResourcesGroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourcesGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcesGroupsClient) EXPECT() *MockResourcesGroupsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockResourcesGroupsClient) List(arg0 context.Context, arg1 string, arg2 *int32) (resources.GroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.GroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourcesGroupsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourcesGroupsClient)(nil).List), arg0, arg1, arg2)
}

// MockResourcesPolicyAssignmentsClient is a mock of ResourcesPolicyAssignmentsClient interface.
type MockResourcesPolicyAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesPolicyAssignmentsClientMockRecorder
}

// MockResourcesPolicyAssignmentsClientMockRecorder is the mock recorder for MockResourcesPolicyAssignmentsClient.
type MockResourcesPolicyAssignmentsClientMockRecorder struct {
	mock *MockResourcesPolicyAssignmentsClient
}

// NewMockResourcesPolicyAssignmentsClient creates a new mock instance.
func NewMockResourcesPolicyAssignmentsClient(ctrl *gomock.Controller) *MockResourcesPolicyAssignmentsClient {
	mock := &MockResourcesPolicyAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockResourcesPolicyAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcesPolicyAssignmentsClient) EXPECT() *MockResourcesPolicyAssignmentsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockResourcesPolicyAssignmentsClient) List(arg0 context.Context, arg1, arg2 string, arg3 *int32) (policy.AssignmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(policy.AssignmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourcesPolicyAssignmentsClientMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourcesPolicyAssignmentsClient)(nil).List), arg0, arg1, arg2, arg3)
}

// MockResourcesLinksClient is a mock of ResourcesLinksClient interface.
type MockResourcesLinksClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesLinksClientMockRecorder
}

// MockResourcesLinksClientMockRecorder is the mock recorder for MockResourcesLinksClient.
type MockResourcesLinksClientMockRecorder struct {
	mock *MockResourcesLinksClient
}

// NewMockResourcesLinksClient creates a new mock instance.
func NewMockResourcesLinksClient(ctrl *gomock.Controller) *MockResourcesLinksClient {
	mock := &MockResourcesLinksClient{ctrl: ctrl}
	mock.recorder = &MockResourcesLinksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcesLinksClient) EXPECT() *MockResourcesLinksClientMockRecorder {
	return m.recorder
}

// ListAtSubscription mocks base method.
func (m *MockResourcesLinksClient) ListAtSubscription(arg0 context.Context, arg1 string) (links.ResourceLinkResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAtSubscription", arg0, arg1)
	ret0, _ := ret[0].(links.ResourceLinkResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAtSubscription indicates an expected call of ListAtSubscription.
func (mr *MockResourcesLinksClientMockRecorder) ListAtSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAtSubscription", reflect.TypeOf((*MockResourcesLinksClient)(nil).ListAtSubscription), arg0, arg1)
}
