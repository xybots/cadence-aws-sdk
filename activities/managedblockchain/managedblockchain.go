// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client managedblockchainiface.ManagedBlockchainAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := managedblockchain.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (managedblockchainiface.ManagedBlockchainAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return managedblockchain.New(sess), nil
}

func (a *Activities) CreateMember(ctx context.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMemberWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateNetwork(ctx context.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateNetworkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateNode(ctx context.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProposal(ctx context.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProposalWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMember(ctx context.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMemberWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNode(ctx context.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMember(ctx context.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMemberWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNetwork(ctx context.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNetworkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNode(ctx context.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetProposal(ctx context.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetProposalWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInvitations(ctx context.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInvitationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMembers(ctx context.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMembersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListNetworks(ctx context.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListNetworksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListNodes(ctx context.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListNodesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProposalVotes(ctx context.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProposalVotesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProposals(ctx context.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProposalsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *managedblockchain.ListTagsForResourceInput) (*managedblockchain.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectInvitation(ctx context.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectInvitationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *managedblockchain.TagResourceInput) (*managedblockchain.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *managedblockchain.UntagResourceInput) (*managedblockchain.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMember(ctx context.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMemberWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNode(ctx context.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) VoteOnProposal(ctx context.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.VoteOnProposalWithContext(ctx, input)

	return output, internal.EncodeError(err)
}