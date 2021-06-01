// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package pistub

import (
	"github.com/aws/aws-sdk-go/service/pi"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeDimensionKeys(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error)
	DescribeDimensionKeysAsync(ctx workflow.Context, input *pi.DescribeDimensionKeysInput) *DescribeDimensionKeysFuture

	GetResourceMetrics(ctx workflow.Context, input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error)
	GetResourceMetricsAsync(ctx workflow.Context, input *pi.GetResourceMetricsInput) *GetResourceMetricsFuture
}

func NewClient() Client {
	return &stub{}
}
