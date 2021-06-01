// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package eventbridgestub

import (
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ActivateEventSource(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error)
	ActivateEventSourceAsync(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) *ActivateEventSourceFuture

	CancelReplay(ctx workflow.Context, input *eventbridge.CancelReplayInput) (*eventbridge.CancelReplayOutput, error)
	CancelReplayAsync(ctx workflow.Context, input *eventbridge.CancelReplayInput) *CancelReplayFuture

	CreateArchive(ctx workflow.Context, input *eventbridge.CreateArchiveInput) (*eventbridge.CreateArchiveOutput, error)
	CreateArchiveAsync(ctx workflow.Context, input *eventbridge.CreateArchiveInput) *CreateArchiveFuture

	CreateEventBus(ctx workflow.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error)
	CreateEventBusAsync(ctx workflow.Context, input *eventbridge.CreateEventBusInput) *CreateEventBusFuture

	CreatePartnerEventSource(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error)
	CreatePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) *CreatePartnerEventSourceFuture

	DeactivateEventSource(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error)
	DeactivateEventSourceAsync(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) *DeactivateEventSourceFuture

	DeleteArchive(ctx workflow.Context, input *eventbridge.DeleteArchiveInput) (*eventbridge.DeleteArchiveOutput, error)
	DeleteArchiveAsync(ctx workflow.Context, input *eventbridge.DeleteArchiveInput) *DeleteArchiveFuture

	DeleteEventBus(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error)
	DeleteEventBusAsync(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) *DeleteEventBusFuture

	DeletePartnerEventSource(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error)
	DeletePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) *DeletePartnerEventSourceFuture

	DeleteRule(ctx workflow.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *eventbridge.DeleteRuleInput) *DeleteRuleFuture

	DescribeArchive(ctx workflow.Context, input *eventbridge.DescribeArchiveInput) (*eventbridge.DescribeArchiveOutput, error)
	DescribeArchiveAsync(ctx workflow.Context, input *eventbridge.DescribeArchiveInput) *DescribeArchiveFuture

	DescribeEventBus(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error)
	DescribeEventBusAsync(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) *DescribeEventBusFuture

	DescribeEventSource(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error)
	DescribeEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) *DescribeEventSourceFuture

	DescribePartnerEventSource(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error)
	DescribePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) *DescribePartnerEventSourceFuture

	DescribeReplay(ctx workflow.Context, input *eventbridge.DescribeReplayInput) (*eventbridge.DescribeReplayOutput, error)
	DescribeReplayAsync(ctx workflow.Context, input *eventbridge.DescribeReplayInput) *DescribeReplayFuture

	DescribeRule(ctx workflow.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error)
	DescribeRuleAsync(ctx workflow.Context, input *eventbridge.DescribeRuleInput) *DescribeRuleFuture

	DisableRule(ctx workflow.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error)
	DisableRuleAsync(ctx workflow.Context, input *eventbridge.DisableRuleInput) *DisableRuleFuture

	EnableRule(ctx workflow.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error)
	EnableRuleAsync(ctx workflow.Context, input *eventbridge.EnableRuleInput) *EnableRuleFuture

	ListArchives(ctx workflow.Context, input *eventbridge.ListArchivesInput) (*eventbridge.ListArchivesOutput, error)
	ListArchivesAsync(ctx workflow.Context, input *eventbridge.ListArchivesInput) *ListArchivesFuture

	ListEventBuses(ctx workflow.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error)
	ListEventBusesAsync(ctx workflow.Context, input *eventbridge.ListEventBusesInput) *ListEventBusesFuture

	ListEventSources(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error)
	ListEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) *ListEventSourcesFuture

	ListPartnerEventSourceAccounts(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) *ListPartnerEventSourceAccountsFuture

	ListPartnerEventSources(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error)
	ListPartnerEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) *ListPartnerEventSourcesFuture

	ListReplays(ctx workflow.Context, input *eventbridge.ListReplaysInput) (*eventbridge.ListReplaysOutput, error)
	ListReplaysAsync(ctx workflow.Context, input *eventbridge.ListReplaysInput) *ListReplaysFuture

	ListRuleNamesByTarget(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error)
	ListRuleNamesByTargetAsync(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) *ListRuleNamesByTargetFuture

	ListRules(ctx workflow.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *eventbridge.ListRulesInput) *ListRulesFuture

	ListTagsForResource(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTargetsByRule(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error)
	ListTargetsByRuleAsync(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) *ListTargetsByRuleFuture

	PutEvents(ctx workflow.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error)
	PutEventsAsync(ctx workflow.Context, input *eventbridge.PutEventsInput) *PutEventsFuture

	PutPartnerEvents(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error)
	PutPartnerEventsAsync(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) *PutPartnerEventsFuture

	PutPermission(ctx workflow.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error)
	PutPermissionAsync(ctx workflow.Context, input *eventbridge.PutPermissionInput) *PutPermissionFuture

	PutRule(ctx workflow.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error)
	PutRuleAsync(ctx workflow.Context, input *eventbridge.PutRuleInput) *PutRuleFuture

	PutTargets(ctx workflow.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error)
	PutTargetsAsync(ctx workflow.Context, input *eventbridge.PutTargetsInput) *PutTargetsFuture

	RemovePermission(ctx workflow.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *eventbridge.RemovePermissionInput) *RemovePermissionFuture

	RemoveTargets(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error)
	RemoveTargetsAsync(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) *RemoveTargetsFuture

	StartReplay(ctx workflow.Context, input *eventbridge.StartReplayInput) (*eventbridge.StartReplayOutput, error)
	StartReplayAsync(ctx workflow.Context, input *eventbridge.StartReplayInput) *StartReplayFuture

	TagResource(ctx workflow.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *eventbridge.TagResourceInput) *TagResourceFuture

	TestEventPattern(ctx workflow.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error)
	TestEventPatternAsync(ctx workflow.Context, input *eventbridge.TestEventPatternInput) *TestEventPatternFuture

	UntagResource(ctx workflow.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *eventbridge.UntagResourceInput) *UntagResourceFuture

	UpdateArchive(ctx workflow.Context, input *eventbridge.UpdateArchiveInput) (*eventbridge.UpdateArchiveOutput, error)
	UpdateArchiveAsync(ctx workflow.Context, input *eventbridge.UpdateArchiveInput) *UpdateArchiveFuture
}

func NewClient() Client {
	return &stub{}
}
