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

type stub struct{}

type CreateLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.CreateLifecyclePolicyOutput, error) {
	var output dlm.CreateLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.DeleteLifecyclePolicyOutput, error) {
	var output dlm.DeleteLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLifecyclePoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLifecyclePoliciesFuture) Get(ctx workflow.Context) (*dlm.GetLifecyclePoliciesOutput, error) {
	var output dlm.GetLifecyclePoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.GetLifecyclePolicyOutput, error) {
	var output dlm.GetLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*dlm.ListTagsForResourceOutput, error) {
	var output dlm.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*dlm.TagResourceOutput, error) {
	var output dlm.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*dlm.UntagResourceOutput, error) {
	var output dlm.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.UpdateLifecyclePolicyOutput, error) {
	var output dlm.UpdateLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
	var output dlm.CreateLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-CreateLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) *CreateLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-CreateLifecyclePolicy", input)
	return &CreateLifecyclePolicyFuture{Future: future}
}

func (a *stub) DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
	var output dlm.DeleteLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-DeleteLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLifecyclePolicyAsync(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) *DeleteLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-DeleteLifecyclePolicy", input)
	return &DeleteLifecyclePolicyFuture{Future: future}
}

func (a *stub) GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
	var output dlm.GetLifecyclePoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-GetLifecyclePolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePoliciesAsync(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) *GetLifecyclePoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-GetLifecyclePolicies", input)
	return &GetLifecyclePoliciesFuture{Future: future}
}

func (a *stub) GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
	var output dlm.GetLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-GetLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePolicyAsync(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) *GetLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-GetLifecyclePolicy", input)
	return &GetLifecyclePolicyFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
	var output dlm.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *dlm.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
	var output dlm.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dlm.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
	var output dlm.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dlm.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
	var output dlm.UpdateLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-dlm-UpdateLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) *UpdateLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dlm-UpdateLifecyclePolicy", input)
	return &UpdateLifecyclePolicyFuture{Future: future}
}
