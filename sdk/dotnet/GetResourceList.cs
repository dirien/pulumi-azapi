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
    public static class GetResourceList
    {
        /// <summary>
        /// This resource can list all resources of a specific type under a scope. If the API supports paging, it will automatically fetch all pages and return the full list.
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
        /// data "azapi_resource_list" "listBySubscription" {
        ///   type                   = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// data "azapi_resource_list" "listByResourceGroup" {
        ///   type                   = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// data "azapi_resource_list" "listSubnetsByVnet" {
        ///   type                   = "Microsoft.Network/virtualNetworks/subnets@2021-02-01"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// ```
        /// </summary>
        public static Task<GetResourceListResult> InvokeAsync(GetResourceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceListResult>("azapi:index/getResourceList:getResourceList", args ?? new GetResourceListArgs(), options.WithDefaults());

        /// <summary>
        /// This resource can list all resources of a specific type under a scope. If the API supports paging, it will automatically fetch all pages and return the full list.
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
        /// data "azapi_resource_list" "listBySubscription" {
        ///   type                   = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// data "azapi_resource_list" "listByResourceGroup" {
        ///   type                   = "Microsoft.Automation/automationAccounts@2021-06-22"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// data "azapi_resource_list" "listSubnetsByVnet" {
        ///   type                   = "Microsoft.Network/virtualNetworks/subnets@2021-02-01"
        ///   parent_id              = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1"
        ///   response_export_values = ["*"]
        /// }
        /// 
        /// ```
        /// </summary>
        public static Output<GetResourceListResult> Invoke(GetResourceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceListResult>("azapi:index/getResourceList:getResourceList", args ?? new GetResourceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The parent resource ID to list resources under. e.g. `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup`.
        /// </summary>
        [Input("parentId", required: true)]
        public string ParentId { get; set; } = null!;

        [Input("responseExportValues")]
        private List<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["value"]`, it will set the following json to computed property `output`.
        /// ```
        /// {
        /// "value": [
        /// {
        /// "id": "id1",
        /// "Permissions": "Full"
        /// },
        /// {
        /// "id": "id2",
        /// "Permissions": "Full"
        /// }
        /// ]
        /// }
        /// ```
        /// </summary>
        public List<string> ResponseExportValues
        {
            get => _responseExportValues ?? (_responseExportValues = new List<string>());
            set => _responseExportValues = value;
        }

        /// <summary>
        /// It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
        /// `&lt;api-version&gt;` is version of the API used to manage this azure resource.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetResourceListArgs()
        {
        }
        public static new GetResourceListArgs Empty => new GetResourceListArgs();
    }

    public sealed class GetResourceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The parent resource ID to list resources under. e.g. `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup`.
        /// </summary>
        [Input("parentId", required: true)]
        public Input<string> ParentId { get; set; } = null!;

        [Input("responseExportValues")]
        private InputList<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["value"]`, it will set the following json to computed property `output`.
        /// ```
        /// {
        /// "value": [
        /// {
        /// "id": "id1",
        /// "Permissions": "Full"
        /// },
        /// {
        /// "id": "id2",
        /// "Permissions": "Full"
        /// }
        /// ]
        /// }
        /// ```
        /// </summary>
        public InputList<string> ResponseExportValues
        {
            get => _responseExportValues ?? (_responseExportValues = new InputList<string>());
            set => _responseExportValues = value;
        }

        /// <summary>
        /// It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
        /// `&lt;api-version&gt;` is version of the API used to manage this azure resource.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetResourceListInvokeArgs()
        {
        }
        public static new GetResourceListInvokeArgs Empty => new GetResourceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The output json containing the properties specified in `response_export_values`. Here are some examples to decode json and extract the value.
        /// </summary>
        public readonly string Output;
        public readonly string ParentId;
        public readonly ImmutableArray<string> ResponseExportValues;
        public readonly string Type;

        [OutputConstructor]
        private GetResourceListResult(
            string id,

            string output,

            string parentId,

            ImmutableArray<string> responseExportValues,

            string type)
        {
            Id = id;
            Output = output;
            ParentId = parentId;
            ResponseExportValues = responseExportValues;
            Type = type;
        }
    }
}