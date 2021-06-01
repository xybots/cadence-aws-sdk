// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package timestreamquerystub

import (
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelQuery(ctx workflow.Context, input *timestreamquery.CancelQueryInput) (*timestreamquery.CancelQueryOutput, error)
	CancelQueryAsync(ctx workflow.Context, input *timestreamquery.CancelQueryInput) *CancelQueryFuture

	DescribeEndpoints(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) (*timestreamquery.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) *DescribeEndpointsFuture

	Query(ctx workflow.Context, input *timestreamquery.QueryInput) (*timestreamquery.QueryOutput, error)
	QueryAsync(ctx workflow.Context, input *timestreamquery.QueryInput) *QueryFuture
}

func NewClient() Client {
	return &stub{}
}
