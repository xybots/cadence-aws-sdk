// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackage/mediapackageiface"

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
	client mediapackageiface.MediaPackageAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mediapackage.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mediapackageiface.MediaPackageAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return mediapackage.New(sess), nil
}

func (a *Activities) ConfigureLogs(ctx context.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ConfigureLogsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateChannel(ctx context.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHarvestJob(ctx context.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHarvestJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateOriginEndpoint(ctx context.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateOriginEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteChannel(ctx context.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteOriginEndpoint(ctx context.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteOriginEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeChannel(ctx context.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeHarvestJob(ctx context.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeHarvestJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeOriginEndpoint(ctx context.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeOriginEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListChannels(ctx context.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListChannelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHarvestJobs(ctx context.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHarvestJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListOriginEndpoints(ctx context.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListOriginEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RotateChannelCredentials(ctx context.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RotateChannelCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RotateIngestEndpointCredentials(ctx context.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RotateIngestEndpointCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateChannel(ctx context.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateOriginEndpoint(ctx context.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateOriginEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
