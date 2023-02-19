"""An Azure Python Pulumi program"""

import pulumi
from pulumi_azure import core
from ediri_azapi import Resource, ResourceArgs, ResourceIdentityArgs, UpdateResource, UpdateResourceArgs

location = 'West Europe'

resource_group = core.ResourceGroup('resource-group', location=location)

app_service_plan = Resource('app-service-plan', ResourceArgs(
    type="Microsoft.Web/serverfarms@2020-06-01",
    name="app-service-plan",
    parent_id=resource_group.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "sku": {
            "name": "F1",
        },
        "properties": {
            "reserved": True,
        },
        "kind": "linux",
        "location": location,
    }),
    response_export_values=[
        "*"
    ]
))

site = Resource('site', ResourceArgs(
    type="Microsoft.Web/sites@2021-02-01",
    parent_id=resource_group.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "properties": {
            "serverFarmId": app_service_plan.id,
            "siteConfig": {
                "linuxFxVersion": "node|14-lts",
            },
            "httpsOnly": True,
        },
        "kind": "app,linux",
        "location": location,
    }),
    identity=ResourceIdentityArgs(
        type="SystemAssigned"
    ),
    response_export_values=[
        "*"
    ]
))

UpdateResource('source-control', UpdateResourceArgs(
    type="Microsoft.Web/sites/sourcecontrols@2022-03-01",
    name="web",
    parent_id=site.id,
    ignore_casing=True,
    body=pulumi.Output.json_dumps({
        "properties": {
            "repoUrl": "https://github.com/Azure-Samples/nodejs-docs-hello-world",
            "branch": "main",
            "isManualIntegration": True,
        },
    }),
), depends_on=[site])

pulumi.export('url', site.name.apply(
    lambda name: f'https://{name}.azurewebsites.net'))
