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
    /// <summary>
    /// This resource can manage a subset of any existing Azure resource manager resource's properties.
    /// 
    /// &gt; **Note** This resource is used to add or modify properties on an existing resource.
    /// When delete `azapi.UpdateResource`, no operation will be performed, and these properties will stay unchanged.
    /// If you want to restore the modified properties to some values, you must apply the restored properties before deleting.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Azapi = ediri.Azapi;
    /// using Azurerm = Pulumi.Azurerm;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleazurerm_resource_group = new Azurerm.Index.Azurerm_resource_group("exampleazurerm_resource_group", new()
    ///     {
    ///         Name = "example-rg",
    ///         Location = "west europe",
    ///     });
    /// 
    ///     var exampleazurerm_public_ip = new Azurerm.Index.Azurerm_public_ip("exampleazurerm_public_ip", new()
    ///     {
    ///         Name = "example-ip",
    ///         Location = exampleazurerm_resource_group.Location,
    ///         ResourceGroupName = exampleazurerm_resource_group.Name,
    ///         AllocationMethod = "Static",
    ///     });
    /// 
    ///     var exampleazurerm_lb = new Azurerm.Index.Azurerm_lb("exampleazurerm_lb", new()
    ///     {
    ///         Name = "example-lb",
    ///         Location = exampleazurerm_resource_group.Location,
    ///         ResourceGroupName = exampleazurerm_resource_group.Name,
    ///         FrontendIpConfiguration = new[]
    ///         {
    ///             
    ///             {
    ///                 { "name", "PublicIPAddress" },
    ///                 { "publicIpAddressId", exampleazurerm_public_ip.Id },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleazurerm_lb_nat_rule = new Azurerm.Index.Azurerm_lb_nat_rule("exampleazurerm_lb_nat_rule", new()
    ///     {
    ///         ResourceGroupName = exampleazurerm_resource_group.Name,
    ///         LoadbalancerId = exampleazurerm_lb.Id,
    ///         Name = "RDPAccess",
    ///         Protocol = "Tcp",
    ///         FrontendPort = 3389,
    ///         BackendPort = 3389,
    ///         FrontendIpConfigurationName = "PublicIPAddress",
    ///     });
    /// 
    ///     var exampleUpdateResource = new Azapi.UpdateResource("exampleUpdateResource", new()
    ///     {
    ///         Type = "Microsoft.Network/loadBalancers@2021-03-01",
    ///         ResourceId = exampleazurerm_lb.Id,
    ///         Body = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["properties"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["inboundNatRules"] = new[]
    ///                 {
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["properties"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["idleTimeoutInMinutes"] = 15,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleazurerm_lb_nat_rule,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AzapiResourceType("azapi:index/updateResource:UpdateResource")]
    public partial class UpdateResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A JSON object that contains the request body used to add on an existing azure resource.
        /// </summary>
        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        /// <summary>
        /// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
        /// </summary>
        [Output("ignoreCasing")]
        public Output<bool?> IgnoreCasing { get; private set; } = null!;

        /// <summary>
        /// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
        /// </summary>
        [Output("ignoreMissingProperty")]
        public Output<bool?> IgnoreMissingProperty { get; private set; } = null!;

        /// <summary>
        /// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
        /// </summary>
        [Output("locks")]
        public Output<ImmutableArray<string>> Locks { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the azure resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The output json containing the properties specified in `response_export_values`. Here're some examples to decode json and extract the value.
        /// </summary>
        [Output("output")]
        public Output<string> Output { get; private set; } = null!;

        /// <summary>
        /// The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
        /// - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
        /// - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
        /// - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
        /// - subscription scope: `parent_id` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
        /// - tenant scope: `parent_id` should be `/`
        /// 
        /// For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
        /// </summary>
        [Output("parentId")]
        public Output<string> ParentId { get; private set; } = null!;

        /// <summary>
        /// The ID of an existing azure source. Changing this forces a new azure resource to be created.
        /// 
        /// &gt; **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        [Output("responseExportValues")]
        public Output<ImmutableArray<string>> ResponseExportValues { get; private set; } = null!;

        /// <summary>
        /// It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
        /// `&lt;api-version&gt;` is version of the API used to manage this azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a UpdateResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdateResource(string name, UpdateResourceArgs args, CustomResourceOptions? options = null)
            : base("azapi:index/updateResource:UpdateResource", name, args ?? new UpdateResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdateResource(string name, Input<string> id, UpdateResourceState? state = null, CustomResourceOptions? options = null)
            : base("azapi:index/updateResource:UpdateResource", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-azapi",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UpdateResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdateResource Get(string name, Input<string> id, UpdateResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new UpdateResource(name, id, state, options);
        }
    }

    public sealed class UpdateResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A JSON object that contains the request body used to add on an existing azure resource.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
        /// </summary>
        [Input("ignoreCasing")]
        public Input<bool>? IgnoreCasing { get; set; }

        /// <summary>
        /// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
        /// </summary>
        [Input("ignoreMissingProperty")]
        public Input<bool>? IgnoreMissingProperty { get; set; }

        [Input("locks")]
        private InputList<string>? _locks;

        /// <summary>
        /// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
        /// </summary>
        public InputList<string> Locks
        {
            get => _locks ?? (_locks = new InputList<string>());
            set => _locks = value;
        }

        /// <summary>
        /// Specifies the name of the azure resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
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
        /// The ID of an existing azure source. Changing this forces a new azure resource to be created.
        /// 
        /// &gt; **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        [Input("responseExportValues")]
        private InputList<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
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

        public UpdateResourceArgs()
        {
        }
        public static new UpdateResourceArgs Empty => new UpdateResourceArgs();
    }

    public sealed class UpdateResourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A JSON object that contains the request body used to add on an existing azure resource.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
        /// </summary>
        [Input("ignoreCasing")]
        public Input<bool>? IgnoreCasing { get; set; }

        /// <summary>
        /// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
        /// </summary>
        [Input("ignoreMissingProperty")]
        public Input<bool>? IgnoreMissingProperty { get; set; }

        [Input("locks")]
        private InputList<string>? _locks;

        /// <summary>
        /// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
        /// </summary>
        public InputList<string> Locks
        {
            get => _locks ?? (_locks = new InputList<string>());
            set => _locks = value;
        }

        /// <summary>
        /// Specifies the name of the azure resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The output json containing the properties specified in `response_export_values`. Here're some examples to decode json and extract the value.
        /// </summary>
        [Input("output")]
        public Input<string>? Output { get; set; }

        /// <summary>
        /// The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
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
        /// The ID of an existing azure source. Changing this forces a new azure resource to be created.
        /// 
        /// &gt; **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        [Input("responseExportValues")]
        private InputList<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
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
        [Input("type")]
        public Input<string>? Type { get; set; }

        public UpdateResourceState()
        {
        }
        public static new UpdateResourceState Empty => new UpdateResourceState();
    }
}
