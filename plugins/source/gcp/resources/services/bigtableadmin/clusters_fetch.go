package bigtableadmin

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/bigtable"
)

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	instance := parent.Item.(*bigtable.InstanceInfo)
	gcpClient, err := bigtable.NewInstanceAdminClient(ctx, c.ProjectId, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.Clusters(ctx, instance.Name)
	if err != nil {
		return err
	}

	res <- resp
	return nil
}
