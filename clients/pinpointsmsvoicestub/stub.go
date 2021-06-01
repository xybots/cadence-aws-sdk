// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package pinpointsmsvoicestub

import (
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateConfigurationSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteConfigurationSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetConfigurationSetEventDestinationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetConfigurationSetEventDestinationsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListConfigurationSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListConfigurationSetsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendVoiceMessageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendVoiceMessageFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-CreateConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *CreateConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-CreateConfigurationSet", input)
	return &CreateConfigurationSetFuture{Future: future}
}

func (a *stub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-CreateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *CreateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-CreateConfigurationSetEventDestination", input)
	return &CreateConfigurationSetEventDestinationFuture{Future: future}
}

func (a *stub) DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-DeleteConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *DeleteConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-DeleteConfigurationSet", input)
	return &DeleteConfigurationSetFuture{Future: future}
}

func (a *stub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-DeleteConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *DeleteConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-DeleteConfigurationSetEventDestination", input)
	return &DeleteConfigurationSetEventDestinationFuture{Future: future}
}

func (a *stub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-GetConfigurationSetEventDestinations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *GetConfigurationSetEventDestinationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-GetConfigurationSetEventDestinations", input)
	return &GetConfigurationSetEventDestinationsFuture{Future: future}
}

func (a *stub) ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-ListConfigurationSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *ListConfigurationSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-ListConfigurationSets", input)
	return &ListConfigurationSetsFuture{Future: future}
}

func (a *stub) SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-SendVoiceMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *SendVoiceMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-SendVoiceMessage", input)
	return &SendVoiceMessageFuture{Future: future}
}

func (a *stub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-UpdateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *UpdateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-pinpointsmsvoice-UpdateConfigurationSetEventDestination", input)
	return &UpdateConfigurationSetEventDestinationFuture{Future: future}
}