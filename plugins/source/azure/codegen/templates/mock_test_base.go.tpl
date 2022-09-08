// Auto generated code - DO NOT EDIT.

package {{.AzureService | ToLower }}

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
    {{template "imports.go.tpl" .}}
)

func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.MockTestHelper(t, {{ .AzureSubService }}(), create{{ .AzureSubService }}Mock)
}

{{range .MockHelpers}}
{{.}}
{{end}}