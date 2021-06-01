// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"

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
	client fmsiface.FMSAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := fms.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (fmsiface.FMSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return fms.New(sess), nil
}

func (a *Activities) AssociateAdminAccount(ctx context.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateAdminAccountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAppsList(ctx context.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAppsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNotificationChannel(ctx context.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNotificationChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePolicy(ctx context.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProtocolsList(ctx context.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProtocolsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateAdminAccount(ctx context.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateAdminAccountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAdminAccount(ctx context.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAdminAccountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAppsList(ctx context.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetComplianceDetail(ctx context.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetComplianceDetailWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNotificationChannel(ctx context.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNotificationChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPolicy(ctx context.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetProtectionStatus(ctx context.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetProtectionStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetProtocolsList(ctx context.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetProtocolsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetViolationDetails(ctx context.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetViolationDetailsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAppsLists(ctx context.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAppsListsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListComplianceStatus(ctx context.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListComplianceStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMemberAccounts(ctx context.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMemberAccountsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPolicies(ctx context.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProtocolsLists(ctx context.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProtocolsListsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAppsList(ctx context.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAppsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutNotificationChannel(ctx context.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutNotificationChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutPolicy(ctx context.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutProtocolsList(ctx context.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutProtocolsListWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
