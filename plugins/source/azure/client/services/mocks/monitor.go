// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: ActivityLogAlertsClient,LogProfilesClient,DiagnosticSettingsClient,ActivityLogsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	insights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	insights0 "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	gomock "github.com/golang/mock/gomock"
)

// MockActivityLogAlertsClient is a mock of ActivityLogAlertsClient interface.
type MockActivityLogAlertsClient struct {
	ctrl     *gomock.Controller
	recorder *MockActivityLogAlertsClientMockRecorder
}

// MockActivityLogAlertsClientMockRecorder is the mock recorder for MockActivityLogAlertsClient.
type MockActivityLogAlertsClientMockRecorder struct {
	mock *MockActivityLogAlertsClient
}

// NewMockActivityLogAlertsClient creates a new mock instance.
func NewMockActivityLogAlertsClient(ctrl *gomock.Controller) *MockActivityLogAlertsClient {
	mock := &MockActivityLogAlertsClient{ctrl: ctrl}
	mock.recorder = &MockActivityLogAlertsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivityLogAlertsClient) EXPECT() *MockActivityLogAlertsClientMockRecorder {
	return m.recorder
}

// ListBySubscriptionID mocks base method.
func (m *MockActivityLogAlertsClient) ListBySubscriptionID(arg0 context.Context) (insights.ActivityLogAlertList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscriptionID", arg0)
	ret0, _ := ret[0].(insights.ActivityLogAlertList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscriptionID indicates an expected call of ListBySubscriptionID.
func (mr *MockActivityLogAlertsClientMockRecorder) ListBySubscriptionID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscriptionID", reflect.TypeOf((*MockActivityLogAlertsClient)(nil).ListBySubscriptionID), arg0)
}

// MockLogProfilesClient is a mock of LogProfilesClient interface.
type MockLogProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockLogProfilesClientMockRecorder
}

// MockLogProfilesClientMockRecorder is the mock recorder for MockLogProfilesClient.
type MockLogProfilesClientMockRecorder struct {
	mock *MockLogProfilesClient
}

// NewMockLogProfilesClient creates a new mock instance.
func NewMockLogProfilesClient(ctrl *gomock.Controller) *MockLogProfilesClient {
	mock := &MockLogProfilesClient{ctrl: ctrl}
	mock.recorder = &MockLogProfilesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogProfilesClient) EXPECT() *MockLogProfilesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockLogProfilesClient) List(arg0 context.Context) (insights0.LogProfileCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(insights0.LogProfileCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLogProfilesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLogProfilesClient)(nil).List), arg0)
}

// MockDiagnosticSettingsClient is a mock of DiagnosticSettingsClient interface.
type MockDiagnosticSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosticSettingsClientMockRecorder
}

// MockDiagnosticSettingsClientMockRecorder is the mock recorder for MockDiagnosticSettingsClient.
type MockDiagnosticSettingsClientMockRecorder struct {
	mock *MockDiagnosticSettingsClient
}

// NewMockDiagnosticSettingsClient creates a new mock instance.
func NewMockDiagnosticSettingsClient(ctrl *gomock.Controller) *MockDiagnosticSettingsClient {
	mock := &MockDiagnosticSettingsClient{ctrl: ctrl}
	mock.recorder = &MockDiagnosticSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosticSettingsClient) EXPECT() *MockDiagnosticSettingsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockDiagnosticSettingsClient) List(arg0 context.Context, arg1 string) (insights0.DiagnosticSettingsResourceCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(insights0.DiagnosticSettingsResourceCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDiagnosticSettingsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDiagnosticSettingsClient)(nil).List), arg0, arg1)
}

// MockActivityLogsClient is a mock of ActivityLogsClient interface.
type MockActivityLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockActivityLogsClientMockRecorder
}

// MockActivityLogsClientMockRecorder is the mock recorder for MockActivityLogsClient.
type MockActivityLogsClientMockRecorder struct {
	mock *MockActivityLogsClient
}

// NewMockActivityLogsClient creates a new mock instance.
func NewMockActivityLogsClient(ctrl *gomock.Controller) *MockActivityLogsClient {
	mock := &MockActivityLogsClient{ctrl: ctrl}
	mock.recorder = &MockActivityLogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivityLogsClient) EXPECT() *MockActivityLogsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockActivityLogsClient) List(arg0 context.Context, arg1, arg2 string) (insights0.EventDataCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(insights0.EventDataCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockActivityLogsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockActivityLogsClient)(nil).List), arg0, arg1, arg2)
}
