// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package glacierstub

import (
	"github.com/aws/aws-sdk-go/service/glacier"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AbortMultipartUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AbortMultipartUploadFuture) Get(ctx workflow.Context) (*glacier.AbortMultipartUploadOutput, error) {
	var output glacier.AbortMultipartUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AbortVaultLockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AbortVaultLockFuture) Get(ctx workflow.Context) (*glacier.AbortVaultLockOutput, error) {
	var output glacier.AbortVaultLockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddTagsToVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsToVaultFuture) Get(ctx workflow.Context) (*glacier.AddTagsToVaultOutput, error) {
	var output glacier.AddTagsToVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CompleteMultipartUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CompleteMultipartUploadFuture) Get(ctx workflow.Context) (*glacier.ArchiveCreationOutput, error) {
	var output glacier.ArchiveCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CompleteVaultLockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CompleteVaultLockFuture) Get(ctx workflow.Context) (*glacier.CompleteVaultLockOutput, error) {
	var output glacier.CompleteVaultLockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateVaultFuture) Get(ctx workflow.Context) (*glacier.CreateVaultOutput, error) {
	var output glacier.CreateVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteArchiveFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteArchiveFuture) Get(ctx workflow.Context) (*glacier.DeleteArchiveOutput, error) {
	var output glacier.DeleteArchiveOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteVaultFuture) Get(ctx workflow.Context) (*glacier.DeleteVaultOutput, error) {
	var output glacier.DeleteVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteVaultAccessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteVaultAccessPolicyFuture) Get(ctx workflow.Context) (*glacier.DeleteVaultAccessPolicyOutput, error) {
	var output glacier.DeleteVaultAccessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteVaultNotificationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteVaultNotificationsFuture) Get(ctx workflow.Context) (*glacier.DeleteVaultNotificationsOutput, error) {
	var output glacier.DeleteVaultNotificationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobFuture) Get(ctx workflow.Context) (*glacier.JobDescription, error) {
	var output glacier.JobDescription
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeVaultFuture) Get(ctx workflow.Context) (*glacier.DescribeVaultOutput, error) {
	var output glacier.DescribeVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDataRetrievalPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDataRetrievalPolicyFuture) Get(ctx workflow.Context) (*glacier.GetDataRetrievalPolicyOutput, error) {
	var output glacier.GetDataRetrievalPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetJobOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetJobOutputFuture) Get(ctx workflow.Context) (*glacier.GetJobOutputOutput, error) {
	var output glacier.GetJobOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetVaultAccessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetVaultAccessPolicyFuture) Get(ctx workflow.Context) (*glacier.GetVaultAccessPolicyOutput, error) {
	var output glacier.GetVaultAccessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetVaultLockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetVaultLockFuture) Get(ctx workflow.Context) (*glacier.GetVaultLockOutput, error) {
	var output glacier.GetVaultLockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetVaultNotificationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetVaultNotificationsFuture) Get(ctx workflow.Context) (*glacier.GetVaultNotificationsOutput, error) {
	var output glacier.GetVaultNotificationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InitiateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InitiateJobFuture) Get(ctx workflow.Context) (*glacier.InitiateJobOutput, error) {
	var output glacier.InitiateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InitiateMultipartUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InitiateMultipartUploadFuture) Get(ctx workflow.Context) (*glacier.InitiateMultipartUploadOutput, error) {
	var output glacier.InitiateMultipartUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InitiateVaultLockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InitiateVaultLockFuture) Get(ctx workflow.Context) (*glacier.InitiateVaultLockOutput, error) {
	var output glacier.InitiateVaultLockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobsFuture) Get(ctx workflow.Context) (*glacier.ListJobsOutput, error) {
	var output glacier.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMultipartUploadsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMultipartUploadsFuture) Get(ctx workflow.Context) (*glacier.ListMultipartUploadsOutput, error) {
	var output glacier.ListMultipartUploadsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPartsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPartsFuture) Get(ctx workflow.Context) (*glacier.ListPartsOutput, error) {
	var output glacier.ListPartsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProvisionedCapacityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProvisionedCapacityFuture) Get(ctx workflow.Context) (*glacier.ListProvisionedCapacityOutput, error) {
	var output glacier.ListProvisionedCapacityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForVaultFuture) Get(ctx workflow.Context) (*glacier.ListTagsForVaultOutput, error) {
	var output glacier.ListTagsForVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListVaultsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListVaultsFuture) Get(ctx workflow.Context) (*glacier.ListVaultsOutput, error) {
	var output glacier.ListVaultsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PurchaseProvisionedCapacityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PurchaseProvisionedCapacityFuture) Get(ctx workflow.Context) (*glacier.PurchaseProvisionedCapacityOutput, error) {
	var output glacier.PurchaseProvisionedCapacityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFromVaultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFromVaultFuture) Get(ctx workflow.Context) (*glacier.RemoveTagsFromVaultOutput, error) {
	var output glacier.RemoveTagsFromVaultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetDataRetrievalPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetDataRetrievalPolicyFuture) Get(ctx workflow.Context) (*glacier.SetDataRetrievalPolicyOutput, error) {
	var output glacier.SetDataRetrievalPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetVaultAccessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetVaultAccessPolicyFuture) Get(ctx workflow.Context) (*glacier.SetVaultAccessPolicyOutput, error) {
	var output glacier.SetVaultAccessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetVaultNotificationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetVaultNotificationsFuture) Get(ctx workflow.Context) (*glacier.SetVaultNotificationsOutput, error) {
	var output glacier.SetVaultNotificationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UploadArchiveFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UploadArchiveFuture) Get(ctx workflow.Context) (*glacier.ArchiveCreationOutput, error) {
	var output glacier.ArchiveCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UploadMultipartPartFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UploadMultipartPartFuture) Get(ctx workflow.Context) (*glacier.UploadMultipartPartOutput, error) {
	var output glacier.UploadMultipartPartOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AbortMultipartUpload(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
	var output glacier.AbortMultipartUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-AbortMultipartUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AbortMultipartUploadAsync(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) *AbortMultipartUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-AbortMultipartUpload", input)
	return &AbortMultipartUploadFuture{Future: future}
}

func (a *stub) AbortVaultLock(ctx workflow.Context, input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error) {
	var output glacier.AbortVaultLockOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-AbortVaultLock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AbortVaultLockAsync(ctx workflow.Context, input *glacier.AbortVaultLockInput) *AbortVaultLockFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-AbortVaultLock", input)
	return &AbortVaultLockFuture{Future: future}
}

func (a *stub) AddTagsToVault(ctx workflow.Context, input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error) {
	var output glacier.AddTagsToVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-AddTagsToVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToVaultAsync(ctx workflow.Context, input *glacier.AddTagsToVaultInput) *AddTagsToVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-AddTagsToVault", input)
	return &AddTagsToVaultFuture{Future: future}
}

func (a *stub) CompleteMultipartUpload(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error) {
	var output glacier.ArchiveCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-CompleteMultipartUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CompleteMultipartUploadAsync(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) *CompleteMultipartUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-CompleteMultipartUpload", input)
	return &CompleteMultipartUploadFuture{Future: future}
}

func (a *stub) CompleteVaultLock(ctx workflow.Context, input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error) {
	var output glacier.CompleteVaultLockOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-CompleteVaultLock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CompleteVaultLockAsync(ctx workflow.Context, input *glacier.CompleteVaultLockInput) *CompleteVaultLockFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-CompleteVaultLock", input)
	return &CompleteVaultLockFuture{Future: future}
}

func (a *stub) CreateVault(ctx workflow.Context, input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error) {
	var output glacier.CreateVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-CreateVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVaultAsync(ctx workflow.Context, input *glacier.CreateVaultInput) *CreateVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-CreateVault", input)
	return &CreateVaultFuture{Future: future}
}

func (a *stub) DeleteArchive(ctx workflow.Context, input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error) {
	var output glacier.DeleteArchiveOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteArchive", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteArchiveAsync(ctx workflow.Context, input *glacier.DeleteArchiveInput) *DeleteArchiveFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteArchive", input)
	return &DeleteArchiveFuture{Future: future}
}

