// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package serverlessapplicationrepositorystub

import (
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) *CreateApplicationFuture

	CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
	CreateApplicationVersionAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) *CreateApplicationVersionFuture

	CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
	CreateCloudFormationChangeSetAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) *CreateCloudFormationChangeSetFuture

	CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
	CreateCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) *CreateCloudFormationTemplateFuture

	DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) *DeleteApplicationFuture

	GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error)
	GetApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) *GetApplicationFuture

	GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
	GetApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) *GetApplicationPolicyFuture

	GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
	GetCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) *GetCloudFormationTemplateFuture

	ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
	ListApplicationDependenciesAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) *ListApplicationDependenciesFuture

	ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
	ListApplicationVersionsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) *ListApplicationVersionsFuture

	ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) *ListApplicationsFuture

	PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
	PutApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) *PutApplicationPolicyFuture

	UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error)
	UnshareApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) *UnshareApplicationFuture

	UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) *UpdateApplicationFuture
}

func NewClient() Client {
	return &stub{}
}
