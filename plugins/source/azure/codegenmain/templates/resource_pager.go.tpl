{{template "base.go.tpl" .}}

func fetch{{.AzureService}}{{.AzureSubService}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{ .AzureService }}.{{ .AzureSubService }}
	pager := svc.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}