// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"

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
	client elasticsearchserviceiface.ElasticsearchServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elasticsearchservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elasticsearchserviceiface.ElasticsearchServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return elasticsearchservice.New(sess), nil
}

func (a *Activities) AcceptInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AcceptInboundCrossClusterSearchConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AddTags(ctx context.Context, input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociatePackage(ctx context.Context, input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociatePackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelElasticsearchServiceSoftwareUpdate(ctx context.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelElasticsearchServiceSoftwareUpdateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateElasticsearchDomain(ctx context.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateElasticsearchDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateOutboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateOutboundCrossClusterSearchConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePackage(ctx context.Context, input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteElasticsearchDomain(ctx context.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteElasticsearchDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteElasticsearchServiceRole(ctx context.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteElasticsearchServiceRoleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInboundCrossClusterSearchConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteOutboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteOutboundCrossClusterSearchConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePackage(ctx context.Context, input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeElasticsearchDomain(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeElasticsearchDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeElasticsearchDomainConfig(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeElasticsearchDomainConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeElasticsearchDomains(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeElasticsearchDomainsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeElasticsearchInstanceTypeLimits(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeElasticsearchInstanceTypeLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInboundCrossClusterSearchConnections(ctx context.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInboundCrossClusterSearchConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeOutboundCrossClusterSearchConnections(ctx context.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeOutboundCrossClusterSearchConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePackages(ctx context.Context, input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePackagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReservedElasticsearchInstanceOfferings(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReservedElasticsearchInstanceOfferingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReservedElasticsearchInstances(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReservedElasticsearchInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DissociatePackage(ctx context.Context, input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DissociatePackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCompatibleElasticsearchVersions(ctx context.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCompatibleElasticsearchVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPackageVersionHistory(ctx context.Context, input *elasticsearchservice.GetPackageVersionHistoryInput) (*elasticsearchservice.GetPackageVersionHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPackageVersionHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetUpgradeHistory(ctx context.Context, input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetUpgradeHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetUpgradeStatus(ctx context.Context, input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetUpgradeStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDomainNames(ctx context.Context, input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDomainNamesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDomainsForPackage(ctx context.Context, input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDomainsForPackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListElasticsearchInstanceTypes(ctx context.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListElasticsearchInstanceTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListElasticsearchVersions(ctx context.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListElasticsearchVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPackagesForDomain(ctx context.Context, input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPackagesForDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTags(ctx context.Context, input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PurchaseReservedElasticsearchInstanceOffering(ctx context.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PurchaseReservedElasticsearchInstanceOfferingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectInboundCrossClusterSearchConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTags(ctx context.Context, input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartElasticsearchServiceSoftwareUpdate(ctx context.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartElasticsearchServiceSoftwareUpdateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateElasticsearchDomainConfig(ctx context.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateElasticsearchDomainConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePackage(ctx context.Context, input *elasticsearchservice.UpdatePackageInput) (*elasticsearchservice.UpdatePackageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePackageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpgradeElasticsearchDomain(ctx context.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpgradeElasticsearchDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}