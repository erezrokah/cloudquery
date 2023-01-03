package client

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/cloudquery/plugin-sdk/testdata"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func TestPgPlugin(t *testing.T) {
	p := destination.NewPlugin("postgresql", "development", New)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
		destination.PluginTestSuiteTests{})
}

func TestPgPluginPrimaryKeyRename(t *testing.T) {
	tableName := fmt.Sprintf("cq_test_pk_rename_%d", rand.Intn(100))
	tableWithStalePk := testdata.TestTable(tableName)
	// We simulate that a primary column was renamed
	stalePK := schema.Column{Name: "stale_pk", Type: schema.TypeUUID, CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true}}
	tableWithStalePk.Columns = append(tableWithStalePk.Columns, stalePK)
	p := destination.NewPlugin("postgresql", "development", New)
	ctx := context.Background()

	spec := Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      LogLevelTrace,
	}

	// Init the plugin so we can call migrate
	if err := p.Init(ctx, zerolog.Logger{}, specs.Destination{Name: "stale_pk", Spec: spec}); err != nil {
		t.Fatal(err)
	}

	// Call migrate so the `stale_pk` column is created
	if err := p.Migrate(ctx, []*schema.Table{tableWithStalePk}); err != nil {
		t.Fatal(err)
	}

	// Migrate the table again without the `stale_pk` column
	table := testdata.TestTable(tableName)
	if err := p.Migrate(ctx, []*schema.Table{table}); err != nil {
		t.Fatal(err)
	}

	// Write some data
	resources := []schema.DestinationResource{{TableName: table.Name, Data: testdata.GenTestData(table)}}
	ch := make(chan schema.DestinationResource, len(resources))
	for _, resource := range resources {
		ch <- resource
	}
	close(ch)
	err := p.Write(ctx, schema.Tables{table}, "cq_test_pk_rename", time.Now(), ch)
	require.NoError(t, err)
}
