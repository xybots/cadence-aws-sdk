// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package datasyncstub

import (
	"github.com/aws/aws-sdk-go/service/datasync"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error)
	CancelTaskExecutionAsync(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) *CancelTaskExecutionFuture

	CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error)
	CreateAgentAsync(ctx workflow.Context, input *datasync.CreateAgentInput) *CreateAgentFuture

	CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationEfsAsync(ctx workflow.Context, input *datasync.CreateLocationEfsInput) *CreateLocationEfsFuture

	CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) *CreateLocationFsxWindowsFuture

	CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationNfsAsync(ctx workflow.Context, input *datasync.CreateLocationNfsInput) *CreateLocationNfsFuture

	CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationObjectStorageAsync(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) *CreateLocationObjectStorageFuture

	CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error)
	CreateLocationS3Async(ctx workflow.Context, input *datasync.CreateLocationS3Input) *CreateLocationS3Future

	CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error)
	CreateLocationSmbAsync(ctx workflow.Context, input *datasync.CreateLocationSmbInput) *CreateLocationSmbFuture

	CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error)
	CreateTaskAsync(ctx workflow.Context, input *datasync.CreateTaskInput) *CreateTaskFuture

	DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error)
	DeleteAgentAsync(ctx workflow.Context, input *datasync.DeleteAgentInput) *DeleteAgentFuture

	DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error)
	DeleteLocationAsync(ctx workflow.Context, input *datasync.DeleteLocationInput) *DeleteLocationFuture

	DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error)
	DeleteTaskAsync(ctx workflow.Context, input *datasync.DeleteTaskInput) *DeleteTaskFuture

	DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error)
	DescribeAgentAsync(ctx workflow.Context, input *datasync.DescribeAgentInput) *DescribeAgentFuture

	DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationEfsAsync(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) *DescribeLocationEfsFuture

	DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) *DescribeLocationFsxWindowsFuture

	DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationNfsAsync(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) *DescribeLocationNfsFuture

	DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationObjectStorageAsync(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) *DescribeLocationObjectStorageFuture

	DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationS3Async(ctx workflow.Context, input *datasync.DescribeLocationS3Input) *DescribeLocationS3Future

	DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error)
	DescribeLocationSmbAsync(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) *DescribeLocationSmbFuture

	DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error)
	DescribeTaskAsync(ctx workflow.Context, input *datasync.DescribeTaskInput) *DescribeTaskFuture

	DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error)
	DescribeTaskExecutionAsync(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) *DescribeTaskExecutionFuture

	ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error)
	ListAgentsAsync(ctx workflow.Context, input *datasync.ListAgentsInput) *ListAgentsFuture

	ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error)
	ListLocationsAsync(ctx workflow.Context, input *datasync.ListLocationsInput) *ListLocationsFuture

	ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *datasync.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error)
	ListTaskExecutionsAsync(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) *ListTaskExecutionsFuture

	ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error)
	ListTasksAsync(ctx workflow.Context, input *datasync.ListTasksInput) *ListTasksFuture

	StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error)
	StartTaskExecutionAsync(ctx workflow.Context, input *datasync.StartTaskExecutionInput) *StartTaskExecutionFuture

	TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *datasync.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *datasync.UntagResourceInput) *UntagResourceFuture

	UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error)
	UpdateAgentAsync(ctx workflow.Context, input *datasync.UpdateAgentInput) *UpdateAgentFuture

	UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error)
	UpdateTaskAsync(ctx workflow.Context, input *datasync.UpdateTaskInput) *UpdateTaskFuture

	UpdateTaskExecution(ctx workflow.Context, input *datasync.UpdateTaskExecutionInput) (*datasync.UpdateTaskExecutionOutput, error)
	UpdateTaskExecutionAsync(ctx workflow.Context, input *datasync.UpdateTaskExecutionInput) *UpdateTaskExecutionFuture
}

func NewClient() Client {
	return &stub{}
}
