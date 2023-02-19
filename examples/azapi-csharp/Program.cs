using Pulumi;
using Azure = Pulumi.Azure;
using Azapi = ediri.Azapi;
using System.Collections.Generic;
using System.Text.Json;


return await Pulumi.Deployment.RunAsync(() =>
{
    var location = "West Europe";
    // Create an Azure Resource Group
    var resourceGroup = new Azure.Core.ResourceGroup("resource-group", new Azure.Core.ResourceGroupArgs
    {
        Location = location
    });

    var appServicePlan = new Azapi.Resource("app-service-plan", new Azapi.ResourceArgs
    {
        Type = "Microsoft.Web/serverfarms@2020-06-01",
        Name = "app-service-plan",
        ParentId = resourceGroup.Id,
        IgnoreCasing = true,
        Body=  JsonSerializer.Serialize(new
        {
            kind = "linux",
            location = location,
            sku = new
            {
                name = "F1",
            },
            properties = new
            {
                reserved = true,
            }
        }),
        ResponseExportValues = new List<string> { 
            "*" 
        }
    });

    var site = new Azapi.Resource("site", new Azapi.ResourceArgs
    {
        Type = "Microsoft.Web/sites@2021-02-01",
        ParentId = resourceGroup.Id,
        IgnoreCasing = true,
        Body=  appServicePlan.Id.Apply(id => JsonSerializer.Serialize(new
        {
            kind = "app,linux",
            location = location,
            properties = new
            {
                serverFarmId = id,
                siteConfig = new
                {
                    linuxFxVersion = "node|14-lts",
                },
                httpsOnly = true,
            }
        })),
        Identity = new Azapi.Inputs.ResourceIdentityArgs
        {
            Type = "SystemAssigned"
        },
        ResponseExportValues = new List<string> { 
            "*"
        }
    });

    new Azapi.UpdateResource("source-control", new Azapi.UpdateResourceArgs
    {
        Type = "Microsoft.Web/sites/sourcecontrols@2022-03-01",
        Name = "web",
        ParentId= site.Id,
        IgnoreCasing = true,
        Body = JsonSerializer.Serialize(new
        {
            properties = new
            {
                repoUrl = "https://github.com/Azure-Samples/nodejs-docs-hello-world",
                branch = "main",
                isManualIntegration = true,
            }
        })
    }, new CustomResourceOptions
    {
        DependsOn = { site }
    }); 

    // Export the primary key of the Storage Account
    return new Dictionary<string, object?>
    {
        ["url"] = Pulumi.Output.Format($"https://{site.Name}.azurewebsites.net")
    };
});