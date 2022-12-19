// Code generated by codegen; DO NOT EDIT.

package elastictranscoder

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Presets() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_presets",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-presets.html`,
		Resolver:    fetchElastictranscoderPresets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elastictranscoder"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "audio",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Audio"),
			},
			{
				Name:     "container",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Container"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "thumbnails",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Thumbnails"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "video",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Video"),
			},
		},
	}
}
