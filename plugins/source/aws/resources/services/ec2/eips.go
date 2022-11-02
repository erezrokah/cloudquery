// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Eips() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_eips",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html",
		Resolver:    fetchEc2Eips,
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
				Name:     "allocation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllocationId"),
			},
			{
				Name:     "association_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationId"),
			},
			{
				Name:     "carrier_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CarrierIp"),
			},
			{
				Name:     "customer_owned_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerOwnedIp"),
			},
			{
				Name:     "customer_owned_ipv4_pool",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerOwnedIpv4Pool"),
			},
			{
				Name:     "domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Domain"),
			},
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceId"),
			},
			{
				Name:     "network_border_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkBorderGroup"),
			},
			{
				Name:     "network_interface_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkInterfaceId"),
			},
			{
				Name:     "network_interface_owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkInterfaceOwnerId"),
			},
			{
				Name:     "private_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateIpAddress"),
			},
			{
				Name:     "public_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIp"),
			},
			{
				Name:     "public_ipv4_pool",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIpv4Pool"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
