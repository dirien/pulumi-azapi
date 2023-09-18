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

// This resource can perform resource action which gets information from an existing resource.
// It's recommended to use `ResourceAction` data source to perform readonly action, please use `ResourceAction` resource,
// if user wants to perform actions which change a resource's state.
//
// ## Example Usage
//
// Here's an example to use the `ResourceAction` data source to get a provider's permissions.
//
// Here's an example to use the `ResourceAction` data source to perform a provider action.
func LookupResourceAction(ctx *pulumi.Context, args *LookupResourceActionArgs, opts ...pulumi.InvokeOption) (*LookupResourceActionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceActionResult
	err := ctx.Invoke("azapi:index/getResourceAction:getResourceAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceAction.
type LookupResourceActionArgs struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action *string `pulumi:"action"`
	// A JSON object that contains the request body.
	Body *string `pulumi:"body"`
	// Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
	Method *string `pulumi:"method"`
	// The ID of an existing azure source.
	ResourceId *string `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// A collection of values returned by getResourceAction.
type LookupResourceActionResult struct {
	Action *string `pulumi:"action"`
	Body   *string `pulumi:"body"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Method *string `pulumi:"method"`
	// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
	Output               string   `pulumi:"output"`
	ResourceId           *string  `pulumi:"resourceId"`
	ResponseExportValues []string `pulumi:"responseExportValues"`
	Type                 string   `pulumi:"type"`
}

func LookupResourceActionOutput(ctx *pulumi.Context, args LookupResourceActionOutputArgs, opts ...pulumi.InvokeOption) LookupResourceActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceActionResult, error) {
			args := v.(LookupResourceActionArgs)
			r, err := LookupResourceAction(ctx, &args, opts...)
			var s LookupResourceActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceActionResultOutput)
}

// A collection of arguments for invoking getResourceAction.
type LookupResourceActionOutputArgs struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// A JSON object that contains the request body.
	Body pulumi.StringPtrInput `pulumi:"body"`
	// Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
	Method pulumi.StringPtrInput `pulumi:"method"`
	// The ID of an existing azure source.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayInput `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (LookupResourceActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceActionArgs)(nil)).Elem()
}

// A collection of values returned by getResourceAction.
type LookupResourceActionResultOutput struct{ *pulumi.OutputState }

func (LookupResourceActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceActionResult)(nil)).Elem()
}

func (o LookupResourceActionResultOutput) ToLookupResourceActionResultOutput() LookupResourceActionResultOutput {
	return o
}

func (o LookupResourceActionResultOutput) ToLookupResourceActionResultOutputWithContext(ctx context.Context) LookupResourceActionResultOutput {
	return o
}

func (o LookupResourceActionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourceActionResult] {
	return pulumix.Output[LookupResourceActionResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupResourceActionResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceActionResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o LookupResourceActionResultOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceActionResult) *string { return v.Body }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResourceActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceActionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceActionResultOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceActionResult) *string { return v.Method }).(pulumi.StringPtrOutput)
}

// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
func (o LookupResourceActionResultOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceActionResult) string { return v.Output }).(pulumi.StringOutput)
}

func (o LookupResourceActionResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceActionResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupResourceActionResultOutput) ResponseExportValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourceActionResult) []string { return v.ResponseExportValues }).(pulumi.StringArrayOutput)
}

func (o LookupResourceActionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceActionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceActionResultOutput{})
}
