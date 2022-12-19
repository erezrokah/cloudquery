// Code generated by codegen; DO NOT EDIT.

package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ActiveReceiptRuleSets() *schema.Table {
	return &schema.Table{
		Name:        "aws_ses_active_receipt_rule_sets",
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html`,
		Resolver:    fetchSesActiveReceiptRuleSets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Metadata.Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Metadata.CreatedTimestamp"),
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
		},
	}
}
