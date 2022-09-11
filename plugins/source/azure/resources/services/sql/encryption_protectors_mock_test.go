// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func createEncryptionProtectorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLEncryptionProtectorsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			EncryptionProtectors: mockClient,
		},
	}

	data := sql.EncryptionProtector{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewEncryptionProtectorListResultPage(sql.EncryptionProtectorListResult{Value: &[]sql.EncryptionProtector{data}}, func(ctx context.Context, result sql.EncryptionProtectorListResult) (sql.EncryptionProtectorListResult, error) {
		return sql.EncryptionProtectorListResult{}, nil
	})

	mockClient.EXPECT().Get(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
