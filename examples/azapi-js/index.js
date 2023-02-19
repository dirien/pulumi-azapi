const pulumi = require("@pulumi/pulumi");
const azure = require("@pulumi/azure");
const azapi = require("@ediri/azapi");

const location = "West Europe";

const resourceGroup = new azure.core.ResourceGroup("resource-group", {
    location: location,
});

const appServicePlan = new azapi.Resource("app-service-plan", {
    type: "Microsoft.Web/serverfarms@2020-06-01",
    name: "app-service-plan",
    parentId: resourceGroup.id,
    ignoreCasing: true,
    body: pulumi.jsonStringify(
        {
            sku: {
                name: "F1",
            },
            properties: {
                reserved: true,
            },
            kind: "linux",
            location: location,
        }
    ),
    responseExportValues: [
        "*"
    ]
})

const site = new azapi.Resource("site", {
    type: "Microsoft.Web/sites@2021-02-01",
    parentId: resourceGroup.id,
    ignoreCasing: true,
    body: pulumi.jsonStringify(
        {
            properties: {
                serverFarmId: appServicePlan.id,
                siteConfig: {
                    linuxFxVersion: "node|14-lts",
                },
                httpsOnly: true,
            },
            kind: "app,linux",
            location: location,
        }
    ),
    identity: {
        type: "SystemAssigned"
    },
    responseExportValues: [
        "*"
    ]
})

new azapi.UpdateResource("source-control", {
    type: "Microsoft.Web/sites/sourcecontrols@2022-03-01",
    name: "web",
    parentId: site.id,
    ignoreCasing: true,
    body: pulumi.jsonStringify(
        {
            properties: {
                repoUrl: "https://github.com/Azure-Samples/nodejs-docs-hello-world",
                branch: "main",
                isManualIntegration: true,
            }
        })
}, {
    dependsOn: [
        site
    ]
})

exports.url = pulumi.interpolate`https://${site.name}.azurewebsites.net`;