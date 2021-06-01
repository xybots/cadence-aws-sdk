// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package connectstub

import (
	"github.com/aws/aws-sdk-go/service/connect"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateApprovedOrigin(ctx workflow.Context, input *connect.AssociateApprovedOriginInput) (*connect.AssociateApprovedOriginOutput, error)
	AssociateApprovedOriginAsync(ctx workflow.Context, input *connect.AssociateApprovedOriginInput) *AssociateApprovedOriginFuture

	AssociateInstanceStorageConfig(ctx workflow.Context, input *connect.AssociateInstanceStorageConfigInput) (*connect.AssociateInstanceStorageConfigOutput, error)
	AssociateInstanceStorageConfigAsync(ctx workflow.Context, input *connect.AssociateInstanceStorageConfigInput) *AssociateInstanceStorageConfigFuture

	AssociateLambdaFunction(ctx workflow.Context, input *connect.AssociateLambdaFunctionInput) (*connect.AssociateLambdaFunctionOutput, error)
	AssociateLambdaFunctionAsync(ctx workflow.Context, input *connect.AssociateLambdaFunctionInput) *AssociateLambdaFunctionFuture

	AssociateLexBot(ctx workflow.Context, input *connect.AssociateLexBotInput) (*connect.AssociateLexBotOutput, error)
	AssociateLexBotAsync(ctx workflow.Context, input *connect.AssociateLexBotInput) *AssociateLexBotFuture

	AssociateQueueQuickConnects(ctx workflow.Context, input *connect.AssociateQueueQuickConnectsInput) (*connect.AssociateQueueQuickConnectsOutput, error)
	AssociateQueueQuickConnectsAsync(ctx workflow.Context, input *connect.AssociateQueueQuickConnectsInput) *AssociateQueueQuickConnectsFuture

	AssociateRoutingProfileQueues(ctx workflow.Context, input *connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error)
	AssociateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.AssociateRoutingProfileQueuesInput) *AssociateRoutingProfileQueuesFuture

	AssociateSecurityKey(ctx workflow.Context, input *connect.AssociateSecurityKeyInput) (*connect.AssociateSecurityKeyOutput, error)
	AssociateSecurityKeyAsync(ctx workflow.Context, input *connect.AssociateSecurityKeyInput) *AssociateSecurityKeyFuture

	CreateContactFlow(ctx workflow.Context, input *connect.CreateContactFlowInput) (*connect.CreateContactFlowOutput, error)
	CreateContactFlowAsync(ctx workflow.Context, input *connect.CreateContactFlowInput) *CreateContactFlowFuture

	CreateInstance(ctx workflow.Context, input *connect.CreateInstanceInput) (*connect.CreateInstanceOutput, error)
	CreateInstanceAsync(ctx workflow.Context, input *connect.CreateInstanceInput) *CreateInstanceFuture

	CreateIntegrationAssociation(ctx workflow.Context, input *connect.CreateIntegrationAssociationInput) (*connect.CreateIntegrationAssociationOutput, error)
	CreateIntegrationAssociationAsync(ctx workflow.Context, input *connect.CreateIntegrationAssociationInput) *CreateIntegrationAssociationFuture

	CreateQueue(ctx workflow.Context, input *connect.CreateQueueInput) (*connect.CreateQueueOutput, error)
	CreateQueueAsync(ctx workflow.Context, input *connect.CreateQueueInput) *CreateQueueFuture

	CreateQuickConnect(ctx workflow.Context, input *connect.CreateQuickConnectInput) (*connect.CreateQuickConnectOutput, error)
	CreateQuickConnectAsync(ctx workflow.Context, input *connect.CreateQuickConnectInput) *CreateQuickConnectFuture

	CreateRoutingProfile(ctx workflow.Context, input *connect.CreateRoutingProfileInput) (*connect.CreateRoutingProfileOutput, error)
	CreateRoutingProfileAsync(ctx workflow.Context, input *connect.CreateRoutingProfileInput) *CreateRoutingProfileFuture

	CreateUseCase(ctx workflow.Context, input *connect.CreateUseCaseInput) (*connect.CreateUseCaseOutput, error)
	CreateUseCaseAsync(ctx workflow.Context, input *connect.CreateUseCaseInput) *CreateUseCaseFuture

	CreateUser(ctx workflow.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *connect.CreateUserInput) *CreateUserFuture

	CreateUserHierarchyGroup(ctx workflow.Context, input *connect.CreateUserHierarchyGroupInput) (*connect.CreateUserHierarchyGroupOutput, error)
	CreateUserHierarchyGroupAsync(ctx workflow.Context, input *connect.CreateUserHierarchyGroupInput) *CreateUserHierarchyGroupFuture

	DeleteInstance(ctx workflow.Context, input *connect.DeleteInstanceInput) (*connect.DeleteInstanceOutput, error)
	DeleteInstanceAsync(ctx workflow.Context, input *connect.DeleteInstanceInput) *DeleteInstanceFuture

	DeleteIntegrationAssociation(ctx workflow.Context, input *connect.DeleteIntegrationAssociationInput) (*connect.DeleteIntegrationAssociationOutput, error)
	DeleteIntegrationAssociationAsync(ctx workflow.Context, input *connect.DeleteIntegrationAssociationInput) *DeleteIntegrationAssociationFuture

	DeleteQuickConnect(ctx workflow.Context, input *connect.DeleteQuickConnectInput) (*connect.DeleteQuickConnectOutput, error)
	DeleteQuickConnectAsync(ctx workflow.Context, input *connect.DeleteQuickConnectInput) *DeleteQuickConnectFuture

	DeleteUseCase(ctx workflow.Context, input *connect.DeleteUseCaseInput) (*connect.DeleteUseCaseOutput, error)
	DeleteUseCaseAsync(ctx workflow.Context, input *connect.DeleteUseCaseInput) *DeleteUseCaseFuture

	DeleteUser(ctx workflow.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *connect.DeleteUserInput) *DeleteUserFuture

	DeleteUserHierarchyGroup(ctx workflow.Context, input *connect.DeleteUserHierarchyGroupInput) (*connect.DeleteUserHierarchyGroupOutput, error)
	DeleteUserHierarchyGroupAsync(ctx workflow.Context, input *connect.DeleteUserHierarchyGroupInput) *DeleteUserHierarchyGroupFuture

	DescribeContactFlow(ctx workflow.Context, input *connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error)
	DescribeContactFlowAsync(ctx workflow.Context, input *connect.DescribeContactFlowInput) *DescribeContactFlowFuture

	DescribeHoursOfOperation(ctx workflow.Context, input *connect.DescribeHoursOfOperationInput) (*connect.DescribeHoursOfOperationOutput, error)
	DescribeHoursOfOperationAsync(ctx workflow.Context, input *connect.DescribeHoursOfOperationInput) *DescribeHoursOfOperationFuture

	DescribeInstance(ctx workflow.Context, input *connect.DescribeInstanceInput) (*connect.DescribeInstanceOutput, error)
	DescribeInstanceAsync(ctx workflow.Context, input *connect.DescribeInstanceInput) *DescribeInstanceFuture

	DescribeInstanceAttribute(ctx workflow.Context, input *connect.DescribeInstanceAttributeInput) (*connect.DescribeInstanceAttributeOutput, error)
	DescribeInstanceAttributeAsync(ctx workflow.Context, input *connect.DescribeInstanceAttributeInput) *DescribeInstanceAttributeFuture

	DescribeInstanceStorageConfig(ctx workflow.Context, input *connect.DescribeInstanceStorageConfigInput) (*connect.DescribeInstanceStorageConfigOutput, error)
	DescribeInstanceStorageConfigAsync(ctx workflow.Context, input *connect.DescribeInstanceStorageConfigInput) *DescribeInstanceStorageConfigFuture

	DescribeQueue(ctx workflow.Context, input *connect.DescribeQueueInput) (*connect.DescribeQueueOutput, error)
	DescribeQueueAsync(ctx workflow.Context, input *connect.DescribeQueueInput) *DescribeQueueFuture

	DescribeQuickConnect(ctx workflow.Context, input *connect.DescribeQuickConnectInput) (*connect.DescribeQuickConnectOutput, error)
	DescribeQuickConnectAsync(ctx workflow.Context, input *connect.DescribeQuickConnectInput) *DescribeQuickConnectFuture

	DescribeRoutingProfile(ctx workflow.Context, input *connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error)
	DescribeRoutingProfileAsync(ctx workflow.Context, input *connect.DescribeRoutingProfileInput) *DescribeRoutingProfileFuture

	DescribeUser(ctx workflow.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error)
	DescribeUserAsync(ctx workflow.Context, input *connect.DescribeUserInput) *DescribeUserFuture

	DescribeUserHierarchyGroup(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) *DescribeUserHierarchyGroupFuture

	DescribeUserHierarchyStructure(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) *DescribeUserHierarchyStructureFuture

	DisassociateApprovedOrigin(ctx workflow.Context, input *connect.DisassociateApprovedOriginInput) (*connect.DisassociateApprovedOriginOutput, error)
	DisassociateApprovedOriginAsync(ctx workflow.Context, input *connect.DisassociateApprovedOriginInput) *DisassociateApprovedOriginFuture

	DisassociateInstanceStorageConfig(ctx workflow.Context, input *connect.DisassociateInstanceStorageConfigInput) (*connect.DisassociateInstanceStorageConfigOutput, error)
	DisassociateInstanceStorageConfigAsync(ctx workflow.Context, input *connect.DisassociateInstanceStorageConfigInput) *DisassociateInstanceStorageConfigFuture

	DisassociateLambdaFunction(ctx workflow.Context, input *connect.DisassociateLambdaFunctionInput) (*connect.DisassociateLambdaFunctionOutput, error)
	DisassociateLambdaFunctionAsync(ctx workflow.Context, input *connect.DisassociateLambdaFunctionInput) *DisassociateLambdaFunctionFuture

	DisassociateLexBot(ctx workflow.Context, input *connect.DisassociateLexBotInput) (*connect.DisassociateLexBotOutput, error)
	DisassociateLexBotAsync(ctx workflow.Context, input *connect.DisassociateLexBotInput) *DisassociateLexBotFuture

	DisassociateQueueQuickConnects(ctx workflow.Context, input *connect.DisassociateQueueQuickConnectsInput) (*connect.DisassociateQueueQuickConnectsOutput, error)
	DisassociateQueueQuickConnectsAsync(ctx workflow.Context, input *connect.DisassociateQueueQuickConnectsInput) *DisassociateQueueQuickConnectsFuture

	DisassociateRoutingProfileQueues(ctx workflow.Context, input *connect.DisassociateRoutingProfileQueuesInput) (*connect.DisassociateRoutingProfileQueuesOutput, error)
	DisassociateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.DisassociateRoutingProfileQueuesInput) *DisassociateRoutingProfileQueuesFuture

	DisassociateSecurityKey(ctx workflow.Context, input *connect.DisassociateSecurityKeyInput) (*connect.DisassociateSecurityKeyOutput, error)
	DisassociateSecurityKeyAsync(ctx workflow.Context, input *connect.DisassociateSecurityKeyInput) *DisassociateSecurityKeyFuture

	GetContactAttributes(ctx workflow.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesAsync(ctx workflow.Context, input *connect.GetContactAttributesInput) *GetContactAttributesFuture

	GetCurrentMetricData(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataAsync(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) *GetCurrentMetricDataFuture

	GetFederationToken(ctx workflow.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenAsync(ctx workflow.Context, input *connect.GetFederationTokenInput) *GetFederationTokenFuture

	GetMetricData(ctx workflow.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error)
	GetMetricDataAsync(ctx workflow.Context, input *connect.GetMetricDataInput) *GetMetricDataFuture

	ListApprovedOrigins(ctx workflow.Context, input *connect.ListApprovedOriginsInput) (*connect.ListApprovedOriginsOutput, error)
	ListApprovedOriginsAsync(ctx workflow.Context, input *connect.ListApprovedOriginsInput) *ListApprovedOriginsFuture

	ListContactFlows(ctx workflow.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error)
	ListContactFlowsAsync(ctx workflow.Context, input *connect.ListContactFlowsInput) *ListContactFlowsFuture

	ListHoursOfOperations(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error)
	ListHoursOfOperationsAsync(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) *ListHoursOfOperationsFuture

	ListInstanceAttributes(ctx workflow.Context, input *connect.ListInstanceAttributesInput) (*connect.ListInstanceAttributesOutput, error)
	ListInstanceAttributesAsync(ctx workflow.Context, input *connect.ListInstanceAttributesInput) *ListInstanceAttributesFuture

	ListInstanceStorageConfigs(ctx workflow.Context, input *connect.ListInstanceStorageConfigsInput) (*connect.ListInstanceStorageConfigsOutput, error)
	ListInstanceStorageConfigsAsync(ctx workflow.Context, input *connect.ListInstanceStorageConfigsInput) *ListInstanceStorageConfigsFuture

	ListInstances(ctx workflow.Context, input *connect.ListInstancesInput) (*connect.ListInstancesOutput, error)
	ListInstancesAsync(ctx workflow.Context, input *connect.ListInstancesInput) *ListInstancesFuture

	ListIntegrationAssociations(ctx workflow.Context, input *connect.ListIntegrationAssociationsInput) (*connect.ListIntegrationAssociationsOutput, error)
	ListIntegrationAssociationsAsync(ctx workflow.Context, input *connect.ListIntegrationAssociationsInput) *ListIntegrationAssociationsFuture

	ListLambdaFunctions(ctx workflow.Context, input *connect.ListLambdaFunctionsInput) (*connect.ListLambdaFunctionsOutput, error)
	ListLambdaFunctionsAsync(ctx workflow.Context, input *connect.ListLambdaFunctionsInput) *ListLambdaFunctionsFuture

	ListLexBots(ctx workflow.Context, input *connect.ListLexBotsInput) (*connect.ListLexBotsOutput, error)
	ListLexBotsAsync(ctx workflow.Context, input *connect.ListLexBotsInput) *ListLexBotsFuture

	ListPhoneNumbers(ctx workflow.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error)
	ListPhoneNumbersAsync(ctx workflow.Context, input *connect.ListPhoneNumbersInput) *ListPhoneNumbersFuture

	ListPrompts(ctx workflow.Context, input *connect.ListPromptsInput) (*connect.ListPromptsOutput, error)
	ListPromptsAsync(ctx workflow.Context, input *connect.ListPromptsInput) *ListPromptsFuture

	ListQueueQuickConnects(ctx workflow.Context, input *connect.ListQueueQuickConnectsInput) (*connect.ListQueueQuickConnectsOutput, error)
	ListQueueQuickConnectsAsync(ctx workflow.Context, input *connect.ListQueueQuickConnectsInput) *ListQueueQuickConnectsFuture

	ListQueues(ctx workflow.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error)
	ListQueuesAsync(ctx workflow.Context, input *connect.ListQueuesInput) *ListQueuesFuture

	ListQuickConnects(ctx workflow.Context, input *connect.ListQuickConnectsInput) (*connect.ListQuickConnectsOutput, error)
	ListQuickConnectsAsync(ctx workflow.Context, input *connect.ListQuickConnectsInput) *ListQuickConnectsFuture

	ListRoutingProfileQueues(ctx workflow.Context, input *connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error)
	ListRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.ListRoutingProfileQueuesInput) *ListRoutingProfileQueuesFuture

	ListRoutingProfiles(ctx workflow.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesAsync(ctx workflow.Context, input *connect.ListRoutingProfilesInput) *ListRoutingProfilesFuture

	ListSecurityKeys(ctx workflow.Context, input *connect.ListSecurityKeysInput) (*connect.ListSecurityKeysOutput, error)
	ListSecurityKeysAsync(ctx workflow.Context, input *connect.ListSecurityKeysInput) *ListSecurityKeysFuture

	ListSecurityProfiles(ctx workflow.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesAsync(ctx workflow.Context, input *connect.ListSecurityProfilesInput) *ListSecurityProfilesFuture

	ListTagsForResource(ctx workflow.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *connect.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListUseCases(ctx workflow.Context, input *connect.ListUseCasesInput) (*connect.ListUseCasesOutput, error)
	ListUseCasesAsync(ctx workflow.Context, input *connect.ListUseCasesInput) *ListUseCasesFuture

	ListUserHierarchyGroups(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsAsync(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) *ListUserHierarchyGroupsFuture

	ListUsers(ctx workflow.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *connect.ListUsersInput) *ListUsersFuture

	ResumeContactRecording(ctx workflow.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error)
	ResumeContactRecordingAsync(ctx workflow.Context, input *connect.ResumeContactRecordingInput) *ResumeContactRecordingFuture

	StartChatContact(ctx workflow.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error)
	StartChatContactAsync(ctx workflow.Context, input *connect.StartChatContactInput) *StartChatContactFuture

	StartContactRecording(ctx workflow.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error)
	StartContactRecordingAsync(ctx workflow.Context, input *connect.StartContactRecordingInput) *StartContactRecordingFuture

	StartOutboundVoiceContact(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactAsync(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) *StartOutboundVoiceContactFuture

	StartTaskContact(ctx workflow.Context, input *connect.StartTaskContactInput) (*connect.StartTaskContactOutput, error)
	StartTaskContactAsync(ctx workflow.Context, input *connect.StartTaskContactInput) *StartTaskContactFuture

	StopContact(ctx workflow.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error)
	StopContactAsync(ctx workflow.Context, input *connect.StopContactInput) *StopContactFuture

	StopContactRecording(ctx workflow.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error)
	StopContactRecordingAsync(ctx workflow.Context, input *connect.StopContactRecordingInput) *StopContactRecordingFuture

	SuspendContactRecording(ctx workflow.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error)
	SuspendContactRecordingAsync(ctx workflow.Context, input *connect.SuspendContactRecordingInput) *SuspendContactRecordingFuture

	TagResource(ctx workflow.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *connect.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *connect.UntagResourceInput) *UntagResourceFuture

	UpdateContactAttributes(ctx workflow.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesAsync(ctx workflow.Context, input *connect.UpdateContactAttributesInput) *UpdateContactAttributesFuture

	UpdateContactFlowContent(ctx workflow.Context, input *connect.UpdateContactFlowContentInput) (*connect.UpdateContactFlowContentOutput, error)
	UpdateContactFlowContentAsync(ctx workflow.Context, input *connect.UpdateContactFlowContentInput) *UpdateContactFlowContentFuture

	UpdateContactFlowName(ctx workflow.Context, input *connect.UpdateContactFlowNameInput) (*connect.UpdateContactFlowNameOutput, error)
	UpdateContactFlowNameAsync(ctx workflow.Context, input *connect.UpdateContactFlowNameInput) *UpdateContactFlowNameFuture

	UpdateInstanceAttribute(ctx workflow.Context, input *connect.UpdateInstanceAttributeInput) (*connect.UpdateInstanceAttributeOutput, error)
	UpdateInstanceAttributeAsync(ctx workflow.Context, input *connect.UpdateInstanceAttributeInput) *UpdateInstanceAttributeFuture

	UpdateInstanceStorageConfig(ctx workflow.Context, input *connect.UpdateInstanceStorageConfigInput) (*connect.UpdateInstanceStorageConfigOutput, error)
	UpdateInstanceStorageConfigAsync(ctx workflow.Context, input *connect.UpdateInstanceStorageConfigInput) *UpdateInstanceStorageConfigFuture

	UpdateQueueHoursOfOperation(ctx workflow.Context, input *connect.UpdateQueueHoursOfOperationInput) (*connect.UpdateQueueHoursOfOperationOutput, error)
	UpdateQueueHoursOfOperationAsync(ctx workflow.Context, input *connect.UpdateQueueHoursOfOperationInput) *UpdateQueueHoursOfOperationFuture

	UpdateQueueMaxContacts(ctx workflow.Context, input *connect.UpdateQueueMaxContactsInput) (*connect.UpdateQueueMaxContactsOutput, error)
	UpdateQueueMaxContactsAsync(ctx workflow.Context, input *connect.UpdateQueueMaxContactsInput) *UpdateQueueMaxContactsFuture

	UpdateQueueName(ctx workflow.Context, input *connect.UpdateQueueNameInput) (*connect.UpdateQueueNameOutput, error)
	UpdateQueueNameAsync(ctx workflow.Context, input *connect.UpdateQueueNameInput) *UpdateQueueNameFuture

	UpdateQueueOutboundCallerConfig(ctx workflow.Context, input *connect.UpdateQueueOutboundCallerConfigInput) (*connect.UpdateQueueOutboundCallerConfigOutput, error)
	UpdateQueueOutboundCallerConfigAsync(ctx workflow.Context, input *connect.UpdateQueueOutboundCallerConfigInput) *UpdateQueueOutboundCallerConfigFuture

	UpdateQueueStatus(ctx workflow.Context, input *connect.UpdateQueueStatusInput) (*connect.UpdateQueueStatusOutput, error)
	UpdateQueueStatusAsync(ctx workflow.Context, input *connect.UpdateQueueStatusInput) *UpdateQueueStatusFuture

	UpdateQuickConnectConfig(ctx workflow.Context, input *connect.UpdateQuickConnectConfigInput) (*connect.UpdateQuickConnectConfigOutput, error)
	UpdateQuickConnectConfigAsync(ctx workflow.Context, input *connect.UpdateQuickConnectConfigInput) *UpdateQuickConnectConfigFuture

	UpdateQuickConnectName(ctx workflow.Context, input *connect.UpdateQuickConnectNameInput) (*connect.UpdateQuickConnectNameOutput, error)
	UpdateQuickConnectNameAsync(ctx workflow.Context, input *connect.UpdateQuickConnectNameInput) *UpdateQuickConnectNameFuture

	UpdateRoutingProfileConcurrency(ctx workflow.Context, input *connect.UpdateRoutingProfileConcurrencyInput) (*connect.UpdateRoutingProfileConcurrencyOutput, error)
	UpdateRoutingProfileConcurrencyAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileConcurrencyInput) *UpdateRoutingProfileConcurrencyFuture

	UpdateRoutingProfileDefaultOutboundQueue(ctx workflow.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error)
	UpdateRoutingProfileDefaultOutboundQueueAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) *UpdateRoutingProfileDefaultOutboundQueueFuture

	UpdateRoutingProfileName(ctx workflow.Context, input *connect.UpdateRoutingProfileNameInput) (*connect.UpdateRoutingProfileNameOutput, error)
	UpdateRoutingProfileNameAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileNameInput) *UpdateRoutingProfileNameFuture

	UpdateRoutingProfileQueues(ctx workflow.Context, input *connect.UpdateRoutingProfileQueuesInput) (*connect.UpdateRoutingProfileQueuesOutput, error)
	UpdateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileQueuesInput) *UpdateRoutingProfileQueuesFuture

	UpdateUserHierarchy(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) *UpdateUserHierarchyFuture

	UpdateUserHierarchyGroupName(ctx workflow.Context, input *connect.UpdateUserHierarchyGroupNameInput) (*connect.UpdateUserHierarchyGroupNameOutput, error)
	UpdateUserHierarchyGroupNameAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyGroupNameInput) *UpdateUserHierarchyGroupNameFuture

	UpdateUserHierarchyStructure(ctx workflow.Context, input *connect.UpdateUserHierarchyStructureInput) (*connect.UpdateUserHierarchyStructureOutput, error)
	UpdateUserHierarchyStructureAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyStructureInput) *UpdateUserHierarchyStructureFuture

	UpdateUserIdentityInfo(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoAsync(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) *UpdateUserIdentityInfoFuture

	UpdateUserPhoneConfig(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigAsync(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) *UpdateUserPhoneConfigFuture

	UpdateUserRoutingProfile(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileAsync(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) *UpdateUserRoutingProfileFuture

	UpdateUserSecurityProfiles(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesAsync(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) *UpdateUserSecurityProfilesFuture
}

func NewClient() Client {
	return &stub{}
}