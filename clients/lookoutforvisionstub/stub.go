// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package lookoutforvisionstub

import (
	"github.com/aws/aws-sdk-go/service/lookoutforvision"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDatasetFuture) Get(ctx workflow.Context) (*lookoutforvision.CreateDatasetOutput, error) {
	var output lookoutforvision.CreateDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateModelFuture) Get(ctx workflow.Context) (*lookoutforvision.CreateModelOutput, error) {
	var output lookoutforvision.CreateModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProjectFuture) Get(ctx workflow.Context) (*lookoutforvision.CreateProjectOutput, error) {
	var output lookoutforvision.CreateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDatasetFuture) Get(ctx workflow.Context) (*lookoutforvision.DeleteDatasetOutput, error) {
	var output lookoutforvision.DeleteDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteModelFuture) Get(ctx workflow.Context) (*lookoutforvision.DeleteModelOutput, error) {
	var output lookoutforvision.DeleteModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProjectFuture) Get(ctx workflow.Context) (*lookoutforvision.DeleteProjectOutput, error) {
	var output lookoutforvision.DeleteProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDatasetFuture) Get(ctx workflow.Context) (*lookoutforvision.DescribeDatasetOutput, error) {
	var output lookoutforvision.DescribeDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeModelFuture) Get(ctx workflow.Context) (*lookoutforvision.DescribeModelOutput, error) {
	var output lookoutforvision.DescribeModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeProjectFuture) Get(ctx workflow.Context) (*lookoutforvision.DescribeProjectOutput, error) {
	var output lookoutforvision.DescribeProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DetectAnomaliesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DetectAnomaliesFuture) Get(ctx workflow.Context) (*lookoutforvision.DetectAnomaliesOutput, error) {
	var output lookoutforvision.DetectAnomaliesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDatasetEntriesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDatasetEntriesFuture) Get(ctx workflow.Context) (*lookoutforvision.ListDatasetEntriesOutput, error) {
	var output lookoutforvision.ListDatasetEntriesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListModelsFuture) Get(ctx workflow.Context) (*lookoutforvision.ListModelsOutput, error) {
	var output lookoutforvision.ListModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProjectsFuture) Get(ctx workflow.Context) (*lookoutforvision.ListProjectsOutput, error) {
	var output lookoutforvision.ListProjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartModelFuture) Get(ctx workflow.Context) (*lookoutforvision.StartModelOutput, error) {
	var output lookoutforvision.StartModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopModelFuture) Get(ctx workflow.Context) (*lookoutforvision.StopModelOutput, error) {
	var output lookoutforvision.StopModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDatasetEntriesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDatasetEntriesFuture) Get(ctx workflow.Context) (*lookoutforvision.UpdateDatasetEntriesOutput, error) {
	var output lookoutforvision.UpdateDatasetEntriesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataset(ctx workflow.Context, input *lookoutforvision.CreateDatasetInput) (*lookoutforvision.CreateDatasetOutput, error) {
	var output lookoutforvision.CreateDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatasetAsync(ctx workflow.Context, input *lookoutforvision.CreateDatasetInput) *CreateDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateDataset", input)
	return &CreateDatasetFuture{Future: future}
}

func (a *stub) CreateModel(ctx workflow.Context, input *lookoutforvision.CreateModelInput) (*lookoutforvision.CreateModelOutput, error) {
	var output lookoutforvision.CreateModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateModelAsync(ctx workflow.Context, input *lookoutforvision.CreateModelInput) *CreateModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateModel", input)
	return &CreateModelFuture{Future: future}
}

func (a *stub) CreateProject(ctx workflow.Context, input *lookoutforvision.CreateProjectInput) (*lookoutforvision.CreateProjectOutput, error) {
	var output lookoutforvision.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProjectAsync(ctx workflow.Context, input *lookoutforvision.CreateProjectInput) *CreateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-CreateProject", input)
	return &CreateProjectFuture{Future: future}
}

func (a *stub) DeleteDataset(ctx workflow.Context, input *lookoutforvision.DeleteDatasetInput) (*lookoutforvision.DeleteDatasetOutput, error) {
	var output lookoutforvision.DeleteDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetAsync(ctx workflow.Context, input *lookoutforvision.DeleteDatasetInput) *DeleteDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteDataset", input)
	return &DeleteDatasetFuture{Future: future}
}

func (a *stub) DeleteModel(ctx workflow.Context, input *lookoutforvision.DeleteModelInput) (*lookoutforvision.DeleteModelOutput, error) {
	var output lookoutforvision.DeleteModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteModelAsync(ctx workflow.Context, input *lookoutforvision.DeleteModelInput) *DeleteModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteModel", input)
	return &DeleteModelFuture{Future: future}
}

func (a *stub) DeleteProject(ctx workflow.Context, input *lookoutforvision.DeleteProjectInput) (*lookoutforvision.DeleteProjectOutput, error) {
	var output lookoutforvision.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProjectAsync(ctx workflow.Context, input *lookoutforvision.DeleteProjectInput) *DeleteProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DeleteProject", input)
	return &DeleteProjectFuture{Future: future}
}

func (a *stub) DescribeDataset(ctx workflow.Context, input *lookoutforvision.DescribeDatasetInput) (*lookoutforvision.DescribeDatasetOutput, error) {
	var output lookoutforvision.DescribeDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetAsync(ctx workflow.Context, input *lookoutforvision.DescribeDatasetInput) *DescribeDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeDataset", input)
	return &DescribeDatasetFuture{Future: future}
}

func (a *stub) DescribeModel(ctx workflow.Context, input *lookoutforvision.DescribeModelInput) (*lookoutforvision.DescribeModelOutput, error) {
	var output lookoutforvision.DescribeModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeModelAsync(ctx workflow.Context, input *lookoutforvision.DescribeModelInput) *DescribeModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeModel", input)
	return &DescribeModelFuture{Future: future}
}

func (a *stub) DescribeProject(ctx workflow.Context, input *lookoutforvision.DescribeProjectInput) (*lookoutforvision.DescribeProjectOutput, error) {
	var output lookoutforvision.DescribeProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProjectAsync(ctx workflow.Context, input *lookoutforvision.DescribeProjectInput) *DescribeProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DescribeProject", input)
	return &DescribeProjectFuture{Future: future}
}

func (a *stub) DetectAnomalies(ctx workflow.Context, input *lookoutforvision.DetectAnomaliesInput) (*lookoutforvision.DetectAnomaliesOutput, error) {
	var output lookoutforvision.DetectAnomaliesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DetectAnomalies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DetectAnomaliesAsync(ctx workflow.Context, input *lookoutforvision.DetectAnomaliesInput) *DetectAnomaliesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-DetectAnomalies", input)
	return &DetectAnomaliesFuture{Future: future}
}

func (a *stub) ListDatasetEntries(ctx workflow.Context, input *lookoutforvision.ListDatasetEntriesInput) (*lookoutforvision.ListDatasetEntriesOutput, error) {
	var output lookoutforvision.ListDatasetEntriesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListDatasetEntries", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetEntriesAsync(ctx workflow.Context, input *lookoutforvision.ListDatasetEntriesInput) *ListDatasetEntriesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListDatasetEntries", input)
	return &ListDatasetEntriesFuture{Future: future}
}

