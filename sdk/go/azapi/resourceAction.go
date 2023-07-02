// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can perform any Azure resource manager resource action.
// It's recommended to use `ResourceAction` resource to perform actions which change a resource's state, please use `ResourceAction` data source,
// if user wants to perform readonly action.
//
// > **Note** When delete `ResourceAction`, no operation will be performed.
//
// ## Example Usage
type ResourceAction struct {
	pulumi.CustomResourceState

	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// A JSON object that contains the request body.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
	Locks pulumi.StringArrayOutput `pulumi:"locks"`
	// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
	Method pulumi.StringPtrOutput `pulumi:"method"`
	// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
	Output pulumi.StringOutput `pulumi:"output"`
	// The ID of an existing azure source.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayOutput `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceAction registers a new resource with the given unique name, arguments, and options.
func NewResourceAction(ctx *pulumi.Context,
	name string, args *ResourceActionArgs, opts ...pulumi.ResourceOption) (*ResourceAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ResourceAction
	err := ctx.RegisterResource("azapi:index/resourceAction:ResourceAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceAction gets an existing ResourceAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceActionState, opts ...pulumi.ResourceOption) (*ResourceAction, error) {
	var resource ResourceAction
	err := ctx.ReadResource("azapi:index/resourceAction:ResourceAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceAction resources.
type resourceActionState struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action *string `pulumi:"action"`
	// A JSON object that contains the request body.
	Body *string `pulumi:"body"`
	// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
	Locks []string `pulumi:"locks"`
	// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
	Method *string `pulumi:"method"`
	// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
	Output *string `pulumi:"output"`
	// The ID of an existing azure source.
	ResourceId *string `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type *string `pulumi:"type"`
}

type ResourceActionState struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action pulumi.StringPtrInput
	// A JSON object that contains the request body.
	Body pulumi.StringPtrInput
	// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
	Locks pulumi.StringArrayInput
	// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
	Method pulumi.StringPtrInput
	// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
	Output pulumi.StringPtrInput
	// The ID of an existing azure source.
	ResourceId pulumi.StringPtrInput
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayInput
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringPtrInput
}

func (ResourceActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceActionState)(nil)).Elem()
}

type resourceActionArgs struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action *string `pulumi:"action"`
	// A JSON object that contains the request body.
	Body *string `pulumi:"body"`
	// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
	Locks []string `pulumi:"locks"`
	// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
	Method *string `pulumi:"method"`
	// The ID of an existing azure source.
	ResourceId string `pulumi:"resourceId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ResourceAction resource.
type ResourceActionArgs struct {
	// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
	Action pulumi.StringPtrInput
	// A JSON object that contains the request body.
	Body pulumi.StringPtrInput
	// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
	Locks pulumi.StringArrayInput
	// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
	Method pulumi.StringPtrInput
	// The ID of an existing azure source.
	ResourceId pulumi.StringInput
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
	ResponseExportValues pulumi.StringArrayInput
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput
}

func (ResourceActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceActionArgs)(nil)).Elem()
}

type ResourceActionInput interface {
	pulumi.Input

	ToResourceActionOutput() ResourceActionOutput
	ToResourceActionOutputWithContext(ctx context.Context) ResourceActionOutput
}

func (*ResourceAction) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAction)(nil)).Elem()
}

func (i *ResourceAction) ToResourceActionOutput() ResourceActionOutput {
	return i.ToResourceActionOutputWithContext(context.Background())
}

func (i *ResourceAction) ToResourceActionOutputWithContext(ctx context.Context) ResourceActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceActionOutput)
}

// ResourceActionArrayInput is an input type that accepts ResourceActionArray and ResourceActionArrayOutput values.
// You can construct a concrete instance of `ResourceActionArrayInput` via:
//
//	ResourceActionArray{ ResourceActionArgs{...} }
type ResourceActionArrayInput interface {
	pulumi.Input

	ToResourceActionArrayOutput() ResourceActionArrayOutput
	ToResourceActionArrayOutputWithContext(context.Context) ResourceActionArrayOutput
}

type ResourceActionArray []ResourceActionInput

func (ResourceActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceAction)(nil)).Elem()
}

func (i ResourceActionArray) ToResourceActionArrayOutput() ResourceActionArrayOutput {
	return i.ToResourceActionArrayOutputWithContext(context.Background())
}

