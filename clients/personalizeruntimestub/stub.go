// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package personalizeruntimestub

import (
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GetPersonalizedRankingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetPersonalizedRankingFuture) Get(ctx workflow.Context) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	var output personalizeruntime.GetPersonalizedRankingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRecommendationsFuture) Get(ctx workflow.Context) (*personalizeruntime.GetRecommendationsOutput, error) {
	var output personalizeruntime.GetRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPersonalizedRanking(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	var output personalizeruntime.GetPersonalizedRankingOutput
	err := workflow.ExecuteActivity(ctx, "aws-personalizeruntime-GetPersonalizedRanking", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPersonalizedRankingAsync(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) *GetPersonalizedRankingFuture {
	future := workflow.ExecuteActivity(ctx, "aws-personalizeruntime-GetPersonalizedRanking", input)
	return &GetPersonalizedRankingFuture{Future: future}
}

func (a *stub) GetRecommendations(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
	var output personalizeruntime.GetRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-personalizeruntime-GetRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecommendationsAsync(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) *GetRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-personalizeruntime-GetRecommendations", input)
	return &GetRecommendationsFuture{Future: future}
}