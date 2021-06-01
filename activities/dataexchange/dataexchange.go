// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"

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
	client dataexchangeiface.DataExchangeAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := dataexchange.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (dataexchangeiface.DataExchangeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return dataexchange.New(sess), nil
}

func (a *Activities) CancelJob(ctx context.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDataSet(ctx context.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDataSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateJob(ctx context.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRevision(ctx context.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAsset(ctx context.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAssetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDataSet(ctx context.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDataSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRevision(ctx context.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAsset(ctx context.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAssetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDataSet(ctx context.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDataSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetJob(ctx context.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRevision(ctx context.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDataSetRevisions(ctx context.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDataSetRevisionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDataSets(ctx context.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDataSetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobs(ctx context.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRevisionAssets(ctx context.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRevisionAssetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartJob(ctx context.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateAsset(ctx context.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAssetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDataSet(ctx context.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDataSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRevision(ctx context.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
