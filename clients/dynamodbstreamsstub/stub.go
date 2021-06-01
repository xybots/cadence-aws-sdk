// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package dynamodbstreamsstub

import (
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DescribeStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStreamFuture) Get(ctx workflow.Context) (*dynamodbstreams.DescribeStreamOutput, error) {
	var output dynamodbstreams.DescribeStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRecordsFuture) Get(ctx workflow.Context) (*dynamodbstreams.GetRecordsOutput, error) {
	var output dynamodbstreams.GetRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetShardIteratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetShardIteratorFuture) Get(ctx workflow.Context) (*dynamodbstreams.GetShardIteratorOutput, error) {
	var output dynamodbstreams.GetShardIteratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListStreamsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListStreamsFuture) Get(ctx workflow.Context) (*dynamodbstreams.ListStreamsOutput, error) {
	var output dynamodbstreams.ListStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStream(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
	var output dynamodbstreams.DescribeStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-DescribeStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamAsync(ctx workflow.Context, input *dynamodbstreams.DescribeStreamInput) *DescribeStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-DescribeStream", input)
	return &DescribeStreamFuture{Future: future}
}

func (a *stub) GetRecords(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
	var output dynamodbstreams.GetRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-GetRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecordsAsync(ctx workflow.Context, input *dynamodbstreams.GetRecordsInput) *GetRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-GetRecords", input)
	return &GetRecordsFuture{Future: future}
}

func (a *stub) GetShardIterator(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
	var output dynamodbstreams.GetShardIteratorOutput
	err := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-GetShardIterator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetShardIteratorAsync(ctx workflow.Context, input *dynamodbstreams.GetShardIteratorInput) *GetShardIteratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-GetShardIterator", input)
	return &GetShardIteratorFuture{Future: future}
}

func (a *stub) ListStreams(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
	var output dynamodbstreams.ListStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-ListStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamsAsync(ctx workflow.Context, input *dynamodbstreams.ListStreamsInput) *ListStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dynamodbstreams-ListStreams", input)
	return &ListStreamsFuture{Future: future}
}
