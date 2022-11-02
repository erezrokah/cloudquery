// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancerTlsCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_load_balancer_tls_certificates",
		Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html",
		Resolver:    fetchLightsailLoadBalancerTlsCertificates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "domain_validation_records",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainValidationRecords"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "is_attached",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAttached"),
			},
			{
				Name:     "issued_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("IssuedAt"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyAlgorithm"),
			},
			{
				Name:     "load_balancer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancerName"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotAfter"),
			},
			{
				Name:     "not_before",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotBefore"),
			},
			{
				Name:     "renewal_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RenewalSummary"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "revocation_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevocationReason"),
			},
			{
				Name:     "revoked_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RevokedAt"),
			},
			{
				Name:     "serial",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Serial"),
			},
			{
				Name:     "signature_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SignatureAlgorithm"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subject",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
