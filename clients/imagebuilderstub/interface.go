// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package imagebuilderstub

import (
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelImageCreation(ctx workflow.Context, input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error)
	CancelImageCreationAsync(ctx workflow.Context, input *imagebuilder.CancelImageCreationInput) *CancelImageCreationFuture

	CreateComponent(ctx workflow.Context, input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error)
	CreateComponentAsync(ctx workflow.Context, input *imagebuilder.CreateComponentInput) *CreateComponentFuture

	CreateContainerRecipe(ctx workflow.Context, input *imagebuilder.CreateContainerRecipeInput) (*imagebuilder.CreateContainerRecipeOutput, error)
	CreateContainerRecipeAsync(ctx workflow.Context, input *imagebuilder.CreateContainerRecipeInput) *CreateContainerRecipeFuture

	CreateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error)
	CreateDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.CreateDistributionConfigurationInput) *CreateDistributionConfigurationFuture

	CreateImage(ctx workflow.Context, input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error)
	CreateImageAsync(ctx workflow.Context, input *imagebuilder.CreateImageInput) *CreateImageFuture

	CreateImagePipeline(ctx workflow.Context, input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error)
	CreateImagePipelineAsync(ctx workflow.Context, input *imagebuilder.CreateImagePipelineInput) *CreateImagePipelineFuture

	CreateImageRecipe(ctx workflow.Context, input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error)
	CreateImageRecipeAsync(ctx workflow.Context, input *imagebuilder.CreateImageRecipeInput) *CreateImageRecipeFuture

	CreateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error)
	CreateInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) *CreateInfrastructureConfigurationFuture

	DeleteComponent(ctx workflow.Context, input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error)
	DeleteComponentAsync(ctx workflow.Context, input *imagebuilder.DeleteComponentInput) *DeleteComponentFuture

	DeleteContainerRecipe(ctx workflow.Context, input *imagebuilder.DeleteContainerRecipeInput) (*imagebuilder.DeleteContainerRecipeOutput, error)
	DeleteContainerRecipeAsync(ctx workflow.Context, input *imagebuilder.DeleteContainerRecipeInput) *DeleteContainerRecipeFuture

	DeleteDistributionConfiguration(ctx workflow.Context, input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error)
	DeleteDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.DeleteDistributionConfigurationInput) *DeleteDistributionConfigurationFuture

	DeleteImage(ctx workflow.Context, input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error)
	DeleteImageAsync(ctx workflow.Context, input *imagebuilder.DeleteImageInput) *DeleteImageFuture

	DeleteImagePipeline(ctx workflow.Context, input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error)
	DeleteImagePipelineAsync(ctx workflow.Context, input *imagebuilder.DeleteImagePipelineInput) *DeleteImagePipelineFuture

	DeleteImageRecipe(ctx workflow.Context, input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error)
	DeleteImageRecipeAsync(ctx workflow.Context, input *imagebuilder.DeleteImageRecipeInput) *DeleteImageRecipeFuture

	DeleteInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error)
	DeleteInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) *DeleteInfrastructureConfigurationFuture

	GetComponent(ctx workflow.Context, input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error)
	GetComponentAsync(ctx workflow.Context, input *imagebuilder.GetComponentInput) *GetComponentFuture

	GetComponentPolicy(ctx workflow.Context, input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error)
	GetComponentPolicyAsync(ctx workflow.Context, input *imagebuilder.GetComponentPolicyInput) *GetComponentPolicyFuture

	GetContainerRecipe(ctx workflow.Context, input *imagebuilder.GetContainerRecipeInput) (*imagebuilder.GetContainerRecipeOutput, error)
	GetContainerRecipeAsync(ctx workflow.Context, input *imagebuilder.GetContainerRecipeInput) *GetContainerRecipeFuture

	GetContainerRecipePolicy(ctx workflow.Context, input *imagebuilder.GetContainerRecipePolicyInput) (*imagebuilder.GetContainerRecipePolicyOutput, error)
	GetContainerRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.GetContainerRecipePolicyInput) *GetContainerRecipePolicyFuture

	GetDistributionConfiguration(ctx workflow.Context, input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error)
	GetDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.GetDistributionConfigurationInput) *GetDistributionConfigurationFuture

	GetImage(ctx workflow.Context, input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error)
	GetImageAsync(ctx workflow.Context, input *imagebuilder.GetImageInput) *GetImageFuture

	GetImagePipeline(ctx workflow.Context, input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error)
	GetImagePipelineAsync(ctx workflow.Context, input *imagebuilder.GetImagePipelineInput) *GetImagePipelineFuture

	GetImagePolicy(ctx workflow.Context, input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error)
	GetImagePolicyAsync(ctx workflow.Context, input *imagebuilder.GetImagePolicyInput) *GetImagePolicyFuture

	GetImageRecipe(ctx workflow.Context, input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error)
	GetImageRecipeAsync(ctx workflow.Context, input *imagebuilder.GetImageRecipeInput) *GetImageRecipeFuture

	GetImageRecipePolicy(ctx workflow.Context, input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error)
	GetImageRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.GetImageRecipePolicyInput) *GetImageRecipePolicyFuture

	GetInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error)
	GetInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.GetInfrastructureConfigurationInput) *GetInfrastructureConfigurationFuture

	ImportComponent(ctx workflow.Context, input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error)
	ImportComponentAsync(ctx workflow.Context, input *imagebuilder.ImportComponentInput) *ImportComponentFuture

	ListComponentBuildVersions(ctx workflow.Context, input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error)
	ListComponentBuildVersionsAsync(ctx workflow.Context, input *imagebuilder.ListComponentBuildVersionsInput) *ListComponentBuildVersionsFuture

	ListComponents(ctx workflow.Context, input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error)
	ListComponentsAsync(ctx workflow.Context, input *imagebuilder.ListComponentsInput) *ListComponentsFuture

	ListContainerRecipes(ctx workflow.Context, input *imagebuilder.ListContainerRecipesInput) (*imagebuilder.ListContainerRecipesOutput, error)
	ListContainerRecipesAsync(ctx workflow.Context, input *imagebuilder.ListContainerRecipesInput) *ListContainerRecipesFuture

	ListDistributionConfigurations(ctx workflow.Context, input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error)
	ListDistributionConfigurationsAsync(ctx workflow.Context, input *imagebuilder.ListDistributionConfigurationsInput) *ListDistributionConfigurationsFuture

	ListImageBuildVersions(ctx workflow.Context, input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error)
	ListImageBuildVersionsAsync(ctx workflow.Context, input *imagebuilder.ListImageBuildVersionsInput) *ListImageBuildVersionsFuture

	ListImagePipelineImages(ctx workflow.Context, input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error)
	ListImagePipelineImagesAsync(ctx workflow.Context, input *imagebuilder.ListImagePipelineImagesInput) *ListImagePipelineImagesFuture

	ListImagePipelines(ctx workflow.Context, input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error)
	ListImagePipelinesAsync(ctx workflow.Context, input *imagebuilder.ListImagePipelinesInput) *ListImagePipelinesFuture

	ListImageRecipes(ctx workflow.Context, input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error)
	ListImageRecipesAsync(ctx workflow.Context, input *imagebuilder.ListImageRecipesInput) *ListImageRecipesFuture

	ListImages(ctx workflow.Context, input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error)
	ListImagesAsync(ctx workflow.Context, input *imagebuilder.ListImagesInput) *ListImagesFuture

	ListInfrastructureConfigurations(ctx workflow.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error)
	ListInfrastructureConfigurationsAsync(ctx workflow.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) *ListInfrastructureConfigurationsFuture

	ListTagsForResource(ctx workflow.Context, input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *imagebuilder.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutComponentPolicy(ctx workflow.Context, input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error)
	PutComponentPolicyAsync(ctx workflow.Context, input *imagebuilder.PutComponentPolicyInput) *PutComponentPolicyFuture

	PutContainerRecipePolicy(ctx workflow.Context, input *imagebuilder.PutContainerRecipePolicyInput) (*imagebuilder.PutContainerRecipePolicyOutput, error)
	PutContainerRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.PutContainerRecipePolicyInput) *PutContainerRecipePolicyFuture

	PutImagePolicy(ctx workflow.Context, input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error)
	PutImagePolicyAsync(ctx workflow.Context, input *imagebuilder.PutImagePolicyInput) *PutImagePolicyFuture

	PutImageRecipePolicy(ctx workflow.Context, input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error)
	PutImageRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.PutImageRecipePolicyInput) *PutImageRecipePolicyFuture

	StartImagePipelineExecution(ctx workflow.Context, input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error)
	StartImagePipelineExecutionAsync(ctx workflow.Context, input *imagebuilder.StartImagePipelineExecutionInput) *StartImagePipelineExecutionFuture

	TagResource(ctx workflow.Context, input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *imagebuilder.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *imagebuilder.UntagResourceInput) *UntagResourceFuture

	UpdateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error)
	UpdateDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.UpdateDistributionConfigurationInput) *UpdateDistributionConfigurationFuture

	UpdateImagePipeline(ctx workflow.Context, input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error)
	UpdateImagePipelineAsync(ctx workflow.Context, input *imagebuilder.UpdateImagePipelineInput) *UpdateImagePipelineFuture

	UpdateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error)
	UpdateInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) *UpdateInfrastructureConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}
