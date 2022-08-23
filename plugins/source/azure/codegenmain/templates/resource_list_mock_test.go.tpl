// Auto generated code - DO NOT EDIT.

package {{.AzurePackageName}}

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func create{{ .AzureService }}{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMock{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeData(&data))

	page := {{ .AzurePackageName }}.New{{ .AzureStructName }}ListResultPage({{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult) ({{ .AzurePackageName }}.{{ .AzureStructName }}ListResult, error) {
		return {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{}, nil
	})

	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range .MockListFunctionArgs }}, {{.}}{{ end }}).Return(page, nil)
	return s
}

func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.AzureMockTestHelper(t, {{ .AzureService }}{{ .AzureSubService }}(), create{{ .AzureService }}{{ .AzureSubService }}Mock, client.TestOptions{})
}
