// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package fmsstub

import (
	"github.com/aws/aws-sdk-go/service/fms"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateAdminAccount(ctx workflow.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error)
	AssociateAdminAccountAsync(ctx workflow.Context, input *fms.AssociateAdminAccountInput) *AssociateAdminAccountFuture

	DeleteAppsList(ctx workflow.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error)
	DeleteAppsListAsync(ctx workflow.Context, input *fms.DeleteAppsListInput) *DeleteAppsListFuture

	DeleteNotificationChannel(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error)
	DeleteNotificationChannelAsync(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) *DeleteNotificationChannelFuture

	DeletePolicy(ctx workflow.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *fms.DeletePolicyInput) *DeletePolicyFuture

	DeleteProtocolsList(ctx workflow.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error)
	DeleteProtocolsListAsync(ctx workflow.Context, input *fms.DeleteProtocolsListInput) *DeleteProtocolsListFuture

	DisassociateAdminAccount(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error)
	DisassociateAdminAccountAsync(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) *DisassociateAdminAccountFuture

	GetAdminAccount(ctx workflow.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error)
	GetAdminAccountAsync(ctx workflow.Context, input *fms.GetAdminAccountInput) *GetAdminAccountFuture

	GetAppsList(ctx workflow.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error)
	GetAppsListAsync(ctx workflow.Context, input *fms.GetAppsListInput) *GetAppsListFuture

	GetComplianceDetail(ctx workflow.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error)
	GetComplianceDetailAsync(ctx workflow.Context, input *fms.GetComplianceDetailInput) *GetComplianceDetailFuture

	GetNotificationChannel(ctx workflow.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error)
	GetNotificationChannelAsync(ctx workflow.Context, input *fms.GetNotificationChannelInput) *GetNotificationChannelFuture

	GetPolicy(ctx workflow.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *fms.GetPolicyInput) *GetPolicyFuture

	GetProtectionStatus(ctx workflow.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error)
	GetProtectionStatusAsync(ctx workflow.Context, input *fms.GetProtectionStatusInput) *GetProtectionStatusFuture

	GetProtocolsList(ctx workflow.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error)
	GetProtocolsListAsync(ctx workflow.Context, input *fms.GetProtocolsListInput) *GetProtocolsListFuture

	GetViolationDetails(ctx workflow.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error)
	GetViolationDetailsAsync(ctx workflow.Context, input *fms.GetViolationDetailsInput) *GetViolationDetailsFuture

	ListAppsLists(ctx workflow.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error)
	ListAppsListsAsync(ctx workflow.Context, input *fms.ListAppsListsInput) *ListAppsListsFuture

	ListComplianceStatus(ctx workflow.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error)
	ListComplianceStatusAsync(ctx workflow.Context, input *fms.ListComplianceStatusInput) *ListComplianceStatusFuture

	ListMemberAccounts(ctx workflow.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error)
	ListMemberAccountsAsync(ctx workflow.Context, input *fms.ListMemberAccountsInput) *ListMemberAccountsFuture

	ListPolicies(ctx workflow.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error)
	ListPoliciesAsync(ctx workflow.Context, input *fms.ListPoliciesInput) *ListPoliciesFuture

	ListProtocolsLists(ctx workflow.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error)
	ListProtocolsListsAsync(ctx workflow.Context, input *fms.ListProtocolsListsInput) *ListProtocolsListsFuture

	ListTagsForResource(ctx workflow.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *fms.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutAppsList(ctx workflow.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error)
	PutAppsListAsync(ctx workflow.Context, input *fms.PutAppsListInput) *PutAppsListFuture

	PutNotificationChannel(ctx workflow.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error)
	PutNotificationChannelAsync(ctx workflow.Context, input *fms.PutNotificationChannelInput) *PutNotificationChannelFuture

	PutPolicy(ctx workflow.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error)
	PutPolicyAsync(ctx workflow.Context, input *fms.PutPolicyInput) *PutPolicyFuture

	PutProtocolsList(ctx workflow.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error)
	PutProtocolsListAsync(ctx workflow.Context, input *fms.PutProtocolsListInput) *PutProtocolsListFuture

	TagResource(ctx workflow.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *fms.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *fms.UntagResourceInput) *UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
