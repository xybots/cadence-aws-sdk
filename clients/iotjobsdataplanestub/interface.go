// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iotjobsdataplanestub

import (
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeJobExecution(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error)
	DescribeJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) *DescribeJobExecutionFuture

	GetPendingJobExecutions(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error)
	GetPendingJobExecutionsAsync(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) *GetPendingJobExecutionsFuture

	StartNextPendingJobExecution(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error)
	StartNextPendingJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) *StartNextPendingJobExecutionFuture

	UpdateJobExecution(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error)
	UpdateJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) *UpdateJobExecutionFuture
}

func NewClient() Client {
	return &stub{}
}