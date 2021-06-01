// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package autoscalingplansstub

import (
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateScalingPlan(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error)
	CreateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) *CreateScalingPlanFuture

	DeleteScalingPlan(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error)
	DeleteScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) *DeleteScalingPlanFuture

	DescribeScalingPlanResources(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)
	DescribeScalingPlanResourcesAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) *DescribeScalingPlanResourcesFuture

	DescribeScalingPlans(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error)
	DescribeScalingPlansAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) *DescribeScalingPlansFuture

	GetScalingPlanResourceForecastData(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)
	GetScalingPlanResourceForecastDataAsync(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) *GetScalingPlanResourceForecastDataFuture

	UpdateScalingPlan(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error)
	UpdateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) *UpdateScalingPlanFuture
}

func NewClient() Client {
	return &stub{}
}