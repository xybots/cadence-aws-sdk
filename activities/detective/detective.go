// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"

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
	client detectiveiface.DetectiveAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := detective.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (detectiveiface.DetectiveAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return detective.New(sess), nil
}

func (a *Activities) AcceptInvitation(ctx context.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AcceptInvitationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGraph(ctx context.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGraphWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMembers(ctx context.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMembersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGraph(ctx context.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGraphWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMembers(ctx context.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMembersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateMembership(ctx context.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateMembershipWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMembers(ctx context.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMembersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGraphs(ctx context.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGraphsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInvitations(ctx context.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInvitationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMembers(ctx context.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMembersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectInvitation(ctx context.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectInvitationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartMonitoringMember(ctx context.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartMonitoringMemberWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
