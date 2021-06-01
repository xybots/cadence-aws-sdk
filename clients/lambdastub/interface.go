// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package lambdastub

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddLayerVersionPermission(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error)
	AddLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) *AddLayerVersionPermissionFuture

	AddPermission(ctx workflow.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)
	AddPermissionAsync(ctx workflow.Context, input *lambda.AddPermissionInput) *AddPermissionFuture

	CreateAlias(ctx workflow.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)
	CreateAliasAsync(ctx workflow.Context, input *lambda.CreateAliasInput) *CreateAliasFuture

	CreateCodeSigningConfig(ctx workflow.Context, input *lambda.CreateCodeSigningConfigInput) (*lambda.CreateCodeSigningConfigOutput, error)
	CreateCodeSigningConfigAsync(ctx workflow.Context, input *lambda.CreateCodeSigningConfigInput) *CreateCodeSigningConfigFuture

	CreateEventSourceMapping(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingAsync(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) *CreateEventSourceMappingFuture

	CreateFunction(ctx workflow.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
	CreateFunctionAsync(ctx workflow.Context, input *lambda.CreateFunctionInput) *CreateFunctionFuture

	DeleteAlias(ctx workflow.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)
	DeleteAliasAsync(ctx workflow.Context, input *lambda.DeleteAliasInput) *DeleteAliasFuture

	DeleteCodeSigningConfig(ctx workflow.Context, input *lambda.DeleteCodeSigningConfigInput) (*lambda.DeleteCodeSigningConfigOutput, error)
	DeleteCodeSigningConfigAsync(ctx workflow.Context, input *lambda.DeleteCodeSigningConfigInput) *DeleteCodeSigningConfigFuture

	DeleteEventSourceMapping(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingAsync(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) *DeleteEventSourceMappingFuture

	DeleteFunction(ctx workflow.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionAsync(ctx workflow.Context, input *lambda.DeleteFunctionInput) *DeleteFunctionFuture

	DeleteFunctionCodeSigningConfig(ctx workflow.Context, input *lambda.DeleteFunctionCodeSigningConfigInput) (*lambda.DeleteFunctionCodeSigningConfigOutput, error)
	DeleteFunctionCodeSigningConfigAsync(ctx workflow.Context, input *lambda.DeleteFunctionCodeSigningConfigInput) *DeleteFunctionCodeSigningConfigFuture

	DeleteFunctionConcurrency(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) *DeleteFunctionConcurrencyFuture

	DeleteFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	DeleteFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) *DeleteFunctionEventInvokeConfigFuture

	DeleteLayerVersion(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error)
	DeleteLayerVersionAsync(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) *DeleteLayerVersionFuture

	DeleteProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	DeleteProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) *DeleteProvisionedConcurrencyConfigFuture

	GetAccountSettings(ctx workflow.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsAsync(ctx workflow.Context, input *lambda.GetAccountSettingsInput) *GetAccountSettingsFuture

	GetAlias(ctx workflow.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error)
	GetAliasAsync(ctx workflow.Context, input *lambda.GetAliasInput) *GetAliasFuture

	GetCodeSigningConfig(ctx workflow.Context, input *lambda.GetCodeSigningConfigInput) (*lambda.GetCodeSigningConfigOutput, error)
	GetCodeSigningConfigAsync(ctx workflow.Context, input *lambda.GetCodeSigningConfigInput) *GetCodeSigningConfigFuture

	GetEventSourceMapping(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingAsync(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) *GetEventSourceMappingFuture

	GetFunction(ctx workflow.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
	GetFunctionAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *GetFunctionFuture

	GetFunctionCodeSigningConfig(ctx workflow.Context, input *lambda.GetFunctionCodeSigningConfigInput) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionCodeSigningConfigAsync(ctx workflow.Context, input *lambda.GetFunctionCodeSigningConfigInput) *GetFunctionCodeSigningConfigFuture

	GetFunctionConcurrency(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) *GetFunctionConcurrencyFuture

	GetFunctionConfiguration(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *GetFunctionConfigurationFuture

	GetFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) *GetFunctionEventInvokeConfigFuture

	GetLayerVersion(ctx workflow.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionAsync(ctx workflow.Context, input *lambda.GetLayerVersionInput) *GetLayerVersionFuture

	GetLayerVersionByArn(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionByArnAsync(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) *GetLayerVersionByArnFuture

	GetLayerVersionPolicy(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error)
	GetLayerVersionPolicyAsync(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) *GetLayerVersionPolicyFuture

	GetPolicy(ctx workflow.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *lambda.GetPolicyInput) *GetPolicyFuture

	GetProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) *GetProvisionedConcurrencyConfigFuture

	Invoke(ctx workflow.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error)
	InvokeAsync(ctx workflow.Context, input *lambda.InvokeInput) *InvokeFuture

	ListAliases(ctx workflow.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)
	ListAliasesAsync(ctx workflow.Context, input *lambda.ListAliasesInput) *ListAliasesFuture

	ListCodeSigningConfigs(ctx workflow.Context, input *lambda.ListCodeSigningConfigsInput) (*lambda.ListCodeSigningConfigsOutput, error)
	ListCodeSigningConfigsAsync(ctx workflow.Context, input *lambda.ListCodeSigningConfigsInput) *ListCodeSigningConfigsFuture

	ListEventSourceMappings(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsAsync(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) *ListEventSourceMappingsFuture

	ListFunctionEventInvokeConfigs(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionEventInvokeConfigsAsync(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) *ListFunctionEventInvokeConfigsFuture

	ListFunctions(ctx workflow.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
	ListFunctionsAsync(ctx workflow.Context, input *lambda.ListFunctionsInput) *ListFunctionsFuture

	ListFunctionsByCodeSigningConfig(ctx workflow.Context, input *lambda.ListFunctionsByCodeSigningConfigInput) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListFunctionsByCodeSigningConfigAsync(ctx workflow.Context, input *lambda.ListFunctionsByCodeSigningConfigInput) *ListFunctionsByCodeSigningConfigFuture

	ListLayerVersions(ctx workflow.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error)
	ListLayerVersionsAsync(ctx workflow.Context, input *lambda.ListLayerVersionsInput) *ListLayerVersionsFuture

	ListLayers(ctx workflow.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error)
	ListLayersAsync(ctx workflow.Context, input *lambda.ListLayersInput) *ListLayersFuture

	ListProvisionedConcurrencyConfigs(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListProvisionedConcurrencyConfigsAsync(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) *ListProvisionedConcurrencyConfigsFuture

	ListTags(ctx workflow.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *lambda.ListTagsInput) *ListTagsFuture

	ListVersionsByFunction(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionAsync(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) *ListVersionsByFunctionFuture

	PublishLayerVersion(ctx workflow.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error)
	PublishLayerVersionAsync(ctx workflow.Context, input *lambda.PublishLayerVersionInput) *PublishLayerVersionFuture

	PublishVersion(ctx workflow.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)
	PublishVersionAsync(ctx workflow.Context, input *lambda.PublishVersionInput) *PublishVersionFuture

	PutFunctionCodeSigningConfig(ctx workflow.Context, input *lambda.PutFunctionCodeSigningConfigInput) (*lambda.PutFunctionCodeSigningConfigOutput, error)
	PutFunctionCodeSigningConfigAsync(ctx workflow.Context, input *lambda.PutFunctionCodeSigningConfigInput) *PutFunctionCodeSigningConfigFuture

	PutFunctionConcurrency(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) *PutFunctionConcurrencyFuture

	PutFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	PutFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) *PutFunctionEventInvokeConfigFuture

	PutProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	PutProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) *PutProvisionedConcurrencyConfigFuture

	RemoveLayerVersionPermission(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error)
	RemoveLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) *RemoveLayerVersionPermissionFuture

	RemovePermission(ctx workflow.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *lambda.RemovePermissionInput) *RemovePermissionFuture

	TagResource(ctx workflow.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *lambda.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *lambda.UntagResourceInput) *UntagResourceFuture

	UpdateAlias(ctx workflow.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)
	UpdateAliasAsync(ctx workflow.Context, input *lambda.UpdateAliasInput) *UpdateAliasFuture

	UpdateCodeSigningConfig(ctx workflow.Context, input *lambda.UpdateCodeSigningConfigInput) (*lambda.UpdateCodeSigningConfigOutput, error)
	UpdateCodeSigningConfigAsync(ctx workflow.Context, input *lambda.UpdateCodeSigningConfigInput) *UpdateCodeSigningConfigFuture

	UpdateEventSourceMapping(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingAsync(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) *UpdateEventSourceMappingFuture

	UpdateFunctionCode(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeAsync(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) *UpdateFunctionCodeFuture

	UpdateFunctionConfiguration(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationAsync(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) *UpdateFunctionConfigurationFuture

	UpdateFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
	UpdateFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) *UpdateFunctionEventInvokeConfigFuture

	WaitUntilFunctionActive(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionActiveAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *clients.VoidFuture

	WaitUntilFunctionExists(ctx workflow.Context, input *lambda.GetFunctionInput) error
	WaitUntilFunctionExistsAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *clients.VoidFuture

	WaitUntilFunctionUpdated(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionUpdatedAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}