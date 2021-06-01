// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"

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
	client iotanalyticsiface.IoTAnalyticsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := iotanalytics.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (iotanalyticsiface.IoTAnalyticsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return iotanalytics.New(sess), nil
}

func (a *Activities) BatchPutMessage(ctx context.Context, input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchPutMessageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelPipelineReprocessing(ctx context.Context, input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelPipelineReprocessingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateChannel(ctx context.Context, input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDataset(ctx context.Context, input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDatasetContent(ctx context.Context, input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatasetContentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDatastore(ctx context.Context, input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatastoreWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePipeline(ctx context.Context, input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteChannel(ctx context.Context, input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDataset(ctx context.Context, input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDatasetContent(ctx context.Context, input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDatasetContentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDatastore(ctx context.Context, input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDatastoreWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePipeline(ctx context.Context, input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeChannel(ctx context.Context, input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDataset(ctx context.Context, input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDatastore(ctx context.Context, input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDatastoreWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoggingOptions(ctx context.Context, input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoggingOptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePipeline(ctx context.Context, input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDatasetContent(ctx context.Context, input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDatasetContentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListChannels(ctx context.Context, input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListChannelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatasetContents(ctx context.Context, input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatasetContentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatasets(ctx context.Context, input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatasetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatastores(ctx context.Context, input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatastoresWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPipelines(ctx context.Context, input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPipelinesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutLoggingOptions(ctx context.Context, input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutLoggingOptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RunPipelineActivity(ctx context.Context, input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RunPipelineActivityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SampleChannelData(ctx context.Context, input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SampleChannelDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartPipelineReprocessing(ctx context.Context, input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartPipelineReprocessingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateChannel(ctx context.Context, input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDataset(ctx context.Context, input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDatastore(ctx context.Context, input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDatastoreWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePipeline(ctx context.Context, input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
