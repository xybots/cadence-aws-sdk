// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package wellarchitectedstub

import (
	"github.com/aws/aws-sdk-go/service/wellarchitected"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateLensesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateLensesFuture) Get(ctx workflow.Context) (*wellarchitected.AssociateLensesOutput, error) {
	var output wellarchitected.AssociateLensesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateMilestoneFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateMilestoneFuture) Get(ctx workflow.Context) (*wellarchitected.CreateMilestoneOutput, error) {
	var output wellarchitected.CreateMilestoneOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateWorkloadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateWorkloadFuture) Get(ctx workflow.Context) (*wellarchitected.CreateWorkloadOutput, error) {
	var output wellarchitected.CreateWorkloadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateWorkloadShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateWorkloadShareFuture) Get(ctx workflow.Context) (*wellarchitected.CreateWorkloadShareOutput, error) {
	var output wellarchitected.CreateWorkloadShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteWorkloadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteWorkloadFuture) Get(ctx workflow.Context) (*wellarchitected.DeleteWorkloadOutput, error) {
	var output wellarchitected.DeleteWorkloadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteWorkloadShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteWorkloadShareFuture) Get(ctx workflow.Context) (*wellarchitected.DeleteWorkloadShareOutput, error) {
	var output wellarchitected.DeleteWorkloadShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateLensesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateLensesFuture) Get(ctx workflow.Context) (*wellarchitected.DisassociateLensesOutput, error) {
	var output wellarchitected.DisassociateLensesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAnswerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAnswerFuture) Get(ctx workflow.Context) (*wellarchitected.GetAnswerOutput, error) {
	var output wellarchitected.GetAnswerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLensReviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLensReviewFuture) Get(ctx workflow.Context) (*wellarchitected.GetLensReviewOutput, error) {
	var output wellarchitected.GetLensReviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLensReviewReportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLensReviewReportFuture) Get(ctx workflow.Context) (*wellarchitected.GetLensReviewReportOutput, error) {
	var output wellarchitected.GetLensReviewReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLensVersionDifferenceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLensVersionDifferenceFuture) Get(ctx workflow.Context) (*wellarchitected.GetLensVersionDifferenceOutput, error) {
	var output wellarchitected.GetLensVersionDifferenceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMilestoneFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMilestoneFuture) Get(ctx workflow.Context) (*wellarchitected.GetMilestoneOutput, error) {
	var output wellarchitected.GetMilestoneOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetWorkloadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetWorkloadFuture) Get(ctx workflow.Context) (*wellarchitected.GetWorkloadOutput, error) {
	var output wellarchitected.GetWorkloadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAnswersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAnswersFuture) Get(ctx workflow.Context) (*wellarchitected.ListAnswersOutput, error) {
	var output wellarchitected.ListAnswersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListLensReviewImprovementsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListLensReviewImprovementsFuture) Get(ctx workflow.Context) (*wellarchitected.ListLensReviewImprovementsOutput, error) {
	var output wellarchitected.ListLensReviewImprovementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListLensReviewsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListLensReviewsFuture) Get(ctx workflow.Context) (*wellarchitected.ListLensReviewsOutput, error) {
	var output wellarchitected.ListLensReviewsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListLensesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListLensesFuture) Get(ctx workflow.Context) (*wellarchitected.ListLensesOutput, error) {
	var output wellarchitected.ListLensesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMilestonesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMilestonesFuture) Get(ctx workflow.Context) (*wellarchitected.ListMilestonesOutput, error) {
	var output wellarchitected.ListMilestonesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListNotificationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListNotificationsFuture) Get(ctx workflow.Context) (*wellarchitected.ListNotificationsOutput, error) {
	var output wellarchitected.ListNotificationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListShareInvitationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListShareInvitationsFuture) Get(ctx workflow.Context) (*wellarchitected.ListShareInvitationsOutput, error) {
	var output wellarchitected.ListShareInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWorkloadSharesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWorkloadSharesFuture) Get(ctx workflow.Context) (*wellarchitected.ListWorkloadSharesOutput, error) {
	var output wellarchitected.ListWorkloadSharesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWorkloadsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWorkloadsFuture) Get(ctx workflow.Context) (*wellarchitected.ListWorkloadsOutput, error) {
	var output wellarchitected.ListWorkloadsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateAnswerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateAnswerFuture) Get(ctx workflow.Context) (*wellarchitected.UpdateAnswerOutput, error) {
	var output wellarchitected.UpdateAnswerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateLensReviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateLensReviewFuture) Get(ctx workflow.Context) (*wellarchitected.UpdateLensReviewOutput, error) {
	var output wellarchitected.UpdateLensReviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateShareInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateShareInvitationFuture) Get(ctx workflow.Context) (*wellarchitected.UpdateShareInvitationOutput, error) {
	var output wellarchitected.UpdateShareInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateWorkloadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateWorkloadFuture) Get(ctx workflow.Context) (*wellarchitected.UpdateWorkloadOutput, error) {
	var output wellarchitected.UpdateWorkloadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateWorkloadShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateWorkloadShareFuture) Get(ctx workflow.Context) (*wellarchitected.UpdateWorkloadShareOutput, error) {
	var output wellarchitected.UpdateWorkloadShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpgradeLensReviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpgradeLensReviewFuture) Get(ctx workflow.Context) (*wellarchitected.UpgradeLensReviewOutput, error) {
	var output wellarchitected.UpgradeLensReviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateLenses(ctx workflow.Context, input *wellarchitected.AssociateLensesInput) (*wellarchitected.AssociateLensesOutput, error) {
	var output wellarchitected.AssociateLensesOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-AssociateLenses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateLensesAsync(ctx workflow.Context, input *wellarchitected.AssociateLensesInput) *AssociateLensesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-AssociateLenses", input)
	return &AssociateLensesFuture{Future: future}
}

func (a *stub) CreateMilestone(ctx workflow.Context, input *wellarchitected.CreateMilestoneInput) (*wellarchitected.CreateMilestoneOutput, error) {
	var output wellarchitected.CreateMilestoneOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateMilestone", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMilestoneAsync(ctx workflow.Context, input *wellarchitected.CreateMilestoneInput) *CreateMilestoneFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateMilestone", input)
	return &CreateMilestoneFuture{Future: future}
}

