// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource can manage any Azure resource manager resource.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Azure resource can be imported using the `resource id`, e.g.
 *
 * ```sh
 * $ pulumi import azapi:index/resource:Resource example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
 * ```
 *
 * It also supports specifying API version by using the `resource id` with `api-version` as a query parameter, e.g.
 *
 * ```sh
 * $ pulumi import azapi:index/resource:Resource example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1?api-version=2021-07-01
 * ```
 */
export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceState, opts?: pulumi.CustomResourceOptions): Resource {
        return new Resource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azapi:index/resource:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resource.__pulumiType;
    }

    /**
     * A JSON object that contains the request body used to create and update azure resource.
     */
    public readonly body!: pulumi.Output<string | undefined>;
    /**
     * A `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.ResourceIdentity>;
    /**
     * A list of properties that should be ignored when comparing the `body` with its current state.
     */
    public readonly ignoreBodyChanges!: pulumi.Output<string[] | undefined>;
    /**
     * Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
     */
    public readonly ignoreCasing!: pulumi.Output<boolean | undefined>;
    /**
     * Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
     */
    public readonly ignoreMissingProperty!: pulumi.Output<boolean | undefined>;
    /**
     * The Azure Region where the azure resource should exist.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
     */
    public readonly locks!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the name of the azure resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
     * ```
     * // it will output "registry1.azurecr.io"
     * output "loginServer" {
     * value = jsondecode(azapi_resource.example.output).properties.loginServer
     * }
     */
    public /*out*/ readonly output!: pulumi.Output<string>;
    /**
     * The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
     * - resource group scope: `parentId` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
     * - management group scope: `parentId` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
     * - extension scope: `parentId` should be the ID of the resource you're adding the extension to.
     * - subscription scope: `parentId` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
     * - tenant scope: `parentId` should be `/`
     *
     * For child level resources, the `parentId` should be the ID of its parent resource, for example, subnet resource's `parentId` is the ID of the vnet.
     *
     * For type `Microsoft.Resources/resourceGroups`, the `parentId` could be omitted, it defaults to subscription ID specified in provider or the default subscription(You could check the default subscription by azure cli command: `az account show`).
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * Whether to remove special characters in resource name. Defaults to `false`.
     *
     * @deprecated It will not work in the next minor release and will be removed in the next major release. Please specify the `name` field and remove the special characters in the `name` field instead.
     */
    public readonly removingSpecialChars!: pulumi.Output<boolean | undefined>;
    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `["*"]` will export the full response body.
     * Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
     * ```
     * {
     * "properties" : {
     * "loginServer" : "registry1.azurecr.io"
     * "policies" : {
     * "quarantinePolicy" = {
     * "status" = "disabled"
     * }
     * }
     * }
     * }
     * ```
     */
    public readonly responseExportValues!: pulumi.Output<string[] | undefined>;
    /**
     * Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
     */
    public readonly schemaValidationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags which should be assigned to the azure resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `<api-version>` is version of the API used to manage this azure resource.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceArgs | ResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceState | undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["ignoreBodyChanges"] = state ? state.ignoreBodyChanges : undefined;
            resourceInputs["ignoreCasing"] = state ? state.ignoreCasing : undefined;
            resourceInputs["ignoreMissingProperty"] = state ? state.ignoreMissingProperty : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["locks"] = state ? state.locks : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["output"] = state ? state.output : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["removingSpecialChars"] = state ? state.removingSpecialChars : undefined;
            resourceInputs["responseExportValues"] = state ? state.responseExportValues : undefined;
            resourceInputs["schemaValidationEnabled"] = state ? state.schemaValidationEnabled : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ResourceArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["ignoreBodyChanges"] = args ? args.ignoreBodyChanges : undefined;
            resourceInputs["ignoreCasing"] = args ? args.ignoreCasing : undefined;
            resourceInputs["ignoreMissingProperty"] = args ? args.ignoreMissingProperty : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["locks"] = args ? args.locks : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["removingSpecialChars"] = args ? args.removingSpecialChars : undefined;
            resourceInputs["responseExportValues"] = args ? args.responseExportValues : undefined;
            resourceInputs["schemaValidationEnabled"] = args ? args.schemaValidationEnabled : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["output"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resource resources.
 */
export interface ResourceState {
    /**
     * A JSON object that contains the request body used to create and update azure resource.
     */
    body?: pulumi.Input<string>;
    /**
     * A `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.ResourceIdentity>;
    /**
     * A list of properties that should be ignored when comparing the `body` with its current state.
     */
    ignoreBodyChanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
     */
    ignoreCasing?: pulumi.Input<boolean>;
    /**
     * Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
     */
    ignoreMissingProperty?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the azure resource should exist.
     */
    location?: pulumi.Input<string>;
    /**
     * A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
     */
    locks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the azure resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
     * ```
     * // it will output "registry1.azurecr.io"
     * output "loginServer" {
     * value = jsondecode(azapi_resource.example.output).properties.loginServer
     * }
     */
    output?: pulumi.Input<string>;
    /**
     * The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
     * - resource group scope: `parentId` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
     * - management group scope: `parentId` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
     * - extension scope: `parentId` should be the ID of the resource you're adding the extension to.
     * - subscription scope: `parentId` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
     * - tenant scope: `parentId` should be `/`
     *
     * For child level resources, the `parentId` should be the ID of its parent resource, for example, subnet resource's `parentId` is the ID of the vnet.
     *
     * For type `Microsoft.Resources/resourceGroups`, the `parentId` could be omitted, it defaults to subscription ID specified in provider or the default subscription(You could check the default subscription by azure cli command: `az account show`).
     */
    parentId?: pulumi.Input<string>;
    /**
     * Whether to remove special characters in resource name. Defaults to `false`.
     *
     * @deprecated It will not work in the next minor release and will be removed in the next major release. Please specify the `name` field and remove the special characters in the `name` field instead.
     */
    removingSpecialChars?: pulumi.Input<boolean>;
    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `["*"]` will export the full response body.
     * Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
     * ```
     * {
     * "properties" : {
     * "loginServer" : "registry1.azurecr.io"
     * "policies" : {
     * "quarantinePolicy" = {
     * "status" = "disabled"
     * }
     * }
     * }
     * }
     * ```
     */
    responseExportValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
     */
    schemaValidationEnabled?: pulumi.Input<boolean>;
    /**
     * A mapping of tags which should be assigned to the azure resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `<api-version>` is version of the API used to manage this azure resource.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    /**
     * A JSON object that contains the request body used to create and update azure resource.
     */
    body?: pulumi.Input<string>;
    /**
     * A `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.ResourceIdentity>;
    /**
     * A list of properties that should be ignored when comparing the `body` with its current state.
     */
    ignoreBodyChanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
     */
    ignoreCasing?: pulumi.Input<boolean>;
    /**
     * Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
     */
    ignoreMissingProperty?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the azure resource should exist.
     */
    location?: pulumi.Input<string>;
    /**
     * A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
     */
    locks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the azure resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the azure resource in which this resource is created. Changing this forces a new resource to be created. It supports different kinds of deployment scope for **top level** resources: 
     * - resource group scope: `parentId` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
     * - management group scope: `parentId` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
     * - extension scope: `parentId` should be the ID of the resource you're adding the extension to.
     * - subscription scope: `parentId` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
     * - tenant scope: `parentId` should be `/`
     *
     * For child level resources, the `parentId` should be the ID of its parent resource, for example, subnet resource's `parentId` is the ID of the vnet.
     *
     * For type `Microsoft.Resources/resourceGroups`, the `parentId` could be omitted, it defaults to subscription ID specified in provider or the default subscription(You could check the default subscription by azure cli command: `az account show`).
     */
    parentId?: pulumi.Input<string>;
    /**
     * Whether to remove special characters in resource name. Defaults to `false`.
     *
     * @deprecated It will not work in the next minor release and will be removed in the next major release. Please specify the `name` field and remove the special characters in the `name` field instead.
     */
    removingSpecialChars?: pulumi.Input<boolean>;
    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `["*"]` will export the full response body.
     * Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
     * ```
     * {
     * "properties" : {
     * "loginServer" : "registry1.azurecr.io"
     * "policies" : {
     * "quarantinePolicy" = {
     * "status" = "disabled"
     * }
     * }
     * }
     * }
     * ```
     */
    responseExportValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
     */
    schemaValidationEnabled?: pulumi.Input<boolean>;
    /**
     * A mapping of tags which should be assigned to the azure resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `<api-version>` is version of the API used to manage this azure resource.
     */
    type: pulumi.Input<string>;
}
