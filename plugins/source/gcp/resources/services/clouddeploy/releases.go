// Code generated by codegen; DO NOT EDIT.

package clouddeploy

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/deploy/apiv1/deploypb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/deploy/apiv1"
)

func Releases() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouddeploy_releases",
		Description: `https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases#Release`,
		Resolver:    fetchReleases,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouddeploy.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uid"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "abandoned",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Abandoned"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "render_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("RenderStartTime"),
			},
			{
				Name:     "render_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("RenderEndTime"),
			},
			{
				Name:     "skaffold_config_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SkaffoldConfigUri"),
			},
			{
				Name:     "skaffold_config_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SkaffoldConfigPath"),
			},
			{
				Name:     "build_artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildArtifacts"),
			},
			{
				Name:     "delivery_pipeline_snapshot",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPipelineSnapshot"),
			},
			{
				Name:     "target_snapshots",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetSnapshots"),
			},
			{
				Name:     "render_state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("RenderState"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "skaffold_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SkaffoldVersion"),
			},
			{
				Name:     "target_artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetArtifacts"),
			},
			{
				Name:     "target_renders",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetRenders"),
			},
		},

		Relations: []*schema.Table{
			Rollouts(),
		},
	}
}

func fetchReleases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListReleasesRequest{
		Parent: parent.Item.(*pb.DeliveryPipeline).Name,
	}
	gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListReleases(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
