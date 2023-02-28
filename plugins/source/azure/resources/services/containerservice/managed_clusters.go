package containerservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ManagedClusters() *schema.Table {
	return &schema.Table{
		Name:        "azure_containerservice_managed_clusters",
		Resolver:    fetchManagedClusters,
		Description: "https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/list?tabs=HTTP#managedcluster",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_containerservice_managed_clusters", client.Namespacemicrosoft_containerservice),
		Transform:   transformers.TransformWithStruct(&armcontainerservice.ManagedCluster{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
		Relations:   []*schema.Table{clusterUpgradeProfiles()},
	}
}

func fetchManagedClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcontainerservice.NewManagedClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
