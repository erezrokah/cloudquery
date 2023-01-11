package peering

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServiceProviders() *schema.Table {
	return &schema.Table{
		Name:      "azure_peering_service_providers",
		Resolver:  fetchServiceProviders,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_peering_service_providers", client.Namespacemicrosoft_peering),
		Transform: transformers.TransformWithStruct(&armpeering.ServiceProvider{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
