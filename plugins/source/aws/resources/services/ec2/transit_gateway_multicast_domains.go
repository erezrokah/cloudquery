// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayMulticastDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_transit_gateway_multicast_domains",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomain.html",
		Resolver:    fetchEc2TransitGatewayMulticastDomains,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "transit_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayId"),
			},
			{
				Name:     "transit_gateway_multicast_domain_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayMulticastDomainArn"),
			},
			{
				Name:     "transit_gateway_multicast_domain_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayMulticastDomainId"),
			},
		},
	}
}
