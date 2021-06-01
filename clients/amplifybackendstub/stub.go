// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package amplifybackendstub

import (
	"github.com/aws/aws-sdk-go/service/amplifybackend"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CloneBackendFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloneBackendFuture) Get(ctx workflow.Context) (*amplifybackend.CloneBackendOutput, error) {
	var output amplifybackend.CloneBackendOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBackendFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBackendFuture) Get(ctx workflow.Context) (*amplifybackend.CreateBackendOutput, error) {
	var output amplifybackend.CreateBackendOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBackendAPIFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBackendAPIFuture) Get(ctx workflow.Context) (*amplifybackend.CreateBackendAPIOutput, error) {
	var output amplifybackend.CreateBackendAPIOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBackendAuthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBackendAuthFuture) Get(ctx workflow.Context) (*amplifybackend.CreateBackendAuthOutput, error) {
	var output amplifybackend.CreateBackendAuthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBackendConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBackendConfigFuture) Get(ctx workflow.Context) (*amplifybackend.CreateBackendConfigOutput, error) {
	var output amplifybackend.CreateBackendConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateTokenFuture) Get(ctx workflow.Context) (*amplifybackend.CreateTokenOutput, error) {
	var output amplifybackend.CreateTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBackendFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBackendFuture) Get(ctx workflow.Context) (*amplifybackend.DeleteBackendOutput, error) {
	var output amplifybackend.DeleteBackendOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBackendAPIFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBackendAPIFuture) Get(ctx workflow.Context) (*amplifybackend.DeleteBackendAPIOutput, error) {
	var output amplifybackend.DeleteBackendAPIOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBackendAuthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBackendAuthFuture) Get(ctx workflow.Context) (*amplifybackend.DeleteBackendAuthOutput, error) {
	var output amplifybackend.DeleteBackendAuthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTokenFuture) Get(ctx workflow.Context) (*amplifybackend.DeleteTokenOutput, error) {
	var output amplifybackend.DeleteTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GenerateBackendAPIModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GenerateBackendAPIModelsFuture) Get(ctx workflow.Context) (*amplifybackend.GenerateBackendAPIModelsOutput, error) {
	var output amplifybackend.GenerateBackendAPIModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendFuture) Get(ctx workflow.Context) (*amplifybackend.GetBackendOutput, error) {
	var output amplifybackend.GetBackendOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendAPIFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendAPIFuture) Get(ctx workflow.Context) (*amplifybackend.GetBackendAPIOutput, error) {
	var output amplifybackend.GetBackendAPIOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendAPIModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendAPIModelsFuture) Get(ctx workflow.Context) (*amplifybackend.GetBackendAPIModelsOutput, error) {
	var output amplifybackend.GetBackendAPIModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendAuthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendAuthFuture) Get(ctx workflow.Context) (*amplifybackend.GetBackendAuthOutput, error) {
	var output amplifybackend.GetBackendAuthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBackendJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBackendJobFuture) Get(ctx workflow.Context) (*amplifybackend.GetBackendJobOutput, error) {
	var output amplifybackend.GetBackendJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTokenFuture) Get(ctx workflow.Context) (*amplifybackend.GetTokenOutput, error) {
	var output amplifybackend.GetTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBackendJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBackendJobsFuture) Get(ctx workflow.Context) (*amplifybackend.ListBackendJobsOutput, error) {
	var output amplifybackend.ListBackendJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveAllBackendsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveAllBackendsFuture) Get(ctx workflow.Context) (*amplifybackend.RemoveAllBackendsOutput, error) {
	var output amplifybackend.RemoveAllBackendsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveBackendConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveBackendConfigFuture) Get(ctx workflow.Context) (*amplifybackend.RemoveBackendConfigOutput, error) {
	var output amplifybackend.RemoveBackendConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBackendAPIFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBackendAPIFuture) Get(ctx workflow.Context) (*amplifybackend.UpdateBackendAPIOutput, error) {
	var output amplifybackend.UpdateBackendAPIOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBackendAuthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBackendAuthFuture) Get(ctx workflow.Context) (*amplifybackend.UpdateBackendAuthOutput, error) {
	var output amplifybackend.UpdateBackendAuthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBackendConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBackendConfigFuture) Get(ctx workflow.Context) (*amplifybackend.UpdateBackendConfigOutput, error) {
	var output amplifybackend.UpdateBackendConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBackendJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBackendJobFuture) Get(ctx workflow.Context) (*amplifybackend.UpdateBackendJobOutput, error) {
	var output amplifybackend.UpdateBackendJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CloneBackend(ctx workflow.Context, input *amplifybackend.CloneBackendInput) (*amplifybackend.CloneBackendOutput, error) {
	var output amplifybackend.CloneBackendOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CloneBackend", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CloneBackendAsync(ctx workflow.Context, input *amplifybackend.CloneBackendInput) *CloneBackendFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CloneBackend", input)
	return &CloneBackendFuture{Future: future}
}

func (a *stub) CreateBackend(ctx workflow.Context, input *amplifybackend.CreateBackendInput) (*amplifybackend.CreateBackendOutput, error) {
	var output amplifybackend.CreateBackendOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackend", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackendAsync(ctx workflow.Context, input *amplifybackend.CreateBackendInput) *CreateBackendFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackend", input)
	return &CreateBackendFuture{Future: future}
}

func (a *stub) CreateBackendAPI(ctx workflow.Context, input *amplifybackend.CreateBackendAPIInput) (*amplifybackend.CreateBackendAPIOutput, error) {
	var output amplifybackend.CreateBackendAPIOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendAPI", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackendAPIAsync(ctx workflow.Context, input *amplifybackend.CreateBackendAPIInput) *CreateBackendAPIFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendAPI", input)
	return &CreateBackendAPIFuture{Future: future}
}

func (a *stub) CreateBackendAuth(ctx workflow.Context, input *amplifybackend.CreateBackendAuthInput) (*amplifybackend.CreateBackendAuthOutput, error) {
	var output amplifybackend.CreateBackendAuthOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendAuth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackendAuthAsync(ctx workflow.Context, input *amplifybackend.CreateBackendAuthInput) *CreateBackendAuthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendAuth", input)
	return &CreateBackendAuthFuture{Future: future}
}

func (a *stub) CreateBackendConfig(ctx workflow.Context, input *amplifybackend.CreateBackendConfigInput) (*amplifybackend.CreateBackendConfigOutput, error) {
	var output amplifybackend.CreateBackendConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackendConfigAsync(ctx workflow.Context, input *amplifybackend.CreateBackendConfigInput) *CreateBackendConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateBackendConfig", input)
	return &CreateBackendConfigFuture{Future: future}
}

func (a *stub) CreateToken(ctx workflow.Context, input *amplifybackend.CreateTokenInput) (*amplifybackend.CreateTokenOutput, error) {
	var output amplifybackend.CreateTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTokenAsync(ctx workflow.Context, input *amplifybackend.CreateTokenInput) *CreateTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-CreateToken", input)
	return &CreateTokenFuture{Future: future}
}

func (a *stub) DeleteBackend(ctx workflow.Context, input *amplifybackend.DeleteBackendInput) (*amplifybackend.DeleteBackendOutput, error) {
	var output amplifybackend.DeleteBackendOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackend", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackendAsync(ctx workflow.Context, input *amplifybackend.DeleteBackendInput) *DeleteBackendFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackend", input)
	return &DeleteBackendFuture{Future: future}
}

func (a *stub) DeleteBackendAPI(ctx workflow.Context, input *amplifybackend.DeleteBackendAPIInput) (*amplifybackend.DeleteBackendAPIOutput, error) {
	var output amplifybackend.DeleteBackendAPIOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackendAPI", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackendAPIAsync(ctx workflow.Context, input *amplifybackend.DeleteBackendAPIInput) *DeleteBackendAPIFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackendAPI", input)
	return &DeleteBackendAPIFuture{Future: future}
}

func (a *stub) DeleteBackendAuth(ctx workflow.Context, input *amplifybackend.DeleteBackendAuthInput) (*amplifybackend.DeleteBackendAuthOutput, error) {
	var output amplifybackend.DeleteBackendAuthOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackendAuth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackendAuthAsync(ctx workflow.Context, input *amplifybackend.DeleteBackendAuthInput) *DeleteBackendAuthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteBackendAuth", input)
	return &DeleteBackendAuthFuture{Future: future}
}

func (a *stub) DeleteToken(ctx workflow.Context, input *amplifybackend.DeleteTokenInput) (*amplifybackend.DeleteTokenOutput, error) {
	var output amplifybackend.DeleteTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTokenAsync(ctx workflow.Context, input *amplifybackend.DeleteTokenInput) *DeleteTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-DeleteToken", input)
	return &DeleteTokenFuture{Future: future}
}

func (a *stub) GenerateBackendAPIModels(ctx workflow.Context, input *amplifybackend.GenerateBackendAPIModelsInput) (*amplifybackend.GenerateBackendAPIModelsOutput, error) {
	var output amplifybackend.GenerateBackendAPIModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GenerateBackendAPIModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateBackendAPIModelsAsync(ctx workflow.Context, input *amplifybackend.GenerateBackendAPIModelsInput) *GenerateBackendAPIModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GenerateBackendAPIModels", input)
	return &GenerateBackendAPIModelsFuture{Future: future}
}

func (a *stub) GetBackend(ctx workflow.Context, input *amplifybackend.GetBackendInput) (*amplifybackend.GetBackendOutput, error) {
	var output amplifybackend.GetBackendOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackend", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendAsync(ctx workflow.Context, input *amplifybackend.GetBackendInput) *GetBackendFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackend", input)
	return &GetBackendFuture{Future: future}
}

func (a *stub) GetBackendAPI(ctx workflow.Context, input *amplifybackend.GetBackendAPIInput) (*amplifybackend.GetBackendAPIOutput, error) {
	var output amplifybackend.GetBackendAPIOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAPI", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendAPIAsync(ctx workflow.Context, input *amplifybackend.GetBackendAPIInput) *GetBackendAPIFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAPI", input)
	return &GetBackendAPIFuture{Future: future}
}

func (a *stub) GetBackendAPIModels(ctx workflow.Context, input *amplifybackend.GetBackendAPIModelsInput) (*amplifybackend.GetBackendAPIModelsOutput, error) {
	var output amplifybackend.GetBackendAPIModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAPIModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendAPIModelsAsync(ctx workflow.Context, input *amplifybackend.GetBackendAPIModelsInput) *GetBackendAPIModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAPIModels", input)
	return &GetBackendAPIModelsFuture{Future: future}
}

func (a *stub) GetBackendAuth(ctx workflow.Context, input *amplifybackend.GetBackendAuthInput) (*amplifybackend.GetBackendAuthOutput, error) {
	var output amplifybackend.GetBackendAuthOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAuth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendAuthAsync(ctx workflow.Context, input *amplifybackend.GetBackendAuthInput) *GetBackendAuthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendAuth", input)
	return &GetBackendAuthFuture{Future: future}
}

func (a *stub) GetBackendJob(ctx workflow.Context, input *amplifybackend.GetBackendJobInput) (*amplifybackend.GetBackendJobOutput, error) {
	var output amplifybackend.GetBackendJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBackendJobAsync(ctx workflow.Context, input *amplifybackend.GetBackendJobInput) *GetBackendJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetBackendJob", input)
	return &GetBackendJobFuture{Future: future}
}

func (a *stub) GetToken(ctx workflow.Context, input *amplifybackend.GetTokenInput) (*amplifybackend.GetTokenOutput, error) {
	var output amplifybackend.GetTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTokenAsync(ctx workflow.Context, input *amplifybackend.GetTokenInput) *GetTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-GetToken", input)
	return &GetTokenFuture{Future: future}
}

func (a *stub) ListBackendJobs(ctx workflow.Context, input *amplifybackend.ListBackendJobsInput) (*amplifybackend.ListBackendJobsOutput, error) {
	var output amplifybackend.ListBackendJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-ListBackendJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBackendJobsAsync(ctx workflow.Context, input *amplifybackend.ListBackendJobsInput) *ListBackendJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-ListBackendJobs", input)
	return &ListBackendJobsFuture{Future: future}
}

func (a *stub) RemoveAllBackends(ctx workflow.Context, input *amplifybackend.RemoveAllBackendsInput) (*amplifybackend.RemoveAllBackendsOutput, error) {
	var output amplifybackend.RemoveAllBackendsOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-RemoveAllBackends", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveAllBackendsAsync(ctx workflow.Context, input *amplifybackend.RemoveAllBackendsInput) *RemoveAllBackendsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-RemoveAllBackends", input)
	return &RemoveAllBackendsFuture{Future: future}
}

func (a *stub) RemoveBackendConfig(ctx workflow.Context, input *amplifybackend.RemoveBackendConfigInput) (*amplifybackend.RemoveBackendConfigOutput, error) {
	var output amplifybackend.RemoveBackendConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-RemoveBackendConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveBackendConfigAsync(ctx workflow.Context, input *amplifybackend.RemoveBackendConfigInput) *RemoveBackendConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-RemoveBackendConfig", input)
	return &RemoveBackendConfigFuture{Future: future}
}

func (a *stub) UpdateBackendAPI(ctx workflow.Context, input *amplifybackend.UpdateBackendAPIInput) (*amplifybackend.UpdateBackendAPIOutput, error) {
	var output amplifybackend.UpdateBackendAPIOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendAPI", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBackendAPIAsync(ctx workflow.Context, input *amplifybackend.UpdateBackendAPIInput) *UpdateBackendAPIFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendAPI", input)
	return &UpdateBackendAPIFuture{Future: future}
}

func (a *stub) UpdateBackendAuth(ctx workflow.Context, input *amplifybackend.UpdateBackendAuthInput) (*amplifybackend.UpdateBackendAuthOutput, error) {
	var output amplifybackend.UpdateBackendAuthOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendAuth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBackendAuthAsync(ctx workflow.Context, input *amplifybackend.UpdateBackendAuthInput) *UpdateBackendAuthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendAuth", input)
	return &UpdateBackendAuthFuture{Future: future}
}

func (a *stub) UpdateBackendConfig(ctx workflow.Context, input *amplifybackend.UpdateBackendConfigInput) (*amplifybackend.UpdateBackendConfigOutput, error) {
	var output amplifybackend.UpdateBackendConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBackendConfigAsync(ctx workflow.Context, input *amplifybackend.UpdateBackendConfigInput) *UpdateBackendConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendConfig", input)
	return &UpdateBackendConfigFuture{Future: future}
}

func (a *stub) UpdateBackendJob(ctx workflow.Context, input *amplifybackend.UpdateBackendJobInput) (*amplifybackend.UpdateBackendJobOutput, error) {
	var output amplifybackend.UpdateBackendJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBackendJobAsync(ctx workflow.Context, input *amplifybackend.UpdateBackendJobInput) *UpdateBackendJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-amplifybackend-UpdateBackendJob", input)
	return &UpdateBackendJobFuture{Future: future}
}
