// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package kinesisvideomediastub

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error)
	GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *GetMediaFuture
}

func NewClient() Client {
	return &stub{}
}
