// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package translatestub

import (
	"github.com/aws/aws-sdk-go/service/translate"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateParallelDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateParallelDataFuture) Get(ctx workflow.Context) (*translate.CreateParallelDataOutput, error) {
	var output translate.CreateParallelDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteParallelDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteParallelDataFuture) Get(ctx workflow.Context) (*translate.DeleteParallelDataOutput, error) {
	var output translate.DeleteParallelDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTerminologyFuture) Get(ctx workflow.Context) (*translate.DeleteTerminologyOutput, error) {
	var output translate.DeleteTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.DescribeTextTranslationJobOutput, error) {
	var output translate.DescribeTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetParallelDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetParallelDataFuture) Get(ctx workflow.Context) (*translate.GetParallelDataOutput, error) {
	var output translate.GetParallelDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTerminologyFuture) Get(ctx workflow.Context) (*translate.GetTerminologyOutput, error) {
	var output translate.GetTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ImportTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ImportTerminologyFuture) Get(ctx workflow.Context) (*translate.ImportTerminologyOutput, error) {
	var output translate.ImportTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListParallelDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListParallelDataFuture) Get(ctx workflow.Context) (*translate.ListParallelDataOutput, error) {
	var output translate.ListParallelDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTerminologiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTerminologiesFuture) Get(ctx workflow.Context) (*translate.ListTerminologiesOutput, error) {
	var output translate.ListTerminologiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTextTranslationJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTextTranslationJobsFuture) Get(ctx workflow.Context) (*translate.ListTextTranslationJobsOutput, error) {
	var output translate.ListTextTranslationJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.StartTextTranslationJobOutput, error) {
	var output translate.StartTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.StopTextTranslationJobOutput, error) {
	var output translate.StopTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextFuture) Get(ctx workflow.Context) (*translate.TextOutput, error) {
	var output translate.TextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateParallelDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateParallelDataFuture) Get(ctx workflow.Context) (*translate.UpdateParallelDataOutput, error) {
	var output translate.UpdateParallelDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateParallelData(ctx workflow.Context, input *translate.CreateParallelDataInput) (*translate.CreateParallelDataOutput, error) {
	var output translate.CreateParallelDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-CreateParallelData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateParallelDataAsync(ctx workflow.Context, input *translate.CreateParallelDataInput) *CreateParallelDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-CreateParallelData", input)
	return &CreateParallelDataFuture{Future: future}
}

func (a *stub) DeleteParallelData(ctx workflow.Context, input *translate.DeleteParallelDataInput) (*translate.DeleteParallelDataOutput, error) {
	var output translate.DeleteParallelDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-DeleteParallelData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteParallelDataAsync(ctx workflow.Context, input *translate.DeleteParallelDataInput) *DeleteParallelDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-DeleteParallelData", input)
	return &DeleteParallelDataFuture{Future: future}
}

func (a *stub) DeleteTerminology(ctx workflow.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
	var output translate.DeleteTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-DeleteTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTerminologyAsync(ctx workflow.Context, input *translate.DeleteTerminologyInput) *DeleteTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-DeleteTerminology", input)
	return &DeleteTerminologyFuture{Future: future}
}

func (a *stub) DescribeTextTranslationJob(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
	var output translate.DescribeTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-DescribeTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTextTranslationJobAsync(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) *DescribeTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-DescribeTextTranslationJob", input)
	return &DescribeTextTranslationJobFuture{Future: future}
}

func (a *stub) GetParallelData(ctx workflow.Context, input *translate.GetParallelDataInput) (*translate.GetParallelDataOutput, error) {
	var output translate.GetParallelDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-GetParallelData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetParallelDataAsync(ctx workflow.Context, input *translate.GetParallelDataInput) *GetParallelDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-GetParallelData", input)
	return &GetParallelDataFuture{Future: future}
}

func (a *stub) GetTerminology(ctx workflow.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
	var output translate.GetTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-GetTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTerminologyAsync(ctx workflow.Context, input *translate.GetTerminologyInput) *GetTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-GetTerminology", input)
	return &GetTerminologyFuture{Future: future}
}

func (a *stub) ImportTerminology(ctx workflow.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
	var output translate.ImportTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ImportTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportTerminologyAsync(ctx workflow.Context, input *translate.ImportTerminologyInput) *ImportTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ImportTerminology", input)
	return &ImportTerminologyFuture{Future: future}
}

func (a *stub) ListParallelData(ctx workflow.Context, input *translate.ListParallelDataInput) (*translate.ListParallelDataOutput, error) {
	var output translate.ListParallelDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ListParallelData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListParallelDataAsync(ctx workflow.Context, input *translate.ListParallelDataInput) *ListParallelDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ListParallelData", input)
	return &ListParallelDataFuture{Future: future}
}

func (a *stub) ListTerminologies(ctx workflow.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
	var output translate.ListTerminologiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ListTerminologies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTerminologiesAsync(ctx workflow.Context, input *translate.ListTerminologiesInput) *ListTerminologiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ListTerminologies", input)
	return &ListTerminologiesFuture{Future: future}
}

func (a *stub) ListTextTranslationJobs(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
	var output translate.ListTextTranslationJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ListTextTranslationJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTextTranslationJobsAsync(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) *ListTextTranslationJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ListTextTranslationJobs", input)
	return &ListTextTranslationJobsFuture{Future: future}
}

func (a *stub) StartTextTranslationJob(ctx workflow.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
	var output translate.StartTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-StartTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartTextTranslationJobAsync(ctx workflow.Context, input *translate.StartTextTranslationJobInput) *StartTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-StartTextTranslationJob", input)
	return &StartTextTranslationJobFuture{Future: future}
}

func (a *stub) StopTextTranslationJob(ctx workflow.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
	var output translate.StopTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-StopTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopTextTranslationJobAsync(ctx workflow.Context, input *translate.StopTextTranslationJobInput) *StopTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-StopTextTranslationJob", input)
	return &StopTextTranslationJobFuture{Future: future}
}

func (a *stub) Text(ctx workflow.Context, input *translate.TextInput) (*translate.TextOutput, error) {
	var output translate.TextOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-Text", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TextAsync(ctx workflow.Context, input *translate.TextInput) *TextFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-Text", input)
	return &TextFuture{Future: future}
}

func (a *stub) UpdateParallelData(ctx workflow.Context, input *translate.UpdateParallelDataInput) (*translate.UpdateParallelDataOutput, error) {
	var output translate.UpdateParallelDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-UpdateParallelData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateParallelDataAsync(ctx workflow.Context, input *translate.UpdateParallelDataInput) *UpdateParallelDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-UpdateParallelData", input)
	return &UpdateParallelDataFuture{Future: future}
}