package appengine

import (
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_apps",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#Application`,
		Resolver:    fetchApps,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Application{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
