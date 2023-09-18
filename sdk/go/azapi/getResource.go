// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azapi/sdk/go/azapi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can access any existing Azure resource manager resource.
//
// ## Example Usage
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceResult
	err := ctx.Invoke("azapi:index/getResource:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResource.
type LookupResourceArgs struct {
	// An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
	Identity *GetResourceIdentity `pulumi:"identity"`
	// Specifies the name of the azure resource.
	Name *string `pulumi:"name"`
	// The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:
	// - resource group scope: `parentId` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
	// - management group scope: `parentId` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
	// - extension scope: `parentId` should be the ID of the resource you're adding the extension to.
	// - subscription scope: `parentId` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
	// - tenant scope: `parentId` should be `/`
	//
	// For child level resources, the `parentId` should be the ID of its parent resource, for example, subnet resource's `parentId` is the ID of the vnet.
	//
	// For type `Microsoft.Resources/resourceGroups`, the `parentId` could be omitted, it defaults to subscription ID specified in provider or the default subscription(You could check the default subscription by azure cli command: `az account show`).
	ParentId *string `pulumi:"parentId"`
	// The ID of an existing azure source.
	//
	// > **Note:** Configuring `name` and `parentId` is an alternative way to configure `resourceId`.
	ResourceId *string `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// A collection of values returned by getResource.
type LookupResourceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
	Identity GetResourceIdentity `pulumi:"identity"`
	// The Azure Region where the azure resource should exist.
	Location string  `pulumi:"location"`
	Name     *string `pulumi:"name"`
	// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
	Output               string   `pulumi:"output"`
	ParentId             string   `pulumi:"parentId"`
	ResourceId           *string  `pulumi:"resourceId"`
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// A mapping of tags which should be assigned to the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
	Type string `pulumi:"type"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			var s LookupResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceResultOutput)
}

// A collection of arguments for invoking getResource.
type LookupResourceOutputArgs struct {
	// An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
	Identity GetResourceIdentityPtrInput `pulumi:"identity"`
	// Specifies the name of the azure resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:
	// - resource group scope: `parentId` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
	// - management group scope: `parentId` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
	// - extension scope: `parentId` should be the ID of the resource you're adding the extension to.
	// - subscription scope: `parentId` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
	// - tenant scope: `parentId` should be `/`
	//
	// For child level resources, the `parentId` should be the ID of its parent resource, for example, subnet resource's `parentId` is the ID of the vnet.
	//
	// For type `Microsoft.Resources/resourceGroups`, the `parentId` could be omitted, it defaults to subscription ID specified in provider or the default subscription(You could check the default subscription by azure cli command: `az account show`).
	ParentId pulumi.StringPtrInput `pulumi:"parentId"`
	// The ID of an existing azure source.
	//
	// > **Note:** Configuring `name` and `parentId` is an alternative way to configure `resourceId`.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayInput `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}

// A collection of values returned by getResource.
type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourceResult] {
	return pulumix.Output[LookupResourceResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
func (o LookupResourceResultOutput) Identity() GetResourceIdentityOutput {
	return o.ApplyT(func(v LookupResourceResult) GetResourceIdentity { return v.Identity }).(GetResourceIdentityOutput)
}

// The Azure Region where the azure resource should exist.
func (o LookupResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
func (o LookupResourceResultOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Output }).(pulumi.StringOutput)
}

func (o LookupResourceResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o LookupResourceResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) ResponseExportValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourceResult) []string { return v.ResponseExportValues }).(pulumi.StringArrayOutput)
}

// A mapping of tags which should be assigned to the azure resource.
func (o LookupResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
func (o LookupResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}