func (a *stub) CreateWorkload(ctx workflow.Context, input *wellarchitected.CreateWorkloadInput) (*wellarchitected.CreateWorkloadOutput, error) {
	var output wellarchitected.CreateWorkloadOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateWorkload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkloadAsync(ctx workflow.Context, input *wellarchitected.CreateWorkloadInput) *CreateWorkloadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateWorkload", input)
	return &CreateWorkloadFuture{Future: future}
}

func (a *stub) CreateWorkloadShare(ctx workflow.Context, input *wellarchitected.CreateWorkloadShareInput) (*wellarchitected.CreateWorkloadShareOutput, error) {
	var output wellarchitected.CreateWorkloadShareOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateWorkloadShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkloadShareAsync(ctx workflow.Context, input *wellarchitected.CreateWorkloadShareInput) *CreateWorkloadShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-CreateWorkloadShare", input)
	return &CreateWorkloadShareFuture{Future: future}
}

func (a *stub) DeleteWorkload(ctx workflow.Context, input *wellarchitected.DeleteWorkloadInput) (*wellarchitected.DeleteWorkloadOutput, error) {
	var output wellarchitected.DeleteWorkloadOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DeleteWorkload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkloadAsync(ctx workflow.Context, input *wellarchitected.DeleteWorkloadInput) *DeleteWorkloadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DeleteWorkload", input)
	return &DeleteWorkloadFuture{Future: future}
}

func (a *stub) DeleteWorkloadShare(ctx workflow.Context, input *wellarchitected.DeleteWorkloadShareInput) (*wellarchitected.DeleteWorkloadShareOutput, error) {
	var output wellarchitected.DeleteWorkloadShareOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DeleteWorkloadShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkloadShareAsync(ctx workflow.Context, input *wellarchitected.DeleteWorkloadShareInput) *DeleteWorkloadShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DeleteWorkloadShare", input)
	return &DeleteWorkloadShareFuture{Future: future}
}

func (a *stub) DisassociateLenses(ctx workflow.Context, input *wellarchitected.DisassociateLensesInput) (*wellarchitected.DisassociateLensesOutput, error) {
	var output wellarchitected.DisassociateLensesOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DisassociateLenses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateLensesAsync(ctx workflow.Context, input *wellarchitected.DisassociateLensesInput) *DisassociateLensesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-DisassociateLenses", input)
	return &DisassociateLensesFuture{Future: future}
}

func (a *stub) GetAnswer(ctx workflow.Context, input *wellarchitected.GetAnswerInput) (*wellarchitected.GetAnswerOutput, error) {
	var output wellarchitected.GetAnswerOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetAnswer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAnswerAsync(ctx workflow.Context, input *wellarchitected.GetAnswerInput) *GetAnswerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetAnswer", input)
	return &GetAnswerFuture{Future: future}
}

func (a *stub) GetLensReview(ctx workflow.Context, input *wellarchitected.GetLensReviewInput) (*wellarchitected.GetLensReviewOutput, error) {
	var output wellarchitected.GetLensReviewOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensReview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLensReviewAsync(ctx workflow.Context, input *wellarchitected.GetLensReviewInput) *GetLensReviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensReview", input)
	return &GetLensReviewFuture{Future: future}
}

func (a *stub) GetLensReviewReport(ctx workflow.Context, input *wellarchitected.GetLensReviewReportInput) (*wellarchitected.GetLensReviewReportOutput, error) {
	var output wellarchitected.GetLensReviewReportOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensReviewReport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLensReviewReportAsync(ctx workflow.Context, input *wellarchitected.GetLensReviewReportInput) *GetLensReviewReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensReviewReport", input)
	return &GetLensReviewReportFuture{Future: future}
}

func (a *stub) GetLensVersionDifference(ctx workflow.Context, input *wellarchitected.GetLensVersionDifferenceInput) (*wellarchitected.GetLensVersionDifferenceOutput, error) {
	var output wellarchitected.GetLensVersionDifferenceOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensVersionDifference", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLensVersionDifferenceAsync(ctx workflow.Context, input *wellarchitected.GetLensVersionDifferenceInput) *GetLensVersionDifferenceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetLensVersionDifference", input)
	return &GetLensVersionDifferenceFuture{Future: future}
}

func (a *stub) GetMilestone(ctx workflow.Context, input *wellarchitected.GetMilestoneInput) (*wellarchitected.GetMilestoneOutput, error) {
	var output wellarchitected.GetMilestoneOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetMilestone", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMilestoneAsync(ctx workflow.Context, input *wellarchitected.GetMilestoneInput) *GetMilestoneFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetMilestone", input)
	return &GetMilestoneFuture{Future: future}
}

func (a *stub) GetWorkload(ctx workflow.Context, input *wellarchitected.GetWorkloadInput) (*wellarchitected.GetWorkloadOutput, error) {
	var output wellarchitected.GetWorkloadOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetWorkload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetWorkloadAsync(ctx workflow.Context, input *wellarchitected.GetWorkloadInput) *GetWorkloadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-GetWorkload", input)
	return &GetWorkloadFuture{Future: future}
}

func (a *stub) ListAnswers(ctx workflow.Context, input *wellarchitected.ListAnswersInput) (*wellarchitected.ListAnswersOutput, error) {
	var output wellarchitected.ListAnswersOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListAnswers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAnswersAsync(ctx workflow.Context, input *wellarchitected.ListAnswersInput) *ListAnswersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListAnswers", input)
	return &ListAnswersFuture{Future: future}
}

func (a *stub) ListLensReviewImprovements(ctx workflow.Context, input *wellarchitected.ListLensReviewImprovementsInput) (*wellarchitected.ListLensReviewImprovementsOutput, error) {
	var output wellarchitected.ListLensReviewImprovementsOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLensReviewImprovements", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLensReviewImprovementsAsync(ctx workflow.Context, input *wellarchitected.ListLensReviewImprovementsInput) *ListLensReviewImprovementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLensReviewImprovements", input)
	return &ListLensReviewImprovementsFuture{Future: future}
}

func (a *stub) ListLensReviews(ctx workflow.Context, input *wellarchitected.ListLensReviewsInput) (*wellarchitected.ListLensReviewsOutput, error) {
	var output wellarchitected.ListLensReviewsOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLensReviews", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLensReviewsAsync(ctx workflow.Context, input *wellarchitected.ListLensReviewsInput) *ListLensReviewsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLensReviews", input)
	return &ListLensReviewsFuture{Future: future}
}

func (a *stub) ListLenses(ctx workflow.Context, input *wellarchitected.ListLensesInput) (*wellarchitected.ListLensesOutput, error) {
	var output wellarchitected.ListLensesOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLenses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLensesAsync(ctx workflow.Context, input *wellarchitected.ListLensesInput) *ListLensesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListLenses", input)
	return &ListLensesFuture{Future: future}
}

func (a *stub) ListMilestones(ctx workflow.Context, input *wellarchitected.ListMilestonesInput) (*wellarchitected.ListMilestonesOutput, error) {
	var output wellarchitected.ListMilestonesOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListMilestones", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMilestonesAsync(ctx workflow.Context, input *wellarchitected.ListMilestonesInput) *ListMilestonesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListMilestones", input)
	return &ListMilestonesFuture{Future: future}
}

