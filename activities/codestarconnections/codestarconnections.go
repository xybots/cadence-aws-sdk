// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codestarconnections

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"

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
	client codestarconnectionsiface.CodeStarConnectionsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := codestarconnections.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (codestarconnectionsiface.CodeStarConnectionsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return codestarconnections.New(sess), nil
}

func (a *Activities) CreateConnection(ctx context.Context, input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHost(ctx context.Context, input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHostWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConnection(ctx context.Context, input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteHost(ctx context.Context, input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteHostWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConnection(ctx context.Context, input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetHost(ctx context.Context, input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetHostWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListConnections(ctx context.Context, input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHosts(ctx context.Context, input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHostsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateHost(ctx context.Context, input *codestarconnections.UpdateHostInput) (*codestarconnections.UpdateHostOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateHostWithContext(ctx, input)

	return output, internal.EncodeError(err)
}