// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client connectiface.ConnectAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := connect.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (connectiface.ConnectAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return connect.New(sess), nil
}

func (a *Activities) AssociateApprovedOrigin(ctx context.Context, input *connect.AssociateApprovedOriginInput) (*connect.AssociateApprovedOriginOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateApprovedOriginWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateInstanceStorageConfig(ctx context.Context, input *connect.AssociateInstanceStorageConfigInput) (*connect.AssociateInstanceStorageConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateInstanceStorageConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateLambdaFunction(ctx context.Context, input *connect.AssociateLambdaFunctionInput) (*connect.AssociateLambdaFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateLambdaFunctionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateLexBot(ctx context.Context, input *connect.AssociateLexBotInput) (*connect.AssociateLexBotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateLexBotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateQueueQuickConnects(ctx context.Context, input *connect.AssociateQueueQuickConnectsInput) (*connect.AssociateQueueQuickConnectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateQueueQuickConnectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateRoutingProfileQueues(ctx context.Context, input *connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateSecurityKey(ctx context.Context, input *connect.AssociateSecurityKeyInput) (*connect.AssociateSecurityKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateSecurityKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateContactFlow(ctx context.Context, input *connect.CreateContactFlowInput) (*connect.CreateContactFlowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateContactFlowWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateInstance(ctx context.Context, input *connect.CreateInstanceInput) (*connect.CreateInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateIntegrationAssociation(ctx context.Context, input *connect.CreateIntegrationAssociationInput) (*connect.CreateIntegrationAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateIntegrationAssociationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateQueue(ctx context.Context, input *connect.CreateQueueInput) (*connect.CreateQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateQuickConnect(ctx context.Context, input *connect.CreateQuickConnectInput) (*connect.CreateQuickConnectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateQuickConnectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRoutingProfile(ctx context.Context, input *connect.CreateRoutingProfileInput) (*connect.CreateRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUseCase(ctx context.Context, input *connect.CreateUseCaseInput) (*connect.CreateUseCaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUseCaseWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUser(ctx context.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUserHierarchyGroup(ctx context.Context, input *connect.CreateUserHierarchyGroupInput) (*connect.CreateUserHierarchyGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUserHierarchyGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInstance(ctx context.Context, input *connect.DeleteInstanceInput) (*connect.DeleteInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteIntegrationAssociation(ctx context.Context, input *connect.DeleteIntegrationAssociationInput) (*connect.DeleteIntegrationAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteIntegrationAssociationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteQuickConnect(ctx context.Context, input *connect.DeleteQuickConnectInput) (*connect.DeleteQuickConnectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteQuickConnectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUseCase(ctx context.Context, input *connect.DeleteUseCaseInput) (*connect.DeleteUseCaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUseCaseWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUser(ctx context.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUserHierarchyGroup(ctx context.Context, input *connect.DeleteUserHierarchyGroupInput) (*connect.DeleteUserHierarchyGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUserHierarchyGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContactFlow(ctx context.Context, input *connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContactFlowWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeHoursOfOperation(ctx context.Context, input *connect.DescribeHoursOfOperationInput) (*connect.DescribeHoursOfOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeHoursOfOperationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstance(ctx context.Context, input *connect.DescribeInstanceInput) (*connect.DescribeInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstanceAttribute(ctx context.Context, input *connect.DescribeInstanceAttributeInput) (*connect.DescribeInstanceAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstanceAttributeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstanceStorageConfig(ctx context.Context, input *connect.DescribeInstanceStorageConfigInput) (*connect.DescribeInstanceStorageConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstanceStorageConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeQueue(ctx context.Context, input *connect.DescribeQueueInput) (*connect.DescribeQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeQuickConnect(ctx context.Context, input *connect.DescribeQuickConnectInput) (*connect.DescribeQuickConnectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeQuickConnectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRoutingProfile(ctx context.Context, input *connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUser(ctx context.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUserHierarchyGroup(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserHierarchyGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUserHierarchyStructure(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserHierarchyStructureWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateApprovedOrigin(ctx context.Context, input *connect.DisassociateApprovedOriginInput) (*connect.DisassociateApprovedOriginOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateApprovedOriginWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateInstanceStorageConfig(ctx context.Context, input *connect.DisassociateInstanceStorageConfigInput) (*connect.DisassociateInstanceStorageConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateInstanceStorageConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateLambdaFunction(ctx context.Context, input *connect.DisassociateLambdaFunctionInput) (*connect.DisassociateLambdaFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateLambdaFunctionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateLexBot(ctx context.Context, input *connect.DisassociateLexBotInput) (*connect.DisassociateLexBotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateLexBotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateQueueQuickConnects(ctx context.Context, input *connect.DisassociateQueueQuickConnectsInput) (*connect.DisassociateQueueQuickConnectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateQueueQuickConnectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateRoutingProfileQueues(ctx context.Context, input *connect.DisassociateRoutingProfileQueuesInput) (*connect.DisassociateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateSecurityKey(ctx context.Context, input *connect.DisassociateSecurityKeyInput) (*connect.DisassociateSecurityKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateSecurityKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetContactAttributes(ctx context.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetContactAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCurrentMetricData(ctx context.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCurrentMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFederationToken(ctx context.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFederationTokenWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMetricData(ctx context.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListApprovedOrigins(ctx context.Context, input *connect.ListApprovedOriginsInput) (*connect.ListApprovedOriginsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListApprovedOriginsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListContactFlows(ctx context.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListContactFlowsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHoursOfOperations(ctx context.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHoursOfOperationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInstanceAttributes(ctx context.Context, input *connect.ListInstanceAttributesInput) (*connect.ListInstanceAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInstanceAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInstanceStorageConfigs(ctx context.Context, input *connect.ListInstanceStorageConfigsInput) (*connect.ListInstanceStorageConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInstanceStorageConfigsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInstances(ctx context.Context, input *connect.ListInstancesInput) (*connect.ListInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListIntegrationAssociations(ctx context.Context, input *connect.ListIntegrationAssociationsInput) (*connect.ListIntegrationAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListIntegrationAssociationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLambdaFunctions(ctx context.Context, input *connect.ListLambdaFunctionsInput) (*connect.ListLambdaFunctionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLambdaFunctionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLexBots(ctx context.Context, input *connect.ListLexBotsInput) (*connect.ListLexBotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLexBotsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPhoneNumbers(ctx context.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPhoneNumbersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPrompts(ctx context.Context, input *connect.ListPromptsInput) (*connect.ListPromptsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPromptsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQueueQuickConnects(ctx context.Context, input *connect.ListQueueQuickConnectsInput) (*connect.ListQueueQuickConnectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQueueQuickConnectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQueues(ctx context.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQuickConnects(ctx context.Context, input *connect.ListQuickConnectsInput) (*connect.ListQuickConnectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQuickConnectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRoutingProfileQueues(ctx context.Context, input *connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRoutingProfiles(ctx context.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRoutingProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSecurityKeys(ctx context.Context, input *connect.ListSecurityKeysInput) (*connect.ListSecurityKeysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSecurityKeysWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSecurityProfiles(ctx context.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSecurityProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUseCases(ctx context.Context, input *connect.ListUseCasesInput) (*connect.ListUseCasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUseCasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUserHierarchyGroups(ctx context.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUserHierarchyGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUsers(ctx context.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResumeContactRecording(ctx context.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResumeContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartChatContact(ctx context.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartChatContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartContactRecording(ctx context.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartOutboundVoiceContact(ctx context.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartOutboundVoiceContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartTaskContact(ctx context.Context, input *connect.StartTaskContactInput) (*connect.StartTaskContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartTaskContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopContact(ctx context.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopContactRecording(ctx context.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SuspendContactRecording(ctx context.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SuspendContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactAttributes(ctx context.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactFlowContent(ctx context.Context, input *connect.UpdateContactFlowContentInput) (*connect.UpdateContactFlowContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactFlowContentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactFlowName(ctx context.Context, input *connect.UpdateContactFlowNameInput) (*connect.UpdateContactFlowNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactFlowNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInstanceAttribute(ctx context.Context, input *connect.UpdateInstanceAttributeInput) (*connect.UpdateInstanceAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInstanceAttributeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInstanceStorageConfig(ctx context.Context, input *connect.UpdateInstanceStorageConfigInput) (*connect.UpdateInstanceStorageConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInstanceStorageConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueueHoursOfOperation(ctx context.Context, input *connect.UpdateQueueHoursOfOperationInput) (*connect.UpdateQueueHoursOfOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueHoursOfOperationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueueMaxContacts(ctx context.Context, input *connect.UpdateQueueMaxContactsInput) (*connect.UpdateQueueMaxContactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueMaxContactsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueueName(ctx context.Context, input *connect.UpdateQueueNameInput) (*connect.UpdateQueueNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueueOutboundCallerConfig(ctx context.Context, input *connect.UpdateQueueOutboundCallerConfigInput) (*connect.UpdateQueueOutboundCallerConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueOutboundCallerConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueueStatus(ctx context.Context, input *connect.UpdateQueueStatusInput) (*connect.UpdateQueueStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQuickConnectConfig(ctx context.Context, input *connect.UpdateQuickConnectConfigInput) (*connect.UpdateQuickConnectConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQuickConnectConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQuickConnectName(ctx context.Context, input *connect.UpdateQuickConnectNameInput) (*connect.UpdateQuickConnectNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQuickConnectNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileConcurrency(ctx context.Context, input *connect.UpdateRoutingProfileConcurrencyInput) (*connect.UpdateRoutingProfileConcurrencyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileConcurrencyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileDefaultOutboundQueue(ctx context.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileDefaultOutboundQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileName(ctx context.Context, input *connect.UpdateRoutingProfileNameInput) (*connect.UpdateRoutingProfileNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileQueues(ctx context.Context, input *connect.UpdateRoutingProfileQueuesInput) (*connect.UpdateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserHierarchy(ctx context.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserHierarchyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserHierarchyGroupName(ctx context.Context, input *connect.UpdateUserHierarchyGroupNameInput) (*connect.UpdateUserHierarchyGroupNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserHierarchyGroupNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserHierarchyStructure(ctx context.Context, input *connect.UpdateUserHierarchyStructureInput) (*connect.UpdateUserHierarchyStructureOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserHierarchyStructureWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserIdentityInfo(ctx context.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserIdentityInfoWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserPhoneConfig(ctx context.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserPhoneConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserRoutingProfile(ctx context.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserSecurityProfiles(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserSecurityProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}