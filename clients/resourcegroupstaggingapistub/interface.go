// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package resourcegroupstaggingapistub

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error)
	DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *DescribeReportCreationFuture

	GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error)
	GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *GetComplianceSummaryFuture

	GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *GetResourcesFuture

	GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *GetTagKeysFuture

	GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *GetTagValuesFuture

	StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error)
	StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *StartReportCreationFuture

	TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
	TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *TagResourcesFuture

	UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
	UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *UntagResourcesFuture
}

func NewClient() Client {
	return &stub{}
}
