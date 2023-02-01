// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

//go:generate mockgen -package=mocks -destination=../mocks/securityhub.go -source=securityhub.go SecurityhubClient
type SecurityhubClient interface {
	DescribeActionTargets(context.Context, *securityhub.DescribeActionTargetsInput, ...func(*securityhub.Options)) (*securityhub.DescribeActionTargetsOutput, error)
	DescribeHub(context.Context, *securityhub.DescribeHubInput, ...func(*securityhub.Options)) (*securityhub.DescribeHubOutput, error)
	DescribeOrganizationConfiguration(context.Context, *securityhub.DescribeOrganizationConfigurationInput, ...func(*securityhub.Options)) (*securityhub.DescribeOrganizationConfigurationOutput, error)
	DescribeProducts(context.Context, *securityhub.DescribeProductsInput, ...func(*securityhub.Options)) (*securityhub.DescribeProductsOutput, error)
	DescribeStandards(context.Context, *securityhub.DescribeStandardsInput, ...func(*securityhub.Options)) (*securityhub.DescribeStandardsOutput, error)
	DescribeStandardsControls(context.Context, *securityhub.DescribeStandardsControlsInput, ...func(*securityhub.Options)) (*securityhub.DescribeStandardsControlsOutput, error)
	GetAdministratorAccount(context.Context, *securityhub.GetAdministratorAccountInput, ...func(*securityhub.Options)) (*securityhub.GetAdministratorAccountOutput, error)
	GetEnabledStandards(context.Context, *securityhub.GetEnabledStandardsInput, ...func(*securityhub.Options)) (*securityhub.GetEnabledStandardsOutput, error)
	GetFindingAggregator(context.Context, *securityhub.GetFindingAggregatorInput, ...func(*securityhub.Options)) (*securityhub.GetFindingAggregatorOutput, error)
	GetFindings(context.Context, *securityhub.GetFindingsInput, ...func(*securityhub.Options)) (*securityhub.GetFindingsOutput, error)
	GetInsightResults(context.Context, *securityhub.GetInsightResultsInput, ...func(*securityhub.Options)) (*securityhub.GetInsightResultsOutput, error)
	GetInsights(context.Context, *securityhub.GetInsightsInput, ...func(*securityhub.Options)) (*securityhub.GetInsightsOutput, error)
	GetInvitationsCount(context.Context, *securityhub.GetInvitationsCountInput, ...func(*securityhub.Options)) (*securityhub.GetInvitationsCountOutput, error)
	GetMasterAccount(context.Context, *securityhub.GetMasterAccountInput, ...func(*securityhub.Options)) (*securityhub.GetMasterAccountOutput, error)
	GetMembers(context.Context, *securityhub.GetMembersInput, ...func(*securityhub.Options)) (*securityhub.GetMembersOutput, error)
	ListEnabledProductsForImport(context.Context, *securityhub.ListEnabledProductsForImportInput, ...func(*securityhub.Options)) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListFindingAggregators(context.Context, *securityhub.ListFindingAggregatorsInput, ...func(*securityhub.Options)) (*securityhub.ListFindingAggregatorsOutput, error)
	ListInvitations(context.Context, *securityhub.ListInvitationsInput, ...func(*securityhub.Options)) (*securityhub.ListInvitationsOutput, error)
	ListMembers(context.Context, *securityhub.ListMembersInput, ...func(*securityhub.Options)) (*securityhub.ListMembersOutput, error)
	ListOrganizationAdminAccounts(context.Context, *securityhub.ListOrganizationAdminAccountsInput, ...func(*securityhub.Options)) (*securityhub.ListOrganizationAdminAccountsOutput, error)
	ListTagsForResource(context.Context, *securityhub.ListTagsForResourceInput, ...func(*securityhub.Options)) (*securityhub.ListTagsForResourceOutput, error)
}
