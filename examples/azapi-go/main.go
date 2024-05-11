package main

import (
	"github.com/dirien/pulumi-azapi/sdk/go/azapi"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		location := pulumi.String("West Europe")
		cfg := config.New(ctx, "")
		cfg.GetSecret("secret")
		resourceGroup, err := core.NewResourceGroup(ctx, "resource-group", &core.ResourceGroupArgs{
			Location: location,
		})
		if err != nil {
			return err
		}

		appServicePlan, err := azapi.NewResource(ctx, "app-service-plan", &azapi.ResourceArgs{
			Type:     pulumi.String("Microsoft.Web/serverfarms@2020-06-01"),
			Name:     pulumi.String("app-service-plan"),
			ParentId: resourceGroup.ID(),

			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"sku": pulumi.Map{
					"name": pulumi.String("F1"),
				},
				"properties": pulumi.Map{
					"reserved": pulumi.Bool(true),
				},
				"kind":     pulumi.String("linux"),
				"location": location,
			}),
			ResponseExportValues: pulumi.StringArray{
				pulumi.String("*"),
			},
		})
		if err != nil {
			return err
		}

		site, err := azapi.NewResource(ctx, "site", &azapi.ResourceArgs{
			Type:         pulumi.String("Microsoft.Web/sites@2021-02-01"),
			ParentId:     resourceGroup.ID(),
			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"properties": pulumi.Map{
					"serverFarmId": appServicePlan.ID(),
					"siteConfig": pulumi.Map{
						"linuxFxVersion": pulumi.String("node|14-lts"),
					},
					"httpsOnly": pulumi.Bool(true),
				},
				"kind":     pulumi.String("app,linux"),
				"location": location,
			}),
			Identity: &azapi.ResourceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			ResponseExportValues: pulumi.StringArray{
				pulumi.String("*"),
			},
		})
		if err != nil {
			return err
		}

		_, err = azapi.NewUpdateResource(ctx, "source-control", &azapi.UpdateResourceArgs{
			Type:         pulumi.String("Microsoft.Web/sites/sourcecontrols@2022-09-01"),
			Name:         pulumi.String("web"),
			ParentId:     site.ID(),
			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"properties": pulumi.Map{
					"repoUrl":             pulumi.String("https://github.com/Azure-Samples/nodejs-docs-hello-world"),
					"branch":              pulumi.String("main"),
					"isManualIntegration": pulumi.Bool(true),
				},
			}),
		}, pulumi.DependsOn([]pulumi.Resource{site}))
		if err != nil {
			return err
		}
		ctx.Export("url", pulumi.Sprintf("https://%s.azurewebsites.net", site.Name))
		return nil
	})
}
