// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"

func Armmysqlflexibleservers() []*Table {
	tables := []*Table{
		{
			NewFunc:        armmysqlflexibleservers.NewServersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/flexibleServers",
			Namespace:      "microsoft.dbformysql",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_dbformysql)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ServersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmysqlflexibleservers())
}
