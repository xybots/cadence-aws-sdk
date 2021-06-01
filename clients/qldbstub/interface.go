// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package qldbstub

import (
	"github.com/aws/aws-sdk-go/service/qldb"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelJournalKinesisStream(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error)
	CancelJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.CancelJournalKinesisStreamInput) *CancelJournalKinesisStreamFuture

	CreateLedger(ctx workflow.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error)
	CreateLedgerAsync(ctx workflow.Context, input *qldb.CreateLedgerInput) *CreateLedgerFuture

	DeleteLedger(ctx workflow.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error)
	DeleteLedgerAsync(ctx workflow.Context, input *qldb.DeleteLedgerInput) *DeleteLedgerFuture

	DescribeJournalKinesisStream(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error)
	DescribeJournalKinesisStreamAsync(ctx workflow.Context, input *qldb.DescribeJournalKinesisStreamInput) *DescribeJournalKinesisStreamFuture

	DescribeJournalS3Export(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error)
	DescribeJournalS3ExportAsync(ctx workflow.Context, input *qldb.DescribeJournalS3ExportInput) *DescribeJournalS3ExportFuture

	DescribeLedger(ctx workflow.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error)
	DescribeLedgerAsync(ctx workflow.Context, input *qldb.DescribeLedgerInput) *DescribeLedgerFuture

	ExportJournalToS3(ctx workflow.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error)
	ExportJournalToS3Async(ctx workflow.Context, input *qldb.ExportJournalToS3Input) *ExportJournalToS3Future

	GetBlock(ctx workflow.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error)
	GetBlockAsync(ctx workflow.Context, input *qldb.GetBlockInput) *GetBlockFuture

	GetDigest(ctx workflow.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error)
	GetDigestAsync(ctx workflow.Context, input *qldb.GetDigestInput) *GetDigestFuture

	GetRevision(ctx workflow.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error)
	GetRevisionAsync(ctx workflow.Context, input *qldb.GetRevisionInput) *GetRevisionFuture

	ListJournalKinesisStreamsForLedger(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error)
	ListJournalKinesisStreamsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) *ListJournalKinesisStreamsForLedgerFuture

	ListJournalS3Exports(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error)
	ListJournalS3ExportsAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsInput) *ListJournalS3ExportsFuture

	ListJournalS3ExportsForLedger(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error)
	ListJournalS3ExportsForLedgerAsync(ctx workflow.Context, input *qldb.ListJournalS3ExportsForLedgerInput) *ListJournalS3ExportsForLedgerFuture

	ListLedgers(ctx workflow.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error)
	ListLedgersAsync(ctx workflow.Context, input *qldb.ListLedgersInput) *ListLedgersFuture

	ListTagsForResource(ctx workflow.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *qldb.ListTagsForResourceInput) *ListTagsForResourceFuture

	StreamJournalToKinesis(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error)
	StreamJournalToKinesisAsync(ctx workflow.Context, input *qldb.StreamJournalToKinesisInput) *StreamJournalToKinesisFuture

	TagResource(ctx workflow.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *qldb.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *qldb.UntagResourceInput) *UntagResourceFuture

	UpdateLedger(ctx workflow.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error)
	UpdateLedgerAsync(ctx workflow.Context, input *qldb.UpdateLedgerInput) *UpdateLedgerFuture
}

func NewClient() Client {
	return &stub{}
}