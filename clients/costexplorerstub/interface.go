// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package costexplorerstub

import (
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CreateAnomalyMonitorFuture

	CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CreateAnomalySubscriptionFuture

	CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CreateCostCategoryDefinitionFuture

	DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *DeleteAnomalyMonitorFuture

	DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *DeleteAnomalySubscriptionFuture

	DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *DeleteCostCategoryDefinitionFuture

	DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *DescribeCostCategoryDefinitionFuture

	GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *GetAnomaliesFuture

	GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *GetAnomalyMonitorsFuture

	GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *GetAnomalySubscriptionsFuture

	GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *GetCostAndUsageFuture

	GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *GetCostAndUsageWithResourcesFuture

	GetCostCategories(ctx workflow.Context, input *costexplorer.GetCostCategoriesInput) (*costexplorer.GetCostCategoriesOutput, error)
	GetCostCategoriesAsync(ctx workflow.Context, input *costexplorer.GetCostCategoriesInput) *GetCostCategoriesFuture

	GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *GetCostForecastFuture

	GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *GetDimensionValuesFuture

	GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *GetReservationCoverageFuture

	GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *GetReservationPurchaseRecommendationFuture

	GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *GetReservationUtilizationFuture

	GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *GetRightsizingRecommendationFuture

	GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *GetSavingsPlansCoverageFuture

	GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *GetSavingsPlansPurchaseRecommendationFuture

	GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *GetSavingsPlansUtilizationFuture

	GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *GetSavingsPlansUtilizationDetailsFuture

	GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *GetTagsFuture

	GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *GetUsageForecastFuture

	ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *ListCostCategoryDefinitionsFuture

	ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *ProvideAnomalyFeedbackFuture

	UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *UpdateAnomalyMonitorFuture

	UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *UpdateAnomalySubscriptionFuture

	UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *UpdateCostCategoryDefinitionFuture
}

func NewClient() Client {
	return &stub{}
}