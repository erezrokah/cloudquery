// Code generated by MockGen. DO NOT EDIT.
// Source: waf.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	waf "github.com/aws/aws-sdk-go-v2/service/waf"
	gomock "github.com/golang/mock/gomock"
)

// MockWafClient is a mock of WafClient interface.
type MockWafClient struct {
	ctrl     *gomock.Controller
	recorder *MockWafClientMockRecorder
}

// MockWafClientMockRecorder is the mock recorder for MockWafClient.
type MockWafClientMockRecorder struct {
	mock *MockWafClient
}

// NewMockWafClient creates a new mock instance.
func NewMockWafClient(ctrl *gomock.Controller) *MockWafClient {
	mock := &MockWafClient{ctrl: ctrl}
	mock.recorder = &MockWafClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWafClient) EXPECT() *MockWafClientMockRecorder {
	return m.recorder
}

// GetByteMatchSet mocks base method.
func (m *MockWafClient) GetByteMatchSet(arg0 context.Context, arg1 *waf.GetByteMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetByteMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByteMatchSet", varargs...)
	ret0, _ := ret[0].(*waf.GetByteMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByteMatchSet indicates an expected call of GetByteMatchSet.
func (mr *MockWafClientMockRecorder) GetByteMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByteMatchSet", reflect.TypeOf((*MockWafClient)(nil).GetByteMatchSet), varargs...)
}

// GetChangeToken mocks base method.
func (m *MockWafClient) GetChangeToken(arg0 context.Context, arg1 *waf.GetChangeTokenInput, arg2 ...func(*waf.Options)) (*waf.GetChangeTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChangeToken", varargs...)
	ret0, _ := ret[0].(*waf.GetChangeTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChangeToken indicates an expected call of GetChangeToken.
func (mr *MockWafClientMockRecorder) GetChangeToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChangeToken", reflect.TypeOf((*MockWafClient)(nil).GetChangeToken), varargs...)
}

// GetChangeTokenStatus mocks base method.
func (m *MockWafClient) GetChangeTokenStatus(arg0 context.Context, arg1 *waf.GetChangeTokenStatusInput, arg2 ...func(*waf.Options)) (*waf.GetChangeTokenStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChangeTokenStatus", varargs...)
	ret0, _ := ret[0].(*waf.GetChangeTokenStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChangeTokenStatus indicates an expected call of GetChangeTokenStatus.
func (mr *MockWafClientMockRecorder) GetChangeTokenStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChangeTokenStatus", reflect.TypeOf((*MockWafClient)(nil).GetChangeTokenStatus), varargs...)
}

// GetGeoMatchSet mocks base method.
func (m *MockWafClient) GetGeoMatchSet(arg0 context.Context, arg1 *waf.GetGeoMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetGeoMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGeoMatchSet", varargs...)
	ret0, _ := ret[0].(*waf.GetGeoMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeoMatchSet indicates an expected call of GetGeoMatchSet.
func (mr *MockWafClientMockRecorder) GetGeoMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeoMatchSet", reflect.TypeOf((*MockWafClient)(nil).GetGeoMatchSet), varargs...)
}

// GetIPSet mocks base method.
func (m *MockWafClient) GetIPSet(arg0 context.Context, arg1 *waf.GetIPSetInput, arg2 ...func(*waf.Options)) (*waf.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIPSet", varargs...)
	ret0, _ := ret[0].(*waf.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPSet indicates an expected call of GetIPSet.
func (mr *MockWafClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPSet", reflect.TypeOf((*MockWafClient)(nil).GetIPSet), varargs...)
}

// GetLoggingConfiguration mocks base method.
func (m *MockWafClient) GetLoggingConfiguration(arg0 context.Context, arg1 *waf.GetLoggingConfigurationInput, arg2 ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*waf.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggingConfiguration indicates an expected call of GetLoggingConfiguration.
func (mr *MockWafClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggingConfiguration", reflect.TypeOf((*MockWafClient)(nil).GetLoggingConfiguration), varargs...)
}

// GetPermissionPolicy mocks base method.
func (m *MockWafClient) GetPermissionPolicy(arg0 context.Context, arg1 *waf.GetPermissionPolicyInput, arg2 ...func(*waf.Options)) (*waf.GetPermissionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPermissionPolicy", varargs...)
	ret0, _ := ret[0].(*waf.GetPermissionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionPolicy indicates an expected call of GetPermissionPolicy.
func (mr *MockWafClientMockRecorder) GetPermissionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionPolicy", reflect.TypeOf((*MockWafClient)(nil).GetPermissionPolicy), varargs...)
}

// GetRateBasedRule mocks base method.
func (m *MockWafClient) GetRateBasedRule(arg0 context.Context, arg1 *waf.GetRateBasedRuleInput, arg2 ...func(*waf.Options)) (*waf.GetRateBasedRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedRule", varargs...)
	ret0, _ := ret[0].(*waf.GetRateBasedRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateBasedRule indicates an expected call of GetRateBasedRule.
func (mr *MockWafClientMockRecorder) GetRateBasedRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedRule", reflect.TypeOf((*MockWafClient)(nil).GetRateBasedRule), varargs...)
}

// GetRateBasedRuleManagedKeys mocks base method.
func (m *MockWafClient) GetRateBasedRuleManagedKeys(arg0 context.Context, arg1 *waf.GetRateBasedRuleManagedKeysInput, arg2 ...func(*waf.Options)) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedRuleManagedKeys", varargs...)
	ret0, _ := ret[0].(*waf.GetRateBasedRuleManagedKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateBasedRuleManagedKeys indicates an expected call of GetRateBasedRuleManagedKeys.
func (mr *MockWafClientMockRecorder) GetRateBasedRuleManagedKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedRuleManagedKeys", reflect.TypeOf((*MockWafClient)(nil).GetRateBasedRuleManagedKeys), varargs...)
}

// GetRegexMatchSet mocks base method.
func (m *MockWafClient) GetRegexMatchSet(arg0 context.Context, arg1 *waf.GetRegexMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetRegexMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexMatchSet", varargs...)
	ret0, _ := ret[0].(*waf.GetRegexMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegexMatchSet indicates an expected call of GetRegexMatchSet.
func (mr *MockWafClientMockRecorder) GetRegexMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexMatchSet", reflect.TypeOf((*MockWafClient)(nil).GetRegexMatchSet), varargs...)
}

// GetRegexPatternSet mocks base method.
func (m *MockWafClient) GetRegexPatternSet(arg0 context.Context, arg1 *waf.GetRegexPatternSetInput, arg2 ...func(*waf.Options)) (*waf.GetRegexPatternSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegexPatternSet", varargs...)
	ret0, _ := ret[0].(*waf.GetRegexPatternSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegexPatternSet indicates an expected call of GetRegexPatternSet.
func (mr *MockWafClientMockRecorder) GetRegexPatternSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegexPatternSet", reflect.TypeOf((*MockWafClient)(nil).GetRegexPatternSet), varargs...)
}

// GetRule mocks base method.
func (m *MockWafClient) GetRule(arg0 context.Context, arg1 *waf.GetRuleInput, arg2 ...func(*waf.Options)) (*waf.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRule", varargs...)
	ret0, _ := ret[0].(*waf.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRule indicates an expected call of GetRule.
func (mr *MockWafClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockWafClient)(nil).GetRule), varargs...)
}

// GetRuleGroup mocks base method.
func (m *MockWafClient) GetRuleGroup(arg0 context.Context, arg1 *waf.GetRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*waf.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleGroup indicates an expected call of GetRuleGroup.
func (mr *MockWafClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafClient)(nil).GetRuleGroup), varargs...)
}

// GetSampledRequests mocks base method.
func (m *MockWafClient) GetSampledRequests(arg0 context.Context, arg1 *waf.GetSampledRequestsInput, arg2 ...func(*waf.Options)) (*waf.GetSampledRequestsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSampledRequests", varargs...)
	ret0, _ := ret[0].(*waf.GetSampledRequestsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSampledRequests indicates an expected call of GetSampledRequests.
func (mr *MockWafClientMockRecorder) GetSampledRequests(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSampledRequests", reflect.TypeOf((*MockWafClient)(nil).GetSampledRequests), varargs...)
}

// GetSizeConstraintSet mocks base method.
func (m *MockWafClient) GetSizeConstraintSet(arg0 context.Context, arg1 *waf.GetSizeConstraintSetInput, arg2 ...func(*waf.Options)) (*waf.GetSizeConstraintSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSizeConstraintSet", varargs...)
	ret0, _ := ret[0].(*waf.GetSizeConstraintSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSizeConstraintSet indicates an expected call of GetSizeConstraintSet.
func (mr *MockWafClientMockRecorder) GetSizeConstraintSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSizeConstraintSet", reflect.TypeOf((*MockWafClient)(nil).GetSizeConstraintSet), varargs...)
}

// GetSqlInjectionMatchSet mocks base method.
func (m *MockWafClient) GetSqlInjectionMatchSet(arg0 context.Context, arg1 *waf.GetSqlInjectionMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetSqlInjectionMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSqlInjectionMatchSet", varargs...)
	ret0, _ := ret[0].(*waf.GetSqlInjectionMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSqlInjectionMatchSet indicates an expected call of GetSqlInjectionMatchSet.
func (mr *MockWafClientMockRecorder) GetSqlInjectionMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSqlInjectionMatchSet", reflect.TypeOf((*MockWafClient)(nil).GetSqlInjectionMatchSet), varargs...)
}

// GetWebACL mocks base method.
func (m *MockWafClient) GetWebACL(arg0 context.Context, arg1 *waf.GetWebACLInput, arg2 ...func(*waf.Options)) (*waf.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*waf.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACL indicates an expected call of GetWebACL.
func (mr *MockWafClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafClient)(nil).GetWebACL), varargs...)
}

// GetXssMatchSet mocks base method.
func (m *MockWafClient) GetXssMatchSet(arg0 context.Context, arg1 *waf.GetXssMatchSetInput, arg2 ...func(*waf.Options)) (*waf.GetXssMatchSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetXssMatchSet", varargs...)
	ret0, _ := ret[0].(*waf.GetXssMatchSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetXssMatchSet indicates an expected call of GetXssMatchSet.
func (mr *MockWafClientMockRecorder) GetXssMatchSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXssMatchSet", reflect.TypeOf((*MockWafClient)(nil).GetXssMatchSet), varargs...)
}

// ListActivatedRulesInRuleGroup mocks base method.
func (m *MockWafClient) ListActivatedRulesInRuleGroup(arg0 context.Context, arg1 *waf.ListActivatedRulesInRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListActivatedRulesInRuleGroup", varargs...)
	ret0, _ := ret[0].(*waf.ListActivatedRulesInRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivatedRulesInRuleGroup indicates an expected call of ListActivatedRulesInRuleGroup.
func (mr *MockWafClientMockRecorder) ListActivatedRulesInRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivatedRulesInRuleGroup", reflect.TypeOf((*MockWafClient)(nil).ListActivatedRulesInRuleGroup), varargs...)
}

// ListByteMatchSets mocks base method.
func (m *MockWafClient) ListByteMatchSets(arg0 context.Context, arg1 *waf.ListByteMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListByteMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByteMatchSets", varargs...)
	ret0, _ := ret[0].(*waf.ListByteMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByteMatchSets indicates an expected call of ListByteMatchSets.
func (mr *MockWafClientMockRecorder) ListByteMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByteMatchSets", reflect.TypeOf((*MockWafClient)(nil).ListByteMatchSets), varargs...)
}

// ListGeoMatchSets mocks base method.
func (m *MockWafClient) ListGeoMatchSets(arg0 context.Context, arg1 *waf.ListGeoMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListGeoMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGeoMatchSets", varargs...)
	ret0, _ := ret[0].(*waf.ListGeoMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGeoMatchSets indicates an expected call of ListGeoMatchSets.
func (mr *MockWafClientMockRecorder) ListGeoMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGeoMatchSets", reflect.TypeOf((*MockWafClient)(nil).ListGeoMatchSets), varargs...)
}

// ListIPSets mocks base method.
func (m *MockWafClient) ListIPSets(arg0 context.Context, arg1 *waf.ListIPSetsInput, arg2 ...func(*waf.Options)) (*waf.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIPSets", varargs...)
	ret0, _ := ret[0].(*waf.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIPSets indicates an expected call of ListIPSets.
func (mr *MockWafClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPSets", reflect.TypeOf((*MockWafClient)(nil).ListIPSets), varargs...)
}

// ListLoggingConfigurations mocks base method.
func (m *MockWafClient) ListLoggingConfigurations(arg0 context.Context, arg1 *waf.ListLoggingConfigurationsInput, arg2 ...func(*waf.Options)) (*waf.ListLoggingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLoggingConfigurations", varargs...)
	ret0, _ := ret[0].(*waf.ListLoggingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoggingConfigurations indicates an expected call of ListLoggingConfigurations.
func (mr *MockWafClientMockRecorder) ListLoggingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoggingConfigurations", reflect.TypeOf((*MockWafClient)(nil).ListLoggingConfigurations), varargs...)
}

// ListRateBasedRules mocks base method.
func (m *MockWafClient) ListRateBasedRules(arg0 context.Context, arg1 *waf.ListRateBasedRulesInput, arg2 ...func(*waf.Options)) (*waf.ListRateBasedRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRateBasedRules", varargs...)
	ret0, _ := ret[0].(*waf.ListRateBasedRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRateBasedRules indicates an expected call of ListRateBasedRules.
func (mr *MockWafClientMockRecorder) ListRateBasedRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRateBasedRules", reflect.TypeOf((*MockWafClient)(nil).ListRateBasedRules), varargs...)
}

// ListRegexMatchSets mocks base method.
func (m *MockWafClient) ListRegexMatchSets(arg0 context.Context, arg1 *waf.ListRegexMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListRegexMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexMatchSets", varargs...)
	ret0, _ := ret[0].(*waf.ListRegexMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegexMatchSets indicates an expected call of ListRegexMatchSets.
func (mr *MockWafClientMockRecorder) ListRegexMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexMatchSets", reflect.TypeOf((*MockWafClient)(nil).ListRegexMatchSets), varargs...)
}

// ListRegexPatternSets mocks base method.
func (m *MockWafClient) ListRegexPatternSets(arg0 context.Context, arg1 *waf.ListRegexPatternSetsInput, arg2 ...func(*waf.Options)) (*waf.ListRegexPatternSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegexPatternSets", varargs...)
	ret0, _ := ret[0].(*waf.ListRegexPatternSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegexPatternSets indicates an expected call of ListRegexPatternSets.
func (mr *MockWafClientMockRecorder) ListRegexPatternSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegexPatternSets", reflect.TypeOf((*MockWafClient)(nil).ListRegexPatternSets), varargs...)
}

// ListRuleGroups mocks base method.
func (m *MockWafClient) ListRuleGroups(arg0 context.Context, arg1 *waf.ListRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*waf.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuleGroups indicates an expected call of ListRuleGroups.
func (mr *MockWafClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafClient)(nil).ListRuleGroups), varargs...)
}

// ListRules mocks base method.
func (m *MockWafClient) ListRules(arg0 context.Context, arg1 *waf.ListRulesInput, arg2 ...func(*waf.Options)) (*waf.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*waf.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules.
func (mr *MockWafClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockWafClient)(nil).ListRules), varargs...)
}

// ListSizeConstraintSets mocks base method.
func (m *MockWafClient) ListSizeConstraintSets(arg0 context.Context, arg1 *waf.ListSizeConstraintSetsInput, arg2 ...func(*waf.Options)) (*waf.ListSizeConstraintSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSizeConstraintSets", varargs...)
	ret0, _ := ret[0].(*waf.ListSizeConstraintSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSizeConstraintSets indicates an expected call of ListSizeConstraintSets.
func (mr *MockWafClientMockRecorder) ListSizeConstraintSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSizeConstraintSets", reflect.TypeOf((*MockWafClient)(nil).ListSizeConstraintSets), varargs...)
}

// ListSqlInjectionMatchSets mocks base method.
func (m *MockWafClient) ListSqlInjectionMatchSets(arg0 context.Context, arg1 *waf.ListSqlInjectionMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSqlInjectionMatchSets", varargs...)
	ret0, _ := ret[0].(*waf.ListSqlInjectionMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSqlInjectionMatchSets indicates an expected call of ListSqlInjectionMatchSets.
func (mr *MockWafClientMockRecorder) ListSqlInjectionMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSqlInjectionMatchSets", reflect.TypeOf((*MockWafClient)(nil).ListSqlInjectionMatchSets), varargs...)
}

// ListSubscribedRuleGroups mocks base method.
func (m *MockWafClient) ListSubscribedRuleGroups(arg0 context.Context, arg1 *waf.ListSubscribedRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSubscribedRuleGroups", varargs...)
	ret0, _ := ret[0].(*waf.ListSubscribedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscribedRuleGroups indicates an expected call of ListSubscribedRuleGroups.
func (mr *MockWafClientMockRecorder) ListSubscribedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscribedRuleGroups", reflect.TypeOf((*MockWafClient)(nil).ListSubscribedRuleGroups), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockWafClient) ListTagsForResource(arg0 context.Context, arg1 *waf.ListTagsForResourceInput, arg2 ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*waf.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockWafClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafClient)(nil).ListTagsForResource), varargs...)
}

// ListWebACLs mocks base method.
func (m *MockWafClient) ListWebACLs(arg0 context.Context, arg1 *waf.ListWebACLsInput, arg2 ...func(*waf.Options)) (*waf.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*waf.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebACLs indicates an expected call of ListWebACLs.
func (mr *MockWafClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafClient)(nil).ListWebACLs), varargs...)
}

// ListXssMatchSets mocks base method.
func (m *MockWafClient) ListXssMatchSets(arg0 context.Context, arg1 *waf.ListXssMatchSetsInput, arg2 ...func(*waf.Options)) (*waf.ListXssMatchSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListXssMatchSets", varargs...)
	ret0, _ := ret[0].(*waf.ListXssMatchSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListXssMatchSets indicates an expected call of ListXssMatchSets.
func (mr *MockWafClientMockRecorder) ListXssMatchSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXssMatchSets", reflect.TypeOf((*MockWafClient)(nil).ListXssMatchSets), varargs...)
}
