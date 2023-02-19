package main

import (
	"github.com/dirien/pulumi-azapi/go/azapi"
	managedidentity "github.com/pulumi/pulumi-azure-native-sdk/managedidentity/v20230131"
	resources "github.com/pulumi/pulumi-azure-native-sdk/resources/v20220901"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		resourceGroup, err := resources.NewResourceGroup(ctx, "test_rg", &resources.ResourceGroupArgs{
			Location: pulumi.String("westeurope"),
		})
		if err != nil {
			return err
		}

		userAssignedIdentity, err := managedidentity.NewUserAssignedIdentity(ctx, "my-identity", &managedidentity.UserAssignedIdentityArgs{
			ResourceGroupName: resourceGroup.Name,
			Location:          resourceGroup.Location,
		})
		if err != nil {
			return err
		}
		userAssignedIdentityMap := userAssignedIdentity.ID().ToIDOutput().ToStringOutput().ApplyT(func(v string) map[string]interface{} {
			m := make(map[string]interface{})
			m[v] = pulumi.ToStringMap(map[string]string{})
			return m
		}).(pulumi.MapOutput)

		registry, err := azapi.NewResource(ctx, "my-resource", &azapi.ResourceArgs{
			Type:     pulumi.String("Microsoft.ContainerRegistry/registries@2022-12-01"),
			Name:     pulumi.String("engindiriregisry"),
			ParentId: resourceGroup.ID(),

			IgnoreCasing: pulumi.Bool(true),
			Body: pulumi.JSONMarshal(pulumi.Map{
				"sku": pulumi.Map{
					"name": pulumi.String("Standard"),
				},
				"properties": pulumi.Map{
					"adminUserEnabled": pulumi.Bool(true),
				},
				"identity": pulumi.Map{
					"type":                   pulumi.String("SystemAssigned, UserAssigned"),
					"userAssignedIdentities": userAssignedIdentityMap,
				},
				"location": resourceGroup.Location,
			}),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
			ResponseExportValues: pulumi.StringArray{
				pulumi.String("properties.loginServer"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("loginServer", pulumi.JSONUnmarshal(registry.Output).AsMapOutput().ApplyT(func(v interface{}) interface{} {
			return v.(map[string]interface{})["properties"].(map[string]interface{})["loginServer"]
		}))
		return nil
	})
}
