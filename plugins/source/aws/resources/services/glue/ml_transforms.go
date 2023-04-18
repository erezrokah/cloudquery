package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func MlTransforms() *schema.Table {
	tableName := "aws_glue_ml_transforms"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_MLTransform.html`,
		Resolver:    fetchGlueMlTransforms,
		Transform:   transformers.TransformWithStruct(&types.MLTransform{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueMlTransformArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueMlTransformTags,
			},
			{
				Name:     "schema",
				Type:     schema.TypeJSON,
				Resolver: resolveMlTransformsSchema,
			},
		},

		Relations: []*schema.Table{
			mlTransformTaskRuns(),
		},
	}
}

func fetchGlueMlTransforms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	paginator := glue.NewGetMLTransformsPaginator(svc, &glue.GetMLTransformsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Transforms
	}
	return nil
}
func resolveGlueMlTransformArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.MLTransform)
	return resource.Set(c.Name, mlTransformARN(cl, &r))
}
func resolveGlueMlTransformTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.MLTransform)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(mlTransformARN(cl, &r)),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
func resolveMlTransformsSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]string)
	for _, c := range r.Schema {
		j[*c.Name] = *c.DataType
	}
	return resource.Set(c.Name, j)
}
