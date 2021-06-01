// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package inspectorstub

import (
	"github.com/aws/aws-sdk-go/service/inspector"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddAttributesToFindings(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error)
	AddAttributesToFindingsAsync(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) *AddAttributesToFindingsFuture

	CreateAssessmentTarget(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error)
	CreateAssessmentTargetAsync(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) *CreateAssessmentTargetFuture

	CreateAssessmentTemplate(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error)
	CreateAssessmentTemplateAsync(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) *CreateAssessmentTemplateFuture

	CreateExclusionsPreview(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error)
	CreateExclusionsPreviewAsync(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) *CreateExclusionsPreviewFuture

	CreateResourceGroup(ctx workflow.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error)
	CreateResourceGroupAsync(ctx workflow.Context, input *inspector.CreateResourceGroupInput) *CreateResourceGroupFuture

	DeleteAssessmentRun(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error)
	DeleteAssessmentRunAsync(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) *DeleteAssessmentRunFuture

	DeleteAssessmentTarget(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error)
	DeleteAssessmentTargetAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) *DeleteAssessmentTargetFuture

	DeleteAssessmentTemplate(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error)
	DeleteAssessmentTemplateAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) *DeleteAssessmentTemplateFuture

	DescribeAssessmentRuns(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error)
	DescribeAssessmentRunsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) *DescribeAssessmentRunsFuture

	DescribeAssessmentTargets(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error)
	DescribeAssessmentTargetsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) *DescribeAssessmentTargetsFuture

	DescribeAssessmentTemplates(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error)
	DescribeAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) *DescribeAssessmentTemplatesFuture

	DescribeCrossAccountAccessRole(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
	DescribeCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) *DescribeCrossAccountAccessRoleFuture

	DescribeExclusions(ctx workflow.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error)
	DescribeExclusionsAsync(ctx workflow.Context, input *inspector.DescribeExclusionsInput) *DescribeExclusionsFuture

	DescribeFindings(ctx workflow.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error)
	DescribeFindingsAsync(ctx workflow.Context, input *inspector.DescribeFindingsInput) *DescribeFindingsFuture

	DescribeResourceGroups(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error)
	DescribeResourceGroupsAsync(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) *DescribeResourceGroupsFuture

	DescribeRulesPackages(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error)
	DescribeRulesPackagesAsync(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) *DescribeRulesPackagesFuture

	GetAssessmentReport(ctx workflow.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error)
	GetAssessmentReportAsync(ctx workflow.Context, input *inspector.GetAssessmentReportInput) *GetAssessmentReportFuture

	GetExclusionsPreview(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error)
	GetExclusionsPreviewAsync(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) *GetExclusionsPreviewFuture

	GetTelemetryMetadata(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error)
	GetTelemetryMetadataAsync(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) *GetTelemetryMetadataFuture

	ListAssessmentRunAgents(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error)
	ListAssessmentRunAgentsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) *ListAssessmentRunAgentsFuture

	ListAssessmentRuns(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error)
	ListAssessmentRunsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) *ListAssessmentRunsFuture

	ListAssessmentTargets(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error)
	ListAssessmentTargetsAsync(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) *ListAssessmentTargetsFuture

	ListAssessmentTemplates(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error)
	ListAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) *ListAssessmentTemplatesFuture

	ListEventSubscriptions(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error)
	ListEventSubscriptionsAsync(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) *ListEventSubscriptionsFuture

	ListExclusions(ctx workflow.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error)
	ListExclusionsAsync(ctx workflow.Context, input *inspector.ListExclusionsInput) *ListExclusionsFuture

	ListFindings(ctx workflow.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error)
	ListFindingsAsync(ctx workflow.Context, input *inspector.ListFindingsInput) *ListFindingsFuture

	ListRulesPackages(ctx workflow.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error)
	ListRulesPackagesAsync(ctx workflow.Context, input *inspector.ListRulesPackagesInput) *ListRulesPackagesFuture

	ListTagsForResource(ctx workflow.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *inspector.ListTagsForResourceInput) *ListTagsForResourceFuture

	PreviewAgents(ctx workflow.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error)
	PreviewAgentsAsync(ctx workflow.Context, input *inspector.PreviewAgentsInput) *PreviewAgentsFuture

	RegisterCrossAccountAccessRole(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error)
	RegisterCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) *RegisterCrossAccountAccessRoleFuture

	RemoveAttributesFromFindings(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error)
	RemoveAttributesFromFindingsAsync(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) *RemoveAttributesFromFindingsFuture

	SetTagsForResource(ctx workflow.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error)
	SetTagsForResourceAsync(ctx workflow.Context, input *inspector.SetTagsForResourceInput) *SetTagsForResourceFuture

	StartAssessmentRun(ctx workflow.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error)
	StartAssessmentRunAsync(ctx workflow.Context, input *inspector.StartAssessmentRunInput) *StartAssessmentRunFuture

	StopAssessmentRun(ctx workflow.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error)
	StopAssessmentRunAsync(ctx workflow.Context, input *inspector.StopAssessmentRunInput) *StopAssessmentRunFuture

	SubscribeToEvent(ctx workflow.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error)
	SubscribeToEventAsync(ctx workflow.Context, input *inspector.SubscribeToEventInput) *SubscribeToEventFuture

	UnsubscribeFromEvent(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error)
	UnsubscribeFromEventAsync(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) *UnsubscribeFromEventFuture

	UpdateAssessmentTarget(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error)
	UpdateAssessmentTargetAsync(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) *UpdateAssessmentTargetFuture
}

func NewClient() Client {
	return &stub{}
}
