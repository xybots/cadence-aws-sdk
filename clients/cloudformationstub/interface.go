// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudformationstub

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelUpdateStack(ctx workflow.Context, input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)
	CancelUpdateStackAsync(ctx workflow.Context, input *cloudformation.CancelUpdateStackInput) *CancelUpdateStackFuture

	ContinueUpdateRollback(ctx workflow.Context, input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error)
	ContinueUpdateRollbackAsync(ctx workflow.Context, input *cloudformation.ContinueUpdateRollbackInput) *ContinueUpdateRollbackFuture

	CreateChangeSet(ctx workflow.Context, input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)
	CreateChangeSetAsync(ctx workflow.Context, input *cloudformation.CreateChangeSetInput) *CreateChangeSetFuture

	CreateStack(ctx workflow.Context, input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
	CreateStackAsync(ctx workflow.Context, input *cloudformation.CreateStackInput) *CreateStackFuture

	CreateStackInstances(ctx workflow.Context, input *cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackInstancesAsync(ctx workflow.Context, input *cloudformation.CreateStackInstancesInput) *CreateStackInstancesFuture

	CreateStackSet(ctx workflow.Context, input *cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error)
	CreateStackSetAsync(ctx workflow.Context, input *cloudformation.CreateStackSetInput) *CreateStackSetFuture

	DeleteChangeSet(ctx workflow.Context, input *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteChangeSetAsync(ctx workflow.Context, input *cloudformation.DeleteChangeSetInput) *DeleteChangeSetFuture

	DeleteStack(ctx workflow.Context, input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)
	DeleteStackAsync(ctx workflow.Context, input *cloudformation.DeleteStackInput) *DeleteStackFuture

	DeleteStackInstances(ctx workflow.Context, input *cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackInstancesAsync(ctx workflow.Context, input *cloudformation.DeleteStackInstancesInput) *DeleteStackInstancesFuture

	DeleteStackSet(ctx workflow.Context, input *cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error)
	DeleteStackSetAsync(ctx workflow.Context, input *cloudformation.DeleteStackSetInput) *DeleteStackSetFuture

	DeregisterType(ctx workflow.Context, input *cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error)
	DeregisterTypeAsync(ctx workflow.Context, input *cloudformation.DeregisterTypeInput) *DeregisterTypeFuture

	DescribeAccountLimits(ctx workflow.Context, input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *cloudformation.DescribeAccountLimitsInput) *DescribeAccountLimitsFuture

	DescribeChangeSet(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetAsync(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) *DescribeChangeSetFuture

	DescribeStackDriftDetectionStatus(ctx workflow.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackDriftDetectionStatusAsync(ctx workflow.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) *DescribeStackDriftDetectionStatusFuture

	DescribeStackEvents(ctx workflow.Context, input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackEventsAsync(ctx workflow.Context, input *cloudformation.DescribeStackEventsInput) *DescribeStackEventsFuture

	DescribeStackInstance(ctx workflow.Context, input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackInstanceAsync(ctx workflow.Context, input *cloudformation.DescribeStackInstanceInput) *DescribeStackInstanceFuture

	DescribeStackResource(ctx workflow.Context, input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourceInput) *DescribeStackResourceFuture

	DescribeStackResourceDrifts(ctx workflow.Context, input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResourceDriftsAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourceDriftsInput) *DescribeStackResourceDriftsFuture

	DescribeStackResources(ctx workflow.Context, input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackResourcesAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourcesInput) *DescribeStackResourcesFuture

	DescribeStackSet(ctx workflow.Context, input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetAsync(ctx workflow.Context, input *cloudformation.DescribeStackSetInput) *DescribeStackSetFuture

	DescribeStackSetOperation(ctx workflow.Context, input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStackSetOperationAsync(ctx workflow.Context, input *cloudformation.DescribeStackSetOperationInput) *DescribeStackSetOperationFuture

	DescribeStacks(ctx workflow.Context, input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	DescribeStacksAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *DescribeStacksFuture

	DescribeType(ctx workflow.Context, input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeAsync(ctx workflow.Context, input *cloudformation.DescribeTypeInput) *DescribeTypeFuture

	DescribeTypeRegistration(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error)
	DescribeTypeRegistrationAsync(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) *DescribeTypeRegistrationFuture

	DetectStackDrift(ctx workflow.Context, input *cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error)
	DetectStackDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackDriftInput) *DetectStackDriftFuture

	DetectStackResourceDrift(ctx workflow.Context, input *cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error)
	DetectStackResourceDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackResourceDriftInput) *DetectStackResourceDriftFuture

	DetectStackSetDrift(ctx workflow.Context, input *cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error)
	DetectStackSetDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackSetDriftInput) *DetectStackSetDriftFuture

	EstimateTemplateCost(ctx workflow.Context, input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)
	EstimateTemplateCostAsync(ctx workflow.Context, input *cloudformation.EstimateTemplateCostInput) *EstimateTemplateCostFuture

	ExecuteChangeSet(ctx workflow.Context, input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)
	ExecuteChangeSetAsync(ctx workflow.Context, input *cloudformation.ExecuteChangeSetInput) *ExecuteChangeSetFuture

	GetStackPolicy(ctx workflow.Context, input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)
	GetStackPolicyAsync(ctx workflow.Context, input *cloudformation.GetStackPolicyInput) *GetStackPolicyFuture

	GetTemplate(ctx workflow.Context, input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)
	GetTemplateAsync(ctx workflow.Context, input *cloudformation.GetTemplateInput) *GetTemplateFuture

	GetTemplateSummary(ctx workflow.Context, input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)
	GetTemplateSummaryAsync(ctx workflow.Context, input *cloudformation.GetTemplateSummaryInput) *GetTemplateSummaryFuture

	ListChangeSets(ctx workflow.Context, input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error)
	ListChangeSetsAsync(ctx workflow.Context, input *cloudformation.ListChangeSetsInput) *ListChangeSetsFuture

	ListExports(ctx workflow.Context, input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error)
	ListExportsAsync(ctx workflow.Context, input *cloudformation.ListExportsInput) *ListExportsFuture

	ListImports(ctx workflow.Context, input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error)
	ListImportsAsync(ctx workflow.Context, input *cloudformation.ListImportsInput) *ListImportsFuture

	ListStackInstances(ctx workflow.Context, input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error)
	ListStackInstancesAsync(ctx workflow.Context, input *cloudformation.ListStackInstancesInput) *ListStackInstancesFuture

	ListStackResources(ctx workflow.Context, input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)
	ListStackResourcesAsync(ctx workflow.Context, input *cloudformation.ListStackResourcesInput) *ListStackResourcesFuture

	ListStackSetOperationResults(ctx workflow.Context, input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperationResultsAsync(ctx workflow.Context, input *cloudformation.ListStackSetOperationResultsInput) *ListStackSetOperationResultsFuture

	ListStackSetOperations(ctx workflow.Context, input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSetOperationsAsync(ctx workflow.Context, input *cloudformation.ListStackSetOperationsInput) *ListStackSetOperationsFuture

	ListStackSets(ctx workflow.Context, input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error)
	ListStackSetsAsync(ctx workflow.Context, input *cloudformation.ListStackSetsInput) *ListStackSetsFuture

	ListStacks(ctx workflow.Context, input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)
	ListStacksAsync(ctx workflow.Context, input *cloudformation.ListStacksInput) *ListStacksFuture

	ListTypeRegistrations(ctx workflow.Context, input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeRegistrationsAsync(ctx workflow.Context, input *cloudformation.ListTypeRegistrationsInput) *ListTypeRegistrationsFuture

	ListTypeVersions(ctx workflow.Context, input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypeVersionsAsync(ctx workflow.Context, input *cloudformation.ListTypeVersionsInput) *ListTypeVersionsFuture

	ListTypes(ctx workflow.Context, input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error)
	ListTypesAsync(ctx workflow.Context, input *cloudformation.ListTypesInput) *ListTypesFuture

	RecordHandlerProgress(ctx workflow.Context, input *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error)
	RecordHandlerProgressAsync(ctx workflow.Context, input *cloudformation.RecordHandlerProgressInput) *RecordHandlerProgressFuture

	RegisterType(ctx workflow.Context, input *cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error)
	RegisterTypeAsync(ctx workflow.Context, input *cloudformation.RegisterTypeInput) *RegisterTypeFuture

	SetStackPolicy(ctx workflow.Context, input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)
	SetStackPolicyAsync(ctx workflow.Context, input *cloudformation.SetStackPolicyInput) *SetStackPolicyFuture

	SetTypeDefaultVersion(ctx workflow.Context, input *cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error)
	SetTypeDefaultVersionAsync(ctx workflow.Context, input *cloudformation.SetTypeDefaultVersionInput) *SetTypeDefaultVersionFuture

	SignalResource(ctx workflow.Context, input *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)
	SignalResourceAsync(ctx workflow.Context, input *cloudformation.SignalResourceInput) *SignalResourceFuture

	StopStackSetOperation(ctx workflow.Context, input *cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error)
	StopStackSetOperationAsync(ctx workflow.Context, input *cloudformation.StopStackSetOperationInput) *StopStackSetOperationFuture

	UpdateStack(ctx workflow.Context, input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)
	UpdateStackAsync(ctx workflow.Context, input *cloudformation.UpdateStackInput) *UpdateStackFuture

	UpdateStackInstances(ctx workflow.Context, input *cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error)
	UpdateStackInstancesAsync(ctx workflow.Context, input *cloudformation.UpdateStackInstancesInput) *UpdateStackInstancesFuture

	UpdateStackSet(ctx workflow.Context, input *cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error)
	UpdateStackSetAsync(ctx workflow.Context, input *cloudformation.UpdateStackSetInput) *UpdateStackSetFuture

	UpdateTerminationProtection(ctx workflow.Context, input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error)
	UpdateTerminationProtectionAsync(ctx workflow.Context, input *cloudformation.UpdateTerminationProtectionInput) *UpdateTerminationProtectionFuture

	ValidateTemplate(ctx workflow.Context, input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
	ValidateTemplateAsync(ctx workflow.Context, input *cloudformation.ValidateTemplateInput) *ValidateTemplateFuture

	WaitUntilChangeSetCreateComplete(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) error
	WaitUntilChangeSetCreateCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) *clients.VoidFuture

	WaitUntilStackCreateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackCreateCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilStackDeleteComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackDeleteCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilStackExists(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackExistsAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilStackImportComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackImportCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilStackRollbackComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackRollbackCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilStackUpdateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackUpdateCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *clients.VoidFuture

	WaitUntilTypeRegistrationComplete(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) error
	WaitUntilTypeRegistrationCompleteAsync(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}