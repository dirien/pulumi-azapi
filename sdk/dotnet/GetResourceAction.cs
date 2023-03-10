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
    public static class GetResourceAction
    {
        /// <summary>
        /// This resource can perform resource action which gets information from an existing resource.
        /// It's recommended to use `azapi.ResourceAction` data source to perform readonly action, please use `azapi.ResourceAction` resource,
        /// if user wants to perform actions which change a resource's state.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Azapi = Pulumi.Azapi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new()
        ///     {
        ///         Location = "west europe",
        ///     });
        /// 
        ///     var exampleAccount = new Azure.Automation.Account("exampleAccount", new()
        ///     {
        ///         ResourceGroupName = exampleResourceGroup.Name,
        ///         Location = exampleResourceGroup.Location,
        ///         SkuName = "Basic",
        ///     });
        /// 
        ///     var exampleResourceAction = Azapi.GetResourceAction.Invoke(new()
        ///     {
        ///         Type = "Microsoft.Automation/automationAccounts@2021-06-22",
        ///         ResourceId = exampleAccount.Id,
        ///         Action = "listKeys",
        ///         ResponseExportValues = new[]
        ///         {
        ///             "*",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourceActionResult> InvokeAsync(GetResourceActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceActionResult>("azapi:index/getResourceAction:getResourceAction", args ?? new GetResourceActionArgs(), options.WithDefaults());

        /// <summary>
        /// This resource can perform resource action which gets information from an existing resource.
        /// It's recommended to use `azapi.ResourceAction` data source to perform readonly action, please use `azapi.ResourceAction` resource,
        /// if user wants to perform actions which change a resource's state.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Azapi = Pulumi.Azapi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new()
        ///     {
        ///         Location = "west europe",
        ///     });
        /// 
        ///     var exampleAccount = new Azure.Automation.Account("exampleAccount", new()
        ///     {
        ///         ResourceGroupName = exampleResourceGroup.Name,
        ///         Location = exampleResourceGroup.Location,
        ///         SkuName = "Basic",
        ///     });
        /// 
        ///     var exampleResourceAction = Azapi.GetResourceAction.Invoke(new()
        ///     {
        ///         Type = "Microsoft.Automation/automationAccounts@2021-06-22",
        ///         ResourceId = exampleAccount.Id,
        ///         Action = "listKeys",
        ///         ResponseExportValues = new[]
        ///         {
        ///             "*",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResourceActionResult> Invoke(GetResourceActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceActionResult>("azapi:index/getResourceAction:getResourceAction", args ?? new GetResourceActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceActionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
        /// </summary>
        [Input("action")]
        public string? Action { get; set; }

        /// <summary>
        /// A JSON object that contains the request body.
        /// </summary>
        [Input("body")]
        public string? Body { get; set; }

        /// <summary>
        /// Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
        /// </summary>
        [Input("method")]
        public string? Method { get; set; }

        /// <summary>
        /// The ID of an existing azure source.
        /// </summary>
        [Input("resourceId")]
        public string? ResourceId { get; set; }

        [Input("responseExportValues")]
        private List<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
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

        public GetResourceActionArgs()
        {
        }
        public static new GetResourceActionArgs Empty => new GetResourceActionArgs();
    }

    public sealed class GetResourceActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A JSON object that contains the request body.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// The ID of an existing azure source.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        [Input("responseExportValues")]
        private InputList<string>? _responseExportValues;

        /// <summary>
        /// A list of path that needs to be exported from response body.
        /// Setting it to `["*"]` will export the full response body.
        /// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
        /// ```csharp
        /// using System.Collections.Generic;
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

        public GetResourceActionInvokeArgs()
        {
        }
        public static new GetResourceActionInvokeArgs Empty => new GetResourceActionInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceActionResult
    {
        public readonly string? Action;
        public readonly string? Body;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Method;
        /// <summary>
        /// The output json containing the properties specified in `response_export_values`. Here are some examples to decode json and extract the value.
        /// </summary>
        public readonly string Output;
        public readonly string? ResourceId;
        public readonly ImmutableArray<string> ResponseExportValues;
        public readonly string Type;

        [OutputConstructor]
        private GetResourceActionResult(
            string? action,

            string? body,

            string id,

            string? method,

            string output,

            string? resourceId,

            ImmutableArray<string> responseExportValues,

            string type)
        {
            Action = action;
            Body = body;
            Id = id;
            Method = method;
            Output = output;
            ResourceId = resourceId;
            ResponseExportValues = responseExportValues;
            Type = type;
        }
    }
}