func (i ResourceActionArray) ToResourceActionArrayOutputWithContext(ctx context.Context) ResourceActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceActionArrayOutput)
}

// ResourceActionMapInput is an input type that accepts ResourceActionMap and ResourceActionMapOutput values.
// You can construct a concrete instance of `ResourceActionMapInput` via:
//
//	ResourceActionMap{ "key": ResourceActionArgs{...} }
type ResourceActionMapInput interface {
	pulumi.Input

	ToResourceActionMapOutput() ResourceActionMapOutput
	ToResourceActionMapOutputWithContext(context.Context) ResourceActionMapOutput
}

type ResourceActionMap map[string]ResourceActionInput

func (ResourceActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceAction)(nil)).Elem()
}

func (i ResourceActionMap) ToResourceActionMapOutput() ResourceActionMapOutput {
	return i.ToResourceActionMapOutputWithContext(context.Background())
}

func (i ResourceActionMap) ToResourceActionMapOutputWithContext(ctx context.Context) ResourceActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceActionMapOutput)
}

type ResourceActionOutput struct{ *pulumi.OutputState }

func (ResourceActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAction)(nil)).Elem()
}

func (o ResourceActionOutput) ToResourceActionOutput() ResourceActionOutput {
	return o
}

func (o ResourceActionOutput) ToResourceActionOutputWithContext(ctx context.Context) ResourceActionOutput {
	return o
}

// The name of the resource action. It's also possible to make Http requests towards the resource ID if leave this field empty.
func (o ResourceActionOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// A JSON object that contains the request body.
func (o ResourceActionOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// A list of ARM resource IDs which are used to avoid modify azapi resources at the same time.
func (o ResourceActionOutput) Locks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringArrayOutput { return v.Locks }).(pulumi.StringArrayOutput)
}

// Specifies the Http method of the azure resource action. Allowed values are `POST`, `PATCH`, `PUT` and `DELETE`. Defaults to `POST`.
func (o ResourceActionOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringPtrOutput { return v.Method }).(pulumi.StringPtrOutput)
}

// The output json containing the properties specified in `responseExportValues`. Here are some examples to decode json and extract the value.
func (o ResourceActionOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringOutput { return v.Output }).(pulumi.StringOutput)
}

// The ID of an existing azure source.
func (o ResourceActionOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// A list of path that needs to be exported from response body.
// Setting it to `["*"]` will export the full response body.
// Here's an example. If it sets to `["keys"]`, it will set the following json to computed property `output`.
func (o ResourceActionOutput) ResponseExportValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringArrayOutput { return v.ResponseExportValues }).(pulumi.StringArrayOutput)
}

// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
// `<api-version>` is version of the API used to manage this azure resource.
func (o ResourceActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ResourceActionArrayOutput struct{ *pulumi.OutputState }

func (ResourceActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceAction)(nil)).Elem()
}

func (o ResourceActionArrayOutput) ToResourceActionArrayOutput() ResourceActionArrayOutput {
	return o
}

func (o ResourceActionArrayOutput) ToResourceActionArrayOutputWithContext(ctx context.Context) ResourceActionArrayOutput {
	return o
}

func (o ResourceActionArrayOutput) Index(i pulumi.IntInput) ResourceActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceAction {
		return vs[0].([]*ResourceAction)[vs[1].(int)]
	}).(ResourceActionOutput)
}

type ResourceActionMapOutput struct{ *pulumi.OutputState }

func (ResourceActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceAction)(nil)).Elem()
}

func (o ResourceActionMapOutput) ToResourceActionMapOutput() ResourceActionMapOutput {
	return o
}

func (o ResourceActionMapOutput) ToResourceActionMapOutputWithContext(ctx context.Context) ResourceActionMapOutput {
	return o
}

func (o ResourceActionMapOutput) MapIndex(k pulumi.StringInput) ResourceActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceAction {
		return vs[0].(map[string]*ResourceAction)[vs[1].(string)]
	}).(ResourceActionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceActionInput)(nil)).Elem(), &ResourceAction{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceActionArrayInput)(nil)).Elem(), ResourceActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceActionMapInput)(nil)).Elem(), ResourceActionMap{})
	pulumi.RegisterOutputType(ResourceActionOutput{})
	pulumi.RegisterOutputType(ResourceActionArrayOutput{})
	pulumi.RegisterOutputType(ResourceActionMapOutput{})
}
