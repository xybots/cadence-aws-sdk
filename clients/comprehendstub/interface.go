// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package comprehendstub

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchDetectDominantLanguage(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error)
	BatchDetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) *BatchDetectDominantLanguageFuture

	BatchDetectEntities(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error)
	BatchDetectEntitiesAsync(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) *BatchDetectEntitiesFuture

	BatchDetectKeyPhrases(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error)
	BatchDetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) *BatchDetectKeyPhrasesFuture

	BatchDetectSentiment(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error)
	BatchDetectSentimentAsync(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) *BatchDetectSentimentFuture

	BatchDetectSyntax(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error)
	BatchDetectSyntaxAsync(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) *BatchDetectSyntaxFuture

	ClassifyDocument(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error)
	ClassifyDocumentAsync(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) *ClassifyDocumentFuture

	CreateDocumentClassifier(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error)
	CreateDocumentClassifierAsync(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) *CreateDocumentClassifierFuture

	CreateEndpoint(ctx workflow.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *comprehend.CreateEndpointInput) *CreateEndpointFuture

	CreateEntityRecognizer(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error)
	CreateEntityRecognizerAsync(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) *CreateEntityRecognizerFuture

	DeleteDocumentClassifier(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error)
	DeleteDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) *DeleteDocumentClassifierFuture

	DeleteEndpoint(ctx workflow.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *comprehend.DeleteEndpointInput) *DeleteEndpointFuture

	DeleteEntityRecognizer(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error)
	DeleteEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) *DeleteEntityRecognizerFuture

	DescribeDocumentClassificationJob(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error)
	DescribeDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) *DescribeDocumentClassificationJobFuture

	DescribeDocumentClassifier(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error)
	DescribeDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) *DescribeDocumentClassifierFuture

	DescribeDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error)
	DescribeDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) *DescribeDominantLanguageDetectionJobFuture

	DescribeEndpoint(ctx workflow.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error)
	DescribeEndpointAsync(ctx workflow.Context, input *comprehend.DescribeEndpointInput) *DescribeEndpointFuture

	DescribeEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error)
	DescribeEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) *DescribeEntitiesDetectionJobFuture

	DescribeEntityRecognizer(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error)
	DescribeEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) *DescribeEntityRecognizerFuture

	DescribeEventsDetectionJob(ctx workflow.Context, input *comprehend.DescribeEventsDetectionJobInput) (*comprehend.DescribeEventsDetectionJobOutput, error)
	DescribeEventsDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeEventsDetectionJobInput) *DescribeEventsDetectionJobFuture

	DescribeKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error)
	DescribeKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) *DescribeKeyPhrasesDetectionJobFuture

	DescribePiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput) (*comprehend.DescribePiiEntitiesDetectionJobOutput, error)
	DescribePiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput) *DescribePiiEntitiesDetectionJobFuture

	DescribeSentimentDetectionJob(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error)
	DescribeSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) *DescribeSentimentDetectionJobFuture

	DescribeTopicsDetectionJob(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error)
	DescribeTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) *DescribeTopicsDetectionJobFuture

	DetectDominantLanguage(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error)
	DetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) *DetectDominantLanguageFuture

	DetectEntities(ctx workflow.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error)
	DetectEntitiesAsync(ctx workflow.Context, input *comprehend.DetectEntitiesInput) *DetectEntitiesFuture

	DetectKeyPhrases(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error)
	DetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) *DetectKeyPhrasesFuture

	DetectPiiEntities(ctx workflow.Context, input *comprehend.DetectPiiEntitiesInput) (*comprehend.DetectPiiEntitiesOutput, error)
	DetectPiiEntitiesAsync(ctx workflow.Context, input *comprehend.DetectPiiEntitiesInput) *DetectPiiEntitiesFuture

	DetectSentiment(ctx workflow.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error)
	DetectSentimentAsync(ctx workflow.Context, input *comprehend.DetectSentimentInput) *DetectSentimentFuture

	DetectSyntax(ctx workflow.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error)
	DetectSyntaxAsync(ctx workflow.Context, input *comprehend.DetectSyntaxInput) *DetectSyntaxFuture

	ListDocumentClassificationJobs(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error)
	ListDocumentClassificationJobsAsync(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) *ListDocumentClassificationJobsFuture

	ListDocumentClassifiers(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error)
	ListDocumentClassifiersAsync(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) *ListDocumentClassifiersFuture

	ListDominantLanguageDetectionJobs(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error)
	ListDominantLanguageDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) *ListDominantLanguageDetectionJobsFuture

	ListEndpoints(ctx workflow.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *comprehend.ListEndpointsInput) *ListEndpointsFuture

	ListEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error)
	ListEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) *ListEntitiesDetectionJobsFuture

	ListEntityRecognizers(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error)
	ListEntityRecognizersAsync(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) *ListEntityRecognizersFuture

	ListEventsDetectionJobs(ctx workflow.Context, input *comprehend.ListEventsDetectionJobsInput) (*comprehend.ListEventsDetectionJobsOutput, error)
	ListEventsDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListEventsDetectionJobsInput) *ListEventsDetectionJobsFuture

	ListKeyPhrasesDetectionJobs(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error)
	ListKeyPhrasesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) *ListKeyPhrasesDetectionJobsFuture

	ListPiiEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput) (*comprehend.ListPiiEntitiesDetectionJobsOutput, error)
	ListPiiEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput) *ListPiiEntitiesDetectionJobsFuture

	ListSentimentDetectionJobs(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error)
	ListSentimentDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) *ListSentimentDetectionJobsFuture

	ListTagsForResource(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTopicsDetectionJobs(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error)
	ListTopicsDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) *ListTopicsDetectionJobsFuture

	StartDocumentClassificationJob(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error)
	StartDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) *StartDocumentClassificationJobFuture

	StartDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error)
	StartDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) *StartDominantLanguageDetectionJobFuture

	StartEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error)
	StartEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) *StartEntitiesDetectionJobFuture

	StartEventsDetectionJob(ctx workflow.Context, input *comprehend.StartEventsDetectionJobInput) (*comprehend.StartEventsDetectionJobOutput, error)
	StartEventsDetectionJobAsync(ctx workflow.Context, input *comprehend.StartEventsDetectionJobInput) *StartEventsDetectionJobFuture

	StartKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error)
	StartKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) *StartKeyPhrasesDetectionJobFuture

	StartPiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartPiiEntitiesDetectionJobInput) (*comprehend.StartPiiEntitiesDetectionJobOutput, error)
	StartPiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartPiiEntitiesDetectionJobInput) *StartPiiEntitiesDetectionJobFuture

	StartSentimentDetectionJob(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error)
	StartSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) *StartSentimentDetectionJobFuture

	StartTopicsDetectionJob(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error)
	StartTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) *StartTopicsDetectionJobFuture

	StopDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error)
	StopDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) *StopDominantLanguageDetectionJobFuture

	StopEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error)
	StopEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) *StopEntitiesDetectionJobFuture

	StopEventsDetectionJob(ctx workflow.Context, input *comprehend.StopEventsDetectionJobInput) (*comprehend.StopEventsDetectionJobOutput, error)
	StopEventsDetectionJobAsync(ctx workflow.Context, input *comprehend.StopEventsDetectionJobInput) *StopEventsDetectionJobFuture

	StopKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error)
	StopKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) *StopKeyPhrasesDetectionJobFuture

	StopPiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopPiiEntitiesDetectionJobInput) (*comprehend.StopPiiEntitiesDetectionJobOutput, error)
	StopPiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopPiiEntitiesDetectionJobInput) *StopPiiEntitiesDetectionJobFuture

	StopSentimentDetectionJob(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error)
	StopSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) *StopSentimentDetectionJobFuture

	StopTrainingDocumentClassifier(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error)
	StopTrainingDocumentClassifierAsync(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) *StopTrainingDocumentClassifierFuture

	StopTrainingEntityRecognizer(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error)
	StopTrainingEntityRecognizerAsync(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) *StopTrainingEntityRecognizerFuture

	TagResource(ctx workflow.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *comprehend.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *comprehend.UntagResourceInput) *UntagResourceFuture

	UpdateEndpoint(ctx workflow.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error)
	UpdateEndpointAsync(ctx workflow.Context, input *comprehend.UpdateEndpointInput) *UpdateEndpointFuture
}

func NewClient() Client {
	return &stub{}
}