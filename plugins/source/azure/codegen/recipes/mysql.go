package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
)

func MySQL() []Resource {
	var serverRelations = []resourceDefinition{
		{
			azureStruct:      &mysql.Configuration{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{"server := parent.Item.(mysql.Server)", `resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return errors.WithStack(err)
			}`},
			listHandler:              valueHandler,
			isRelation:               true,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
	}
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"},
				},
			},
			definitions: append([]resourceDefinition{
				{
					azureStruct:  &mysql.Server{},
					listFunction: "List",
					listHandler:  valueHandler,
					relations:    serverRelations,
				},
			}, serverRelations...),
			serviceNameOverride: "MySQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
