// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-azapi/sdk/go/azapi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can parse an Azure resource ID into its separate fields.
//
// ## Example Usage
func GetResourceId(ctx *pulumi.Context, args *GetResourceIdArgs, opts ...pulumi.InvokeOption) (*GetResourceIdResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetResourceIdResult
	err := ctx.Invoke("azapi:index/getResourceId:getResourceId", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceId.
type GetResourceIdArgs struct {
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
	ParentId *string `pulumi:"parentId"`
	// The ID of an existing azure source.
	//
	// > **Note:** Configuring `name` and `parentId` is an alternative way to configure `resourceId`.
	ResourceId *string `pulumi:"resourceId"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// A collection of values returned by getResourceId.
type GetResourceIdResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the azure resource.
	Name string `pulumi:"name"`
	// The ID of the azure resource in which this resource is created.
	ParentId string `pulumi:"parentId"`
	// The map of the resource ID parts, where the key is the part name and the value is the part value. e.g. `virtualNetworks=myVnet`.
	Parts map[string]string `pulumi:"parts"`
	// The azure resource provider namespace of the azure resource.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource group name of the azure resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceId        string `pulumi:"resourceId"`
	// The subscription ID of the azure resource.
	SubscriptionId string `pulumi:"subscriptionId"`
	Type           string `pulumi:"type"`
}

func GetResourceIdOutput(ctx *pulumi.Context, args GetResourceIdOutputArgs, opts ...pulumi.InvokeOption) GetResourceIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResourceIdResult, error) {
			args := v.(GetResourceIdArgs)
			r, err := GetResourceId(ctx, &args, opts...)
			var s GetResourceIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResourceIdResultOutput)
}

// A collection of arguments for invoking getResourceId.
type GetResourceIdOutputArgs struct {
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
	ParentId pulumi.StringPtrInput `pulumi:"parentId"`
	// The ID of an existing azure source.
	//
	// > **Note:** Configuring `name` and `parentId` is an alternative way to configure `resourceId`.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetResourceIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceIdArgs)(nil)).Elem()
}

// A collection of values returned by getResourceId.
type GetResourceIdResultOutput struct{ *pulumi.OutputState }

func (GetResourceIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceIdResult)(nil)).Elem()
}

func (o GetResourceIdResultOutput) ToGetResourceIdResultOutput() GetResourceIdResultOutput {
	return o
}

func (o GetResourceIdResultOutput) ToGetResourceIdResultOutputWithContext(ctx context.Context) GetResourceIdResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourceIdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the azure resource.
func (o GetResourceIdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the azure resource in which this resource is created.
func (o GetResourceIdResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.ParentId }).(pulumi.StringOutput)
}

// The map of the resource ID parts, where the key is the part name and the value is the part value. e.g. `virtualNetworks=myVnet`.
func (o GetResourceIdResultOutput) Parts() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetResourceIdResult) map[string]string { return v.Parts }).(pulumi.StringMapOutput)
}

// The azure resource provider namespace of the azure resource.
func (o GetResourceIdResultOutput) ProviderNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.ProviderNamespace }).(pulumi.StringOutput)
}

// The resource group name of the azure resource.
func (o GetResourceIdResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o GetResourceIdResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The subscription ID of the azure resource.
func (o GetResourceIdResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o GetResourceIdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourceIdResultOutput{})
}
