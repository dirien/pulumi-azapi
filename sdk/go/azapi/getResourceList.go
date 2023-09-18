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

// This resource can list all resources of a specific type under a scope. If the API supports paging, it will automatically fetch all pages and return the full list.
//
// ## Example Usage
func GetResourceList(ctx *pulumi.Context, args *GetResourceListArgs, opts ...pulumi.InvokeOption) (*GetResourceListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetResourceListResult
	err := ctx.Invoke("azapi:index/getResourceList:getResourceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceList.
type GetResourceListArgs struct {
	// The parent resource ID to list resources under. e.g. `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup`.
	ParentId string `pulumi:"parentId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["value"]`, it will set the following json to computed property `output`.
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// A collection of values returned by getResourceList.
type GetResourceListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
	Output               string   `pulumi:"output"`
	ParentId             string   `pulumi:"parentId"`
	ResponseExportValues []string `pulumi:"responseExportValues"`
	Type                 string   `pulumi:"type"`
}

func GetResourceListOutput(ctx *pulumi.Context, args GetResourceListOutputArgs, opts ...pulumi.InvokeOption) GetResourceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResourceListResult, error) {
			args := v.(GetResourceListArgs)
			r, err := GetResourceList(ctx, &args, opts...)
			var s GetResourceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResourceListResultOutput)
}

// A collection of arguments for invoking getResourceList.
type GetResourceListOutputArgs struct {
	// The parent resource ID to list resources under. e.g. `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup`.
	ParentId pulumi.StringInput `pulumi:"parentId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["value"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayInput `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetResourceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceListArgs)(nil)).Elem()
}

// A collection of values returned by getResourceList.
type GetResourceListResultOutput struct{ *pulumi.OutputState }

func (GetResourceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceListResult)(nil)).Elem()
}

func (o GetResourceListResultOutput) ToGetResourceListResultOutput() GetResourceListResultOutput {
	return o
}

func (o GetResourceListResultOutput) ToGetResourceListResultOutputWithContext(ctx context.Context) GetResourceListResultOutput {
	return o
}

func (o GetResourceListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetResourceListResult] {
	return pulumix.Output[GetResourceListResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceListResult) string { return v.Id }).(pulumi.StringOutput)
}

// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
func (o GetResourceListResultOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceListResult) string { return v.Output }).(pulumi.StringOutput)
}

func (o GetResourceListResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceListResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o GetResourceListResultOutput) ResponseExportValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourceListResult) []string { return v.ResponseExportValues }).(pulumi.StringArrayOutput)
}

func (o GetResourceListResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceListResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourceListResultOutput{})
}