func (a *stub) DeleteVault(ctx workflow.Context, input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error) {
	var output glacier.DeleteVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVaultAsync(ctx workflow.Context, input *glacier.DeleteVaultInput) *DeleteVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVault", input)
	return &DeleteVaultFuture{Future: future}
}

func (a *stub) DeleteVaultAccessPolicy(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error) {
	var output glacier.DeleteVaultAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVaultAccessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) *DeleteVaultAccessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVaultAccessPolicy", input)
	return &DeleteVaultAccessPolicyFuture{Future: future}
}

func (a *stub) DeleteVaultNotifications(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error) {
	var output glacier.DeleteVaultNotificationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVaultNotifications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVaultNotificationsAsync(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) *DeleteVaultNotificationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DeleteVaultNotifications", input)
	return &DeleteVaultNotificationsFuture{Future: future}
}

func (a *stub) DescribeJob(ctx workflow.Context, input *glacier.DescribeJobInput) (*glacier.JobDescription, error) {
	var output glacier.JobDescription
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DescribeJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobAsync(ctx workflow.Context, input *glacier.DescribeJobInput) *DescribeJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DescribeJob", input)
	return &DescribeJobFuture{Future: future}
}

func (a *stub) DescribeVault(ctx workflow.Context, input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
	var output glacier.DescribeVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-DescribeVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeVaultAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) *DescribeVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-DescribeVault", input)
	return &DescribeVaultFuture{Future: future}
}

func (a *stub) GetDataRetrievalPolicy(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
	var output glacier.GetDataRetrievalPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-GetDataRetrievalPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) *GetDataRetrievalPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-GetDataRetrievalPolicy", input)
	return &GetDataRetrievalPolicyFuture{Future: future}
}

func (a *stub) GetJobOutput(ctx workflow.Context, input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
	var output glacier.GetJobOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-GetJobOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobOutputAsync(ctx workflow.Context, input *glacier.GetJobOutputInput) *GetJobOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-GetJobOutput", input)
	return &GetJobOutputFuture{Future: future}
}

func (a *stub) GetVaultAccessPolicy(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
	var output glacier.GetVaultAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultAccessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) *GetVaultAccessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultAccessPolicy", input)
	return &GetVaultAccessPolicyFuture{Future: future}
}

func (a *stub) GetVaultLock(ctx workflow.Context, input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
	var output glacier.GetVaultLockOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultLock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetVaultLockAsync(ctx workflow.Context, input *glacier.GetVaultLockInput) *GetVaultLockFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultLock", input)
	return &GetVaultLockFuture{Future: future}
}

func (a *stub) GetVaultNotifications(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
	var output glacier.GetVaultNotificationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultNotifications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetVaultNotificationsAsync(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) *GetVaultNotificationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-GetVaultNotifications", input)
	return &GetVaultNotificationsFuture{Future: future}
}

func (a *stub) InitiateJob(ctx workflow.Context, input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error) {
	var output glacier.InitiateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitiateJobAsync(ctx workflow.Context, input *glacier.InitiateJobInput) *InitiateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateJob", input)
	return &InitiateJobFuture{Future: future}
}

func (a *stub) InitiateMultipartUpload(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error) {
	var output glacier.InitiateMultipartUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateMultipartUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitiateMultipartUploadAsync(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) *InitiateMultipartUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateMultipartUpload", input)
	return &InitiateMultipartUploadFuture{Future: future}
}

func (a *stub) InitiateVaultLock(ctx workflow.Context, input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error) {
	var output glacier.InitiateVaultLockOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateVaultLock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitiateVaultLockAsync(ctx workflow.Context, input *glacier.InitiateVaultLockInput) *InitiateVaultLockFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-InitiateVaultLock", input)
	return &InitiateVaultLockFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
	var output glacier.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *glacier.ListJobsInput) *ListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListJobs", input)
	return &ListJobsFuture{Future: future}
}

func (a *stub) ListMultipartUploads(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
	var output glacier.ListMultipartUploadsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListMultipartUploads", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMultipartUploadsAsync(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) *ListMultipartUploadsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListMultipartUploads", input)
	return &ListMultipartUploadsFuture{Future: future}
}

func (a *stub) ListParts(ctx workflow.Context, input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
	var output glacier.ListPartsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListParts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartsAsync(ctx workflow.Context, input *glacier.ListPartsInput) *ListPartsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListParts", input)
	return &ListPartsFuture{Future: future}
}

func (a *stub) ListProvisionedCapacity(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
	var output glacier.ListProvisionedCapacityOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListProvisionedCapacity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProvisionedCapacityAsync(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) *ListProvisionedCapacityFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListProvisionedCapacity", input)
	return &ListProvisionedCapacityFuture{Future: future}
}

func (a *stub) ListTagsForVault(ctx workflow.Context, input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
	var output glacier.ListTagsForVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListTagsForVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForVaultAsync(ctx workflow.Context, input *glacier.ListTagsForVaultInput) *ListTagsForVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListTagsForVault", input)
	return &ListTagsForVaultFuture{Future: future}
}

func (a *stub) ListVaults(ctx workflow.Context, input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
	var output glacier.ListVaultsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-ListVaults", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVaultsAsync(ctx workflow.Context, input *glacier.ListVaultsInput) *ListVaultsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-ListVaults", input)
	return &ListVaultsFuture{Future: future}
}

func (a *stub) PurchaseProvisionedCapacity(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error) {
	var output glacier.PurchaseProvisionedCapacityOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-PurchaseProvisionedCapacity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PurchaseProvisionedCapacityAsync(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) *PurchaseProvisionedCapacityFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-PurchaseProvisionedCapacity", input)
	return &PurchaseProvisionedCapacityFuture{Future: future}
}

func (a *stub) RemoveTagsFromVault(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error) {
	var output glacier.RemoveTagsFromVaultOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-RemoveTagsFromVault", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsFromVaultAsync(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) *RemoveTagsFromVaultFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-RemoveTagsFromVault", input)
	return &RemoveTagsFromVaultFuture{Future: future}
}

func (a *stub) SetDataRetrievalPolicy(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error) {
	var output glacier.SetDataRetrievalPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-SetDataRetrievalPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) *SetDataRetrievalPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-SetDataRetrievalPolicy", input)
	return &SetDataRetrievalPolicyFuture{Future: future}
}

func (a *stub) SetVaultAccessPolicy(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error) {
	var output glacier.SetVaultAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-SetVaultAccessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) *SetVaultAccessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-SetVaultAccessPolicy", input)
	return &SetVaultAccessPolicyFuture{Future: future}
}

func (a *stub) SetVaultNotifications(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error) {
	var output glacier.SetVaultNotificationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-SetVaultNotifications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetVaultNotificationsAsync(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) *SetVaultNotificationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-SetVaultNotifications", input)
	return &SetVaultNotificationsFuture{Future: future}
}

func (a *stub) UploadArchive(ctx workflow.Context, input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error) {
	var output glacier.ArchiveCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-UploadArchive", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UploadArchiveAsync(ctx workflow.Context, input *glacier.UploadArchiveInput) *UploadArchiveFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-UploadArchive", input)
	return &UploadArchiveFuture{Future: future}
}

func (a *stub) UploadMultipartPart(ctx workflow.Context, input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error) {
	var output glacier.UploadMultipartPartOutput
	err := workflow.ExecuteActivity(ctx, "aws-glacier-UploadMultipartPart", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UploadMultipartPartAsync(ctx workflow.Context, input *glacier.UploadMultipartPartInput) *UploadMultipartPartFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-UploadMultipartPart", input)
	return &UploadMultipartPartFuture{Future: future}
}

func (a *stub) WaitUntilVaultExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error {
	return workflow.ExecuteActivity(ctx, "aws-glacier-WaitUntilVaultExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilVaultExistsAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-WaitUntilVaultExists", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilVaultNotExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error {
	return workflow.ExecuteActivity(ctx, "aws-glacier-WaitUntilVaultNotExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilVaultNotExistsAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-glacier-WaitUntilVaultNotExists", input)
	return clients.NewVoidFuture(future)
}
