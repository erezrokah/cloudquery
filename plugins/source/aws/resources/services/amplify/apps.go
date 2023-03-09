package amplify

import (
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Apps() *schema.Table {
	tableName := "aws_amplify_apps"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amplify/latest/APIReference/API_ListApps.html`,
		Resolver:    fetchApps,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "amplify"),
		Transform:   transformers.TransformWithStruct(&types.App{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