func (a *stub) ListModels(ctx workflow.Context, input *lookoutforvision.ListModelsInput) (*lookoutforvision.ListModelsOutput, error) {
	var output lookoutforvision.ListModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListModelsAsync(ctx workflow.Context, input *lookoutforvision.ListModelsInput) *ListModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListModels", input)
	return &ListModelsFuture{Future: future}
}

func (a *stub) ListProjects(ctx workflow.Context, input *lookoutforvision.ListProjectsInput) (*lookoutforvision.ListProjectsOutput, error) {
	var output lookoutforvision.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListProjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProjectsAsync(ctx workflow.Context, input *lookoutforvision.ListProjectsInput) *ListProjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-ListProjects", input)
	return &ListProjectsFuture{Future: future}
}

func (a *stub) StartModel(ctx workflow.Context, input *lookoutforvision.StartModelInput) (*lookoutforvision.StartModelOutput, error) {
	var output lookoutforvision.StartModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-StartModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartModelAsync(ctx workflow.Context, input *lookoutforvision.StartModelInput) *StartModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-StartModel", input)
	return &StartModelFuture{Future: future}
}

func (a *stub) StopModel(ctx workflow.Context, input *lookoutforvision.StopModelInput) (*lookoutforvision.StopModelOutput, error) {
	var output lookoutforvision.StopModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-StopModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopModelAsync(ctx workflow.Context, input *lookoutforvision.StopModelInput) *StopModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-StopModel", input)
	return &StopModelFuture{Future: future}
}

func (a *stub) UpdateDatasetEntries(ctx workflow.Context, input *lookoutforvision.UpdateDatasetEntriesInput) (*lookoutforvision.UpdateDatasetEntriesOutput, error) {
	var output lookoutforvision.UpdateDatasetEntriesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-UpdateDatasetEntries", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDatasetEntriesAsync(ctx workflow.Context, input *lookoutforvision.UpdateDatasetEntriesInput) *UpdateDatasetEntriesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lookoutforvision-UpdateDatasetEntries", input)
	return &UpdateDatasetEntriesFuture{Future: future}
}
