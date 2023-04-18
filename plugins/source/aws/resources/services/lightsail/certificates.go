package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Certificates() *schema.Table {
	tableName := "aws_lightsail_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html`,
		Resolver:    fetchLightsailCertificates,
		Transform:   transformers.TransformWithStruct(&types.Certificate{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchLightsailCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := lightsail.GetCertificatesInput{
		IncludeCertificateDetails: true,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetCertificates(ctx, &input)
	if err != nil {
		return err
	}
	for _, cer := range response.Certificates {
		res <- cer.CertificateDetail
	}
	return nil
}
