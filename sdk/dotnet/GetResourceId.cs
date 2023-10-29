// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Azapi
{
    public static class GetResourceId
    {
        /// <summary>
        /// This resource can parse an Azure resource ID into its separate fields.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// terraform {
        ///   required_providers {
        ///     azapi = {
        ///       source = "Azure/azapi"
        ///     }
        ///   }
        /// }
        /// 
        /// provider "azapi" {
        /// }
        /// 
        /// data "azapi_resource_id" "account" {
        ///   type        = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   resource_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Automation/automationAccounts/automationAccount1"
        /// }
        /// 
        /// output "account_name" {
        ///   value = data.azapi_resource_id.account.name
        /// }
        /// 
        /// output "account_resource_group" {
        ///   value = data.azapi_resource_id.account.resource_group_name
        /// }
        /// 
        /// output "account_subscription" {
        ///   value = data.azapi_resource_id.account.subscription_id
        /// }
        /// 
        /// output "account_parent_id" {
        ///   value = data.azapi_resource_id.account.parent_id
        /// }
        /// 
        /// data "azapi_resource_id" "vnet" {
        ///   type      = "Microsoft.Network/virtualNetworks@2021-02-01"
        ///   parent_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1"
        ///   name      = "vnet1"
        /// }
        /// 
        /// output "vnet_id" {
        ///   value = data.azapi_resource_id.vnet.id
        /// }
        /// 
        /// output "vnet_resource_group" {
        ///   value = data.azapi_resource_id.vnet.resource_group_name
        /// }
        /// 
        /// output "vnet_subscription" {
        ///   value = data.azapi_resource_id.vnet.subscription_id
        /// }
        /// ```
        /// </summary>
        public static Task<GetResourceIdResult> InvokeAsync(GetResourceIdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceIdResult>("azapi:index/getResourceId:getResourceId", args ?? new GetResourceIdArgs(), options.WithDefaults());

        /// <summary>
        /// This resource can parse an Azure resource ID into its separate fields.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// terraform {
        ///   required_providers {
        ///     azapi = {
        ///       source = "Azure/azapi"
        ///     }
        ///   }
        /// }
        /// 
        /// provider "azapi" {
        /// }
        /// 
        /// data "azapi_resource_id" "account" {
        ///   type        = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   resource_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Automation/automationAccounts/automationAccount1"
        /// }
        /// 
        /// output "account_name" {
        ///   value = data.azapi_resource_id.account.name
        /// }
        /// 
        /// output "account_resource_group" {
        ///   value = data.azapi_resource_id.account.resource_group_name
        /// }
        /// 
        /// output "account_subscription" {
        ///   value = data.azapi_resource_id.account.subscription_id
        /// }
        /// 
        /// output "account_parent_id" {
        ///   value = data.azapi_resource_id.account.parent_id
        /// }
        /// 
        /// data "azapi_resource_id" "vnet" {
        ///   type      = "Microsoft.Network/virtualNetworks@2021-02-01"
        ///   parent_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1"
        ///   name      = "vnet1"
        /// }
        /// 
        /// output "vnet_id" {
        ///   value = data.azapi_resource_id.vnet.id
        /// }
        /// 
        /// output "vnet_resource_group" {
        ///   value = data.azapi_resource_id.vnet.resource_group_name
        /// }
        /// 
        /// output "vnet_subscription" {
        ///   value = data.azapi_resource_id.vnet.subscription_id
        /// }
        /// ```
        /// </summary>
        public static Output<GetResourceIdResult> Invoke(GetResourceIdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceIdResult>("azapi:index/getResourceId:getResourceId", args ?? new GetResourceIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the azure resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:
        /// - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
        /// - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
        /// - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
        /// - subscription scope: `parent_id` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
        /// - tenant scope: `parent_id` should be `/`
        /// 
        /// For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
        /// </summary>
        [Input("parentId")]
        public string? ParentId { get; set; }

        /// <summary>
        /// The ID of an existing azure source.
        /// 
        /// &gt; **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
        /// </summary>
        [Input("resourceId")]
        public string? ResourceId { get; set; }

        /// <summary>
        /// It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
        /// `&lt;api-version&gt;` is version of the API used to manage this azure resource.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetResourceIdArgs()
        {
        }
        public static new GetResourceIdArgs Empty => new GetResourceIdArgs();
    }

    public sealed class GetResourceIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the azure resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:
        /// - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
        /// - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
        /// - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
        /// - subscription scope: `parent_id` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
        /// - tenant scope: `parent_id` should be `/`
        /// 
        /// For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// The ID of an existing azure source.
        /// 
        /// &gt; **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
        /// `&lt;api-version&gt;` is version of the API used to manage this azure resource.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetResourceIdInvokeArgs()
        {
        }
        public static new GetResourceIdInvokeArgs Empty => new GetResourceIdInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceIdResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the azure resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the azure resource in which this resource is created.
        /// </summary>
        public readonly string ParentId;
        /// <summary>
        /// The map of the resource ID parts, where the key is the part name and the value is the part value. e.g. `virtualNetworks=myVnet`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parts;
        /// <summary>
        /// The azure resource provider namespace of the azure resource.
        /// </summary>
        public readonly string ProviderNamespace;
        /// <summary>
        /// The resource group name of the azure resource.
        /// </summary>
        public readonly string ResourceGroupName;
        public readonly string ResourceId;
        /// <summary>
        /// The subscription ID of the azure resource.
        /// </summary>
        public readonly string SubscriptionId;
        public readonly string Type;

        [OutputConstructor]
        private GetResourceIdResult(
            string id,

            string name,

            string parentId,

            ImmutableDictionary<string, string> parts,

            string providerNamespace,

            string resourceGroupName,

            string resourceId,

            string subscriptionId,

            string type)
        {
            Id = id;
            Name = name;
            ParentId = parentId;
            Parts = parts;
            ProviderNamespace = providerNamespace;
            ResourceGroupName = resourceGroupName;
            ResourceId = resourceId;
            SubscriptionId = subscriptionId;
            Type = type;
        }
    }
}