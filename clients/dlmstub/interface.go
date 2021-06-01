// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package dlmstub

import (
	"github.com/aws/aws-sdk-go/service/dlm"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error)
	CreateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) *CreateLifecyclePolicyFuture

	DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) *DeleteLifecyclePolicyFuture

	GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error)
	GetLifecyclePoliciesAsync(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) *GetLifecyclePoliciesFuture

	GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) *GetLifecyclePolicyFuture

	ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *dlm.ListTagsForResourceInput) *ListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dlm.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dlm.UntagResourceInput) *UntagResourceFuture

	UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error)
	UpdateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) *UpdateLifecyclePolicyFuture
}

func NewClient() Client {
	return &stub{}
}
