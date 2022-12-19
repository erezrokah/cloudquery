// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"

func Armhealthbot() []*Table {
	tables := []*Table{
		{
			NewFunc:        armhealthbot.NewBotsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HealthBot/healthBots",
			Namespace:      "microsoft.healthbot",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_healthbot)`,
			Pager:          `NewListPager`,
			ResponseStruct: "BotsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhealthbot())
}
