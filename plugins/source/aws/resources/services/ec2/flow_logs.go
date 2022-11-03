// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FlowLogs() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_flow_logs",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FlowLog.html",
		Resolver:    fetchEc2FlowLogs,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveFlowLogArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "deliver_cross_account_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliverCrossAccountRole"),
			},
			{
				Name:     "deliver_logs_error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliverLogsErrorMessage"),
			},
			{
				Name:     "deliver_logs_permission_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliverLogsPermissionArn"),
			},
			{
				Name:     "deliver_logs_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliverLogsStatus"),
			},
			{
				Name:     "destination_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DestinationOptions"),
			},
			{
				Name:     "flow_log_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FlowLogId"),
			},
			{
				Name:     "flow_log_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FlowLogStatus"),
			},
			{
				Name:     "log_destination",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogDestination"),
			},
			{
				Name:     "log_destination_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogDestinationType"),
			},
			{
				Name:     "log_format",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogFormat"),
			},
			{
				Name:     "log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupName"),
			},
			{
				Name:     "max_aggregation_interval",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxAggregationInterval"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "traffic_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrafficType"),
			},
		},
	}
}
