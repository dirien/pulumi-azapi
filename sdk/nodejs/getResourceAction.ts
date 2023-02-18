// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can perform resource action which gets information from an existing resource.
 * It's recommended to use `azapi.ResourceAction` data source to perform readonly action, please use `azapi.ResourceAction` resource,
 * if user wants to perform actions which change a resource's state.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azapi from "@pulumi/azapi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "west europe"});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     skuName: "Basic",
 * });
 * const exampleResourceAction = azapi.getResourceActionOutput({
 *     type: "Microsoft.Automation/automationAccounts@2021-06-22",
 *     resourceId: exampleAccount.id,
 *     action: "listKeys",
 *     responseExportValues: ["*"],
 * });
 * ```
 */
export function getResourceAction(args: GetResourceActionArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceActionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azapi:index/getResourceAction:getResourceAction", {
        "action": args.action,
        "body": args.body,
        "method": args.method,
        "resourceId": args.resourceId,
        "responseExportValues": args.responseExportValues,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceAction.
 */
export interface GetResourceActionArgs {
    /**
     * The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
     */
    action?: string;
    /**
     * A JSON object that contains the request body.
     */
    body?: string;
    /**
     * Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
     */
    method?: string;
    /**
     * The ID of an existing azure source.
     */
    resourceId?: string;
    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `["*"]` will export the full response body.
     * Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     */
    responseExportValues?: string[];
    /**
     * It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `<api-version>` is version of the API used to manage this azure resource.
     */
    type: string;
}

/**
 * A collection of values returned by getResourceAction.
 */
export interface GetResourceActionResult {
    readonly action?: string;
    readonly body?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly method?: string;
    /**
     * The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
     */
    readonly output: string;
    readonly resourceId?: string;
    readonly responseExportValues?: string[];
    readonly type: string;
}
/**
 * This resource can perform resource action which gets information from an existing resource.
 * It's recommended to use `azapi.ResourceAction` data source to perform readonly action, please use `azapi.ResourceAction` resource,
 * if user wants to perform actions which change a resource's state.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azapi from "@pulumi/azapi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "west europe"});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     skuName: "Basic",
 * });
 * const exampleResourceAction = azapi.getResourceActionOutput({
 *     type: "Microsoft.Automation/automationAccounts@2021-06-22",
 *     resourceId: exampleAccount.id,
 *     action: "listKeys",
 *     responseExportValues: ["*"],
 * });
 * ```
 */
export function getResourceActionOutput(args: GetResourceActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceActionResult> {
    return pulumi.output(args).apply((a: any) => getResourceAction(a, opts))
}

/**
 * A collection of arguments for invoking getResourceAction.
 */
export interface GetResourceActionOutputArgs {
    /**
     * The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
     */
    action?: pulumi.Input<string>;
    /**
     * A JSON object that contains the request body.
     */
    body?: pulumi.Input<string>;
    /**
     * Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
     */
    method?: pulumi.Input<string>;
    /**
     * The ID of an existing azure source.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `["*"]` will export the full response body.
     * Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     */
    responseExportValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `<api-version>` is version of the API used to manage this azure resource.
     */
    type: pulumi.Input<string>;
}