func (a *stub) ListNotifications(ctx workflow.Context, input *wellarchitected.ListNotificationsInput) (*wellarchitected.ListNotificationsOutput, error) {
	var output wellarchitected.ListNotificationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListNotifications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNotificationsAsync(ctx workflow.Context, input *wellarchitected.ListNotificationsInput) *ListNotificationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListNotifications", input)
	return &ListNotificationsFuture{Future: future}
}

func (a *stub) ListShareInvitations(ctx workflow.Context, input *wellarchitected.ListShareInvitationsInput) (*wellarchitected.ListShareInvitationsOutput, error) {
	var output wellarchitected.ListShareInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListShareInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListShareInvitationsAsync(ctx workflow.Context, input *wellarchitected.ListShareInvitationsInput) *ListShareInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListShareInvitations", input)
	return &ListShareInvitationsFuture{Future: future}
}

func (a *stub) ListWorkloadShares(ctx workflow.Context, input *wellarchitected.ListWorkloadSharesInput) (*wellarchitected.ListWorkloadSharesOutput, error) {
	var output wellarchitected.ListWorkloadSharesOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListWorkloadShares", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkloadSharesAsync(ctx workflow.Context, input *wellarchitected.ListWorkloadSharesInput) *ListWorkloadSharesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListWorkloadShares", input)
	return &ListWorkloadSharesFuture{Future: future}
}

func (a *stub) ListWorkloads(ctx workflow.Context, input *wellarchitected.ListWorkloadsInput) (*wellarchitected.ListWorkloadsOutput, error) {
	var output wellarchitected.ListWorkloadsOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListWorkloads", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkloadsAsync(ctx workflow.Context, input *wellarchitected.ListWorkloadsInput) *ListWorkloadsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-ListWorkloads", input)
	return &ListWorkloadsFuture{Future: future}
}

func (a *stub) UpdateAnswer(ctx workflow.Context, input *wellarchitected.UpdateAnswerInput) (*wellarchitected.UpdateAnswerOutput, error) {
	var output wellarchitected.UpdateAnswerOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateAnswer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAnswerAsync(ctx workflow.Context, input *wellarchitected.UpdateAnswerInput) *UpdateAnswerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateAnswer", input)
	return &UpdateAnswerFuture{Future: future}
}

func (a *stub) UpdateLensReview(ctx workflow.Context, input *wellarchitected.UpdateLensReviewInput) (*wellarchitected.UpdateLensReviewOutput, error) {
	var output wellarchitected.UpdateLensReviewOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateLensReview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateLensReviewAsync(ctx workflow.Context, input *wellarchitected.UpdateLensReviewInput) *UpdateLensReviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateLensReview", input)
	return &UpdateLensReviewFuture{Future: future}
}

func (a *stub) UpdateShareInvitation(ctx workflow.Context, input *wellarchitected.UpdateShareInvitationInput) (*wellarchitected.UpdateShareInvitationOutput, error) {
	var output wellarchitected.UpdateShareInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateShareInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateShareInvitationAsync(ctx workflow.Context, input *wellarchitected.UpdateShareInvitationInput) *UpdateShareInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateShareInvitation", input)
	return &UpdateShareInvitationFuture{Future: future}
}

func (a *stub) UpdateWorkload(ctx workflow.Context, input *wellarchitected.UpdateWorkloadInput) (*wellarchitected.UpdateWorkloadOutput, error) {
	var output wellarchitected.UpdateWorkloadOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateWorkload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWorkloadAsync(ctx workflow.Context, input *wellarchitected.UpdateWorkloadInput) *UpdateWorkloadFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateWorkload", input)
	return &UpdateWorkloadFuture{Future: future}
}

func (a *stub) UpdateWorkloadShare(ctx workflow.Context, input *wellarchitected.UpdateWorkloadShareInput) (*wellarchitected.UpdateWorkloadShareOutput, error) {
	var output wellarchitected.UpdateWorkloadShareOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateWorkloadShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWorkloadShareAsync(ctx workflow.Context, input *wellarchitected.UpdateWorkloadShareInput) *UpdateWorkloadShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpdateWorkloadShare", input)
	return &UpdateWorkloadShareFuture{Future: future}
}

func (a *stub) UpgradeLensReview(ctx workflow.Context, input *wellarchitected.UpgradeLensReviewInput) (*wellarchitected.UpgradeLensReviewOutput, error) {
	var output wellarchitected.UpgradeLensReviewOutput
	err := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpgradeLensReview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpgradeLensReviewAsync(ctx workflow.Context, input *wellarchitected.UpgradeLensReviewInput) *UpgradeLensReviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-wellarchitected-UpgradeLensReview", input)
	return &UpgradeLensReviewFuture{Future: future}
}