// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package augmentedairuntimestub

import (
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error)
	DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *DeleteHumanLoopFuture

	DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error)
	DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *DescribeHumanLoopFuture

	ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error)
	ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *ListHumanLoopsFuture

	StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error)
	StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *StartHumanLoopFuture

	StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error)
	StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *StopHumanLoopFuture
}

func NewClient() Client {
	return &stub{}
}