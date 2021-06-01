// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package sagemakerruntimestub

import (
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type InvokeEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InvokeEndpointFuture) Get(ctx workflow.Context) (*sagemakerruntime.InvokeEndpointOutput, error) {
	var output sagemakerruntime.InvokeEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) InvokeEndpoint(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
	var output sagemakerruntime.InvokeEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws-sagemakerruntime-InvokeEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InvokeEndpointAsync(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) *InvokeEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sagemakerruntime-InvokeEndpoint", input)
	return &InvokeEndpointFuture{Future: future}
}