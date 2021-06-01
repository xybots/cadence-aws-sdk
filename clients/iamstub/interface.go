// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iamstub

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddClientIDToOpenIDConnectProvider(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddClientIDToOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) *AddClientIDToOpenIDConnectProviderFuture

	AddRoleToInstanceProfile(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)
	AddRoleToInstanceProfileAsync(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) *AddRoleToInstanceProfileFuture

	AddUserToGroup(ctx workflow.Context, input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)
	AddUserToGroupAsync(ctx workflow.Context, input *iam.AddUserToGroupInput) *AddUserToGroupFuture

	AttachGroupPolicy(ctx workflow.Context, input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)
	AttachGroupPolicyAsync(ctx workflow.Context, input *iam.AttachGroupPolicyInput) *AttachGroupPolicyFuture

	AttachRolePolicy(ctx workflow.Context, input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)
	AttachRolePolicyAsync(ctx workflow.Context, input *iam.AttachRolePolicyInput) *AttachRolePolicyFuture

	AttachUserPolicy(ctx workflow.Context, input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)
	AttachUserPolicyAsync(ctx workflow.Context, input *iam.AttachUserPolicyInput) *AttachUserPolicyFuture

	ChangePassword(ctx workflow.Context, input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)
	ChangePasswordAsync(ctx workflow.Context, input *iam.ChangePasswordInput) *ChangePasswordFuture

	CreateAccessKey(ctx workflow.Context, input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)
	CreateAccessKeyAsync(ctx workflow.Context, input *iam.CreateAccessKeyInput) *CreateAccessKeyFuture

	CreateAccountAlias(ctx workflow.Context, input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)
	CreateAccountAliasAsync(ctx workflow.Context, input *iam.CreateAccountAliasInput) *CreateAccountAliasFuture

	CreateGroup(ctx workflow.Context, input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *iam.CreateGroupInput) *CreateGroupFuture

	CreateInstanceProfile(ctx workflow.Context, input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)
	CreateInstanceProfileAsync(ctx workflow.Context, input *iam.CreateInstanceProfileInput) *CreateInstanceProfileFuture

	CreateLoginProfile(ctx workflow.Context, input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)
	CreateLoginProfileAsync(ctx workflow.Context, input *iam.CreateLoginProfileInput) *CreateLoginProfileFuture

	CreateOpenIDConnectProvider(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreateOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) *CreateOpenIDConnectProviderFuture

	CreatePolicy(ctx workflow.Context, input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
	CreatePolicyAsync(ctx workflow.Context, input *iam.CreatePolicyInput) *CreatePolicyFuture

	CreatePolicyVersion(ctx workflow.Context, input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyVersionAsync(ctx workflow.Context, input *iam.CreatePolicyVersionInput) *CreatePolicyVersionFuture

	CreateRole(ctx workflow.Context, input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error)
	CreateRoleAsync(ctx workflow.Context, input *iam.CreateRoleInput) *CreateRoleFuture

	CreateSAMLProvider(ctx workflow.Context, input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)
	CreateSAMLProviderAsync(ctx workflow.Context, input *iam.CreateSAMLProviderInput) *CreateSAMLProviderFuture

	CreateServiceLinkedRole(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleAsync(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) *CreateServiceLinkedRoleFuture

	CreateServiceSpecificCredential(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) *CreateServiceSpecificCredentialFuture

	CreateUser(ctx workflow.Context, input *iam.CreateUserInput) (*iam.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *iam.CreateUserInput) *CreateUserFuture

	CreateVirtualMFADevice(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)
	CreateVirtualMFADeviceAsync(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) *CreateVirtualMFADeviceFuture

	DeactivateMFADevice(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)
	DeactivateMFADeviceAsync(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) *DeactivateMFADeviceFuture

	DeleteAccessKey(ctx workflow.Context, input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccessKeyAsync(ctx workflow.Context, input *iam.DeleteAccessKeyInput) *DeleteAccessKeyFuture

	DeleteAccountAlias(ctx workflow.Context, input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountAliasAsync(ctx workflow.Context, input *iam.DeleteAccountAliasInput) *DeleteAccountAliasFuture

	DeleteAccountPasswordPolicy(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) *DeleteAccountPasswordPolicyFuture

	DeleteGroup(ctx workflow.Context, input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *iam.DeleteGroupInput) *DeleteGroupFuture

	DeleteGroupPolicy(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupPolicyAsync(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) *DeleteGroupPolicyFuture

	DeleteInstanceProfile(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileAsync(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) *DeleteInstanceProfileFuture

	DeleteLoginProfile(ctx workflow.Context, input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)
	DeleteLoginProfileAsync(ctx workflow.Context, input *iam.DeleteLoginProfileInput) *DeleteLoginProfileFuture

	DeleteOpenIDConnectProvider(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeleteOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) *DeleteOpenIDConnectProviderFuture

	DeletePolicy(ctx workflow.Context, input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *iam.DeletePolicyInput) *DeletePolicyFuture

	DeletePolicyVersion(ctx workflow.Context, input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyVersionAsync(ctx workflow.Context, input *iam.DeletePolicyVersionInput) *DeletePolicyVersionFuture

	DeleteRole(ctx workflow.Context, input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
	DeleteRoleAsync(ctx workflow.Context, input *iam.DeleteRoleInput) *DeleteRoleFuture

	DeleteRolePermissionsBoundary(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error)
	DeleteRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) *DeleteRolePermissionsBoundaryFuture

	DeleteRolePolicy(ctx workflow.Context, input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)
	DeleteRolePolicyAsync(ctx workflow.Context, input *iam.DeleteRolePolicyInput) *DeleteRolePolicyFuture

	DeleteSAMLProvider(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderAsync(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) *DeleteSAMLProviderFuture

	DeleteSSHPublicKey(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteSSHPublicKeyAsync(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) *DeleteSSHPublicKeyFuture

	DeleteServerCertificate(ctx workflow.Context, input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)
	DeleteServerCertificateAsync(ctx workflow.Context, input *iam.DeleteServerCertificateInput) *DeleteServerCertificateFuture

	DeleteServiceLinkedRole(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleAsync(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) *DeleteServiceLinkedRoleFuture

	DeleteServiceSpecificCredential(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) *DeleteServiceSpecificCredentialFuture

	DeleteSigningCertificate(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)
	DeleteSigningCertificateAsync(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) *DeleteSigningCertificateFuture

	DeleteUser(ctx workflow.Context, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *iam.DeleteUserInput) *DeleteUserFuture

	DeleteUserPermissionsBoundary(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error)
	DeleteUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) *DeleteUserPermissionsBoundaryFuture

	DeleteUserPolicy(ctx workflow.Context, input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserPolicyAsync(ctx workflow.Context, input *iam.DeleteUserPolicyInput) *DeleteUserPolicyFuture

	DeleteVirtualMFADevice(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)
	DeleteVirtualMFADeviceAsync(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) *DeleteVirtualMFADeviceFuture

	DetachGroupPolicy(ctx workflow.Context, input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)
	DetachGroupPolicyAsync(ctx workflow.Context, input *iam.DetachGroupPolicyInput) *DetachGroupPolicyFuture

	DetachRolePolicy(ctx workflow.Context, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)
	DetachRolePolicyAsync(ctx workflow.Context, input *iam.DetachRolePolicyInput) *DetachRolePolicyFuture

	DetachUserPolicy(ctx workflow.Context, input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)
	DetachUserPolicyAsync(ctx workflow.Context, input *iam.DetachUserPolicyInput) *DetachUserPolicyFuture

	EnableMFADevice(ctx workflow.Context, input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)
	EnableMFADeviceAsync(ctx workflow.Context, input *iam.EnableMFADeviceInput) *EnableMFADeviceFuture

	GenerateCredentialReport(ctx workflow.Context, input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)
	GenerateCredentialReportAsync(ctx workflow.Context, input *iam.GenerateCredentialReportInput) *GenerateCredentialReportFuture

	GenerateOrganizationsAccessReport(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error)
	GenerateOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) *GenerateOrganizationsAccessReportFuture

	GenerateServiceLastAccessedDetails(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error)
	GenerateServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) *GenerateServiceLastAccessedDetailsFuture

	GetAccessKeyLastUsed(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedAsync(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) *GetAccessKeyLastUsedFuture

	GetAccountAuthorizationDetails(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountAuthorizationDetailsAsync(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) *GetAccountAuthorizationDetailsFuture

	GetAccountPasswordPolicy(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) *GetAccountPasswordPolicyFuture

	GetAccountSummary(ctx workflow.Context, input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)
	GetAccountSummaryAsync(ctx workflow.Context, input *iam.GetAccountSummaryInput) *GetAccountSummaryFuture

	GetContextKeysForCustomPolicy(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForCustomPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) *GetContextKeysForCustomPolicyFuture

	GetContextKeysForPrincipalPolicy(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) *GetContextKeysForPrincipalPolicyFuture

	GetCredentialReport(ctx workflow.Context, input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)
	GetCredentialReportAsync(ctx workflow.Context, input *iam.GetCredentialReportInput) *GetCredentialReportFuture

	GetGroup(ctx workflow.Context, input *iam.GetGroupInput) (*iam.GetGroupOutput, error)
	GetGroupAsync(ctx workflow.Context, input *iam.GetGroupInput) *GetGroupFuture

	GetGroupPolicy(ctx workflow.Context, input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)
	GetGroupPolicyAsync(ctx workflow.Context, input *iam.GetGroupPolicyInput) *GetGroupPolicyFuture

	GetInstanceProfile(ctx workflow.Context, input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)
	GetInstanceProfileAsync(ctx workflow.Context, input *iam.GetInstanceProfileInput) *GetInstanceProfileFuture

	GetLoginProfile(ctx workflow.Context, input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)
	GetLoginProfileAsync(ctx workflow.Context, input *iam.GetLoginProfileInput) *GetLoginProfileFuture

	GetOpenIDConnectProvider(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) *GetOpenIDConnectProviderFuture

	GetOrganizationsAccessReport(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error)
	GetOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) *GetOrganizationsAccessReportFuture

	GetPolicy(ctx workflow.Context, input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *iam.GetPolicyInput) *GetPolicyFuture

	GetPolicyVersion(ctx workflow.Context, input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)
	GetPolicyVersionAsync(ctx workflow.Context, input *iam.GetPolicyVersionInput) *GetPolicyVersionFuture

	GetRole(ctx workflow.Context, input *iam.GetRoleInput) (*iam.GetRoleOutput, error)
	GetRoleAsync(ctx workflow.Context, input *iam.GetRoleInput) *GetRoleFuture

	GetRolePolicy(ctx workflow.Context, input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)
	GetRolePolicyAsync(ctx workflow.Context, input *iam.GetRolePolicyInput) *GetRolePolicyFuture

	GetSAMLProvider(ctx workflow.Context, input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)
	GetSAMLProviderAsync(ctx workflow.Context, input *iam.GetSAMLProviderInput) *GetSAMLProviderFuture

	GetSSHPublicKey(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)
	GetSSHPublicKeyAsync(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) *GetSSHPublicKeyFuture

	GetServerCertificate(ctx workflow.Context, input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)
	GetServerCertificateAsync(ctx workflow.Context, input *iam.GetServerCertificateInput) *GetServerCertificateFuture

	GetServiceLastAccessedDetails(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error)
	GetServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) *GetServiceLastAccessedDetailsFuture

	GetServiceLastAccessedDetailsWithEntities(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
	GetServiceLastAccessedDetailsWithEntitiesAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) *GetServiceLastAccessedDetailsWithEntitiesFuture

	GetServiceLinkedRoleDeletionStatus(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetServiceLinkedRoleDeletionStatusAsync(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) *GetServiceLinkedRoleDeletionStatusFuture

	GetUser(ctx workflow.Context, input *iam.GetUserInput) (*iam.GetUserOutput, error)
	GetUserAsync(ctx workflow.Context, input *iam.GetUserInput) *GetUserFuture

	GetUserPolicy(ctx workflow.Context, input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)
	GetUserPolicyAsync(ctx workflow.Context, input *iam.GetUserPolicyInput) *GetUserPolicyFuture

	ListAccessKeys(ctx workflow.Context, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
	ListAccessKeysAsync(ctx workflow.Context, input *iam.ListAccessKeysInput) *ListAccessKeysFuture

	ListAccountAliases(ctx workflow.Context, input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)
	ListAccountAliasesAsync(ctx workflow.Context, input *iam.ListAccountAliasesInput) *ListAccountAliasesFuture

	ListAttachedGroupPolicies(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedGroupPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) *ListAttachedGroupPoliciesFuture

	ListAttachedRolePolicies(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesAsync(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) *ListAttachedRolePoliciesFuture

	ListAttachedUserPolicies(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) *ListAttachedUserPoliciesFuture

	ListEntitiesForPolicy(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyAsync(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) *ListEntitiesForPolicyFuture

	ListGroupPolicies(ctx workflow.Context, input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)
	ListGroupPoliciesAsync(ctx workflow.Context, input *iam.ListGroupPoliciesInput) *ListGroupPoliciesFuture

	ListGroups(ctx workflow.Context, input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *iam.ListGroupsInput) *ListGroupsFuture

	ListGroupsForUser(ctx workflow.Context, input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)
	ListGroupsForUserAsync(ctx workflow.Context, input *iam.ListGroupsForUserInput) *ListGroupsForUserFuture

	ListInstanceProfiles(ctx workflow.Context, input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesAsync(ctx workflow.Context, input *iam.ListInstanceProfilesInput) *ListInstanceProfilesFuture

	ListInstanceProfilesForRole(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesForRoleAsync(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) *ListInstanceProfilesForRoleFuture

	ListMFADevices(ctx workflow.Context, input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)
	ListMFADevicesAsync(ctx workflow.Context, input *iam.ListMFADevicesInput) *ListMFADevicesFuture

	ListOpenIDConnectProviders(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListOpenIDConnectProvidersAsync(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) *ListOpenIDConnectProvidersFuture

	ListPolicies(ctx workflow.Context, input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)
	ListPoliciesAsync(ctx workflow.Context, input *iam.ListPoliciesInput) *ListPoliciesFuture

	ListPoliciesGrantingServiceAccess(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
	ListPoliciesGrantingServiceAccessAsync(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) *ListPoliciesGrantingServiceAccessFuture

	ListPolicyVersions(ctx workflow.Context, input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)
	ListPolicyVersionsAsync(ctx workflow.Context, input *iam.ListPolicyVersionsInput) *ListPolicyVersionsFuture

	ListRolePolicies(ctx workflow.Context, input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)
	ListRolePoliciesAsync(ctx workflow.Context, input *iam.ListRolePoliciesInput) *ListRolePoliciesFuture

	ListRoleTags(ctx workflow.Context, input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error)
	ListRoleTagsAsync(ctx workflow.Context, input *iam.ListRoleTagsInput) *ListRoleTagsFuture

	ListRoles(ctx workflow.Context, input *iam.ListRolesInput) (*iam.ListRolesOutput, error)
	ListRolesAsync(ctx workflow.Context, input *iam.ListRolesInput) *ListRolesFuture

	ListSAMLProviders(ctx workflow.Context, input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)
	ListSAMLProvidersAsync(ctx workflow.Context, input *iam.ListSAMLProvidersInput) *ListSAMLProvidersFuture

	ListSSHPublicKeys(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)
	ListSSHPublicKeysAsync(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) *ListSSHPublicKeysFuture

	ListServerCertificates(ctx workflow.Context, input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
	ListServerCertificatesAsync(ctx workflow.Context, input *iam.ListServerCertificatesInput) *ListServerCertificatesFuture

	ListServiceSpecificCredentials(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListServiceSpecificCredentialsAsync(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) *ListServiceSpecificCredentialsFuture

	ListSigningCertificates(ctx workflow.Context, input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)
	ListSigningCertificatesAsync(ctx workflow.Context, input *iam.ListSigningCertificatesInput) *ListSigningCertificatesFuture

	ListUserPolicies(ctx workflow.Context, input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)
	ListUserPoliciesAsync(ctx workflow.Context, input *iam.ListUserPoliciesInput) *ListUserPoliciesFuture

	ListUserTags(ctx workflow.Context, input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error)
	ListUserTagsAsync(ctx workflow.Context, input *iam.ListUserTagsInput) *ListUserTagsFuture

	ListUsers(ctx workflow.Context, input *iam.ListUsersInput) (*iam.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *iam.ListUsersInput) *ListUsersFuture

	ListVirtualMFADevices(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)
	ListVirtualMFADevicesAsync(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) *ListVirtualMFADevicesFuture

	PutGroupPolicy(ctx workflow.Context, input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)
	PutGroupPolicyAsync(ctx workflow.Context, input *iam.PutGroupPolicyInput) *PutGroupPolicyFuture

	PutRolePermissionsBoundary(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error)
	PutRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) *PutRolePermissionsBoundaryFuture

	PutRolePolicy(ctx workflow.Context, input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)
	PutRolePolicyAsync(ctx workflow.Context, input *iam.PutRolePolicyInput) *PutRolePolicyFuture

	PutUserPermissionsBoundary(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error)
	PutUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) *PutUserPermissionsBoundaryFuture

	PutUserPolicy(ctx workflow.Context, input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)
	PutUserPolicyAsync(ctx workflow.Context, input *iam.PutUserPolicyInput) *PutUserPolicyFuture

	RemoveClientIDFromOpenIDConnectProvider(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveClientIDFromOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) *RemoveClientIDFromOpenIDConnectProviderFuture

	RemoveRoleFromInstanceProfile(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfileAsync(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) *RemoveRoleFromInstanceProfileFuture

	RemoveUserFromGroup(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupAsync(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) *RemoveUserFromGroupFuture

	ResetServiceSpecificCredential(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResetServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) *ResetServiceSpecificCredentialFuture

	ResyncMFADevice(ctx workflow.Context, input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)
	ResyncMFADeviceAsync(ctx workflow.Context, input *iam.ResyncMFADeviceInput) *ResyncMFADeviceFuture

	SetDefaultPolicyVersion(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionAsync(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) *SetDefaultPolicyVersionFuture

	SetSecurityTokenServicePreferences(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error)
	SetSecurityTokenServicePreferencesAsync(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) *SetSecurityTokenServicePreferencesFuture

	SimulateCustomPolicy(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulateCustomPolicyAsync(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) *SimulateCustomPolicyFuture

	SimulatePrincipalPolicy(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyAsync(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) *SimulatePrincipalPolicyFuture

	TagRole(ctx workflow.Context, input *iam.TagRoleInput) (*iam.TagRoleOutput, error)
	TagRoleAsync(ctx workflow.Context, input *iam.TagRoleInput) *TagRoleFuture

	TagUser(ctx workflow.Context, input *iam.TagUserInput) (*iam.TagUserOutput, error)
	TagUserAsync(ctx workflow.Context, input *iam.TagUserInput) *TagUserFuture

	UntagRole(ctx workflow.Context, input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error)
	UntagRoleAsync(ctx workflow.Context, input *iam.UntagRoleInput) *UntagRoleFuture

	UntagUser(ctx workflow.Context, input *iam.UntagUserInput) (*iam.UntagUserOutput, error)
	UntagUserAsync(ctx workflow.Context, input *iam.UntagUserInput) *UntagUserFuture

	UpdateAccessKey(ctx workflow.Context, input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccessKeyAsync(ctx workflow.Context, input *iam.UpdateAccessKeyInput) *UpdateAccessKeyFuture

	UpdateAccountPasswordPolicy(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) *UpdateAccountPasswordPolicyFuture

	UpdateAssumeRolePolicy(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateAssumeRolePolicyAsync(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) *UpdateAssumeRolePolicyFuture

	UpdateGroup(ctx workflow.Context, input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *iam.UpdateGroupInput) *UpdateGroupFuture

	UpdateLoginProfile(ctx workflow.Context, input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)
	UpdateLoginProfileAsync(ctx workflow.Context, input *iam.UpdateLoginProfileInput) *UpdateLoginProfileFuture

	UpdateOpenIDConnectProviderThumbprint(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateOpenIDConnectProviderThumbprintAsync(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) *UpdateOpenIDConnectProviderThumbprintFuture

	UpdateRole(ctx workflow.Context, input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error)
	UpdateRoleAsync(ctx workflow.Context, input *iam.UpdateRoleInput) *UpdateRoleFuture

	UpdateRoleDescription(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateRoleDescriptionAsync(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) *UpdateRoleDescriptionFuture

	UpdateSAMLProvider(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderAsync(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) *UpdateSAMLProviderFuture

	UpdateSSHPublicKey(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateSSHPublicKeyAsync(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) *UpdateSSHPublicKeyFuture

	UpdateServerCertificate(ctx workflow.Context, input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)
	UpdateServerCertificateAsync(ctx workflow.Context, input *iam.UpdateServerCertificateInput) *UpdateServerCertificateFuture

	UpdateServiceSpecificCredential(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) *UpdateServiceSpecificCredentialFuture

	UpdateSigningCertificate(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)
	UpdateSigningCertificateAsync(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) *UpdateSigningCertificateFuture

	UpdateUser(ctx workflow.Context, input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *iam.UpdateUserInput) *UpdateUserFuture

	UploadSSHPublicKey(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)
	UploadSSHPublicKeyAsync(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) *UploadSSHPublicKeyFuture

	UploadServerCertificate(ctx workflow.Context, input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)
	UploadServerCertificateAsync(ctx workflow.Context, input *iam.UploadServerCertificateInput) *UploadServerCertificateFuture

	UploadSigningCertificate(ctx workflow.Context, input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
	UploadSigningCertificateAsync(ctx workflow.Context, input *iam.UploadSigningCertificateInput) *UploadSigningCertificateFuture

	WaitUntilInstanceProfileExists(ctx workflow.Context, input *iam.GetInstanceProfileInput) error
	WaitUntilInstanceProfileExistsAsync(ctx workflow.Context, input *iam.GetInstanceProfileInput) *clients.VoidFuture

	WaitUntilPolicyExists(ctx workflow.Context, input *iam.GetPolicyInput) error
	WaitUntilPolicyExistsAsync(ctx workflow.Context, input *iam.GetPolicyInput) *clients.VoidFuture

	WaitUntilRoleExists(ctx workflow.Context, input *iam.GetRoleInput) error
	WaitUntilRoleExistsAsync(ctx workflow.Context, input *iam.GetRoleInput) *clients.VoidFuture

	WaitUntilUserExists(ctx workflow.Context, input *iam.GetUserInput) error
	WaitUntilUserExistsAsync(ctx workflow.Context, input *iam.GetUserInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
