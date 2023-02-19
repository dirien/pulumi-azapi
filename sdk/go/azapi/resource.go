// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage any Azure resource manager resource.
//
// ## Import
//
// Azure resource can be imported using the `resource id`, e.g.
//
// ```sh
//
//	$ pulumi import azapi:index/resource:Resource example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
//
// ```
//
//	It also supports specifying API version by using the `resource id` with `api-version` as a query parameter, e.g.
//
// ```sh
//
//	$ pulumi import azapi:index/resource:Resource example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1?api-version=2021-07-01
//
// ```
type Resource struct {
	pulumi.CustomResourceState

	// A JSON object that contains the request body used to create and update azure resource.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// A `identity` block as defined below.
	Identity ResourceIdentityOutput `pulumi:"identity"`
	// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
	IgnoreCasing pulumi.BoolPtrOutput `pulumi:"ignoreCasing"`
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
	IgnoreMissingProperty pulumi.BoolPtrOutput `pulumi:"ignoreMissingProperty"`
	// The Azure Region where the azure resource should exist.
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	Locks pulumi.StringArrayOutput `pulumi:"locks"`
	// Specifies the name of the azure resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
	Output   pulumi.StringOutput `pulumi:"output"`
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	ResponseExportValues pulumi.StringArrayOutput `pulumi:"responseExportValues"`
	// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
	SchemaValidationEnabled pulumi.BoolPtrOutput `pulumi:"schemaValidationEnabled"`
	// A mapping of tags which should be assigned to the azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("azapi:index/resource:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("azapi:index/resource:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
	// A JSON object that contains the request body used to create and update azure resource.
	Body *string `pulumi:"body"`
	// A `identity` block as defined below.
	Identity *ResourceIdentity `pulumi:"identity"`
	// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
	IgnoreCasing *bool `pulumi:"ignoreCasing"`
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
	IgnoreMissingProperty *bool `pulumi:"ignoreMissingProperty"`
	// The Azure Region where the azure resource should exist.
	Location *string `pulumi:"location"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	Locks []string `pulumi:"locks"`
	// Specifies the name of the azure resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
	Output   *string `pulumi:"output"`
	ParentId *string `pulumi:"parentId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
	SchemaValidationEnabled *bool `pulumi:"schemaValidationEnabled"`
	// A mapping of tags which should be assigned to the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type *string `pulumi:"type"`
}

type ResourceState struct {
	// A JSON object that contains the request body used to create and update azure resource.
	Body pulumi.StringPtrInput
	// A `identity` block as defined below.
	Identity ResourceIdentityPtrInput
	// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
	IgnoreCasing pulumi.BoolPtrInput
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
	IgnoreMissingProperty pulumi.BoolPtrInput
	// The Azure Region where the azure resource should exist.
	Location pulumi.StringPtrInput
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	Locks pulumi.StringArrayInput
	// Specifies the name of the azure resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
	Output   pulumi.StringPtrInput
	ParentId pulumi.StringPtrInput
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	ResponseExportValues pulumi.StringArrayInput
	// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
	SchemaValidationEnabled pulumi.BoolPtrInput
	// A mapping of tags which should be assigned to the azure resource.
	Tags pulumi.StringMapInput
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringPtrInput
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// A JSON object that contains the request body used to create and update azure resource.
	Body *string `pulumi:"body"`
	// A `identity` block as defined below.
	Identity *ResourceIdentity `pulumi:"identity"`
	// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
	IgnoreCasing *bool `pulumi:"ignoreCasing"`
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
	IgnoreMissingProperty *bool `pulumi:"ignoreMissingProperty"`
	// The Azure Region where the azure resource should exist.
	Location *string `pulumi:"location"`
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	Locks []string `pulumi:"locks"`
	// Specifies the name of the azure resource. Changing this forces a new resource to be created.
	Name     *string `pulumi:"name"`
	ParentId string  `pulumi:"parentId"`
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	ResponseExportValues []string `pulumi:"responseExportValues"`
	// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
	SchemaValidationEnabled *bool `pulumi:"schemaValidationEnabled"`
	// A mapping of tags which should be assigned to the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// A JSON object that contains the request body used to create and update azure resource.
	Body pulumi.StringPtrInput
	// A `identity` block as defined below.
	Identity ResourceIdentityPtrInput
	// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
	IgnoreCasing pulumi.BoolPtrInput
	// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
	IgnoreMissingProperty pulumi.BoolPtrInput
	// The Azure Region where the azure resource should exist.
	Location pulumi.StringPtrInput
	// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
	Locks pulumi.StringArrayInput
	// Specifies the name of the azure resource. Changing this forces a new resource to be created.
	Name     pulumi.StringPtrInput
	ParentId pulumi.StringInput
	// A list of path that needs to be exported from response body.
	// Setting it to `["*"]` will export the full response body.
	// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	ResponseExportValues pulumi.StringArrayInput
	// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
	SchemaValidationEnabled pulumi.BoolPtrInput
	// A mapping of tags which should be assigned to the azure resource.
	Tags pulumi.StringMapInput
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

// ResourceArrayInput is an input type that accepts ResourceArray and ResourceArrayOutput values.
// You can construct a concrete instance of `ResourceArrayInput` via:
//
//	ResourceArray{ ResourceArgs{...} }
type ResourceArrayInput interface {
	pulumi.Input

	ToResourceArrayOutput() ResourceArrayOutput
	ToResourceArrayOutputWithContext(context.Context) ResourceArrayOutput
}

type ResourceArray []ResourceInput

func (ResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (i ResourceArray) ToResourceArrayOutput() ResourceArrayOutput {
	return i.ToResourceArrayOutputWithContext(context.Background())
}

func (i ResourceArray) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceArrayOutput)
}

// ResourceMapInput is an input type that accepts ResourceMap and ResourceMapOutput values.
// You can construct a concrete instance of `ResourceMapInput` via:
//
//	ResourceMap{ "key": ResourceArgs{...} }
type ResourceMapInput interface {
	pulumi.Input

	ToResourceMapOutput() ResourceMapOutput
	ToResourceMapOutputWithContext(context.Context) ResourceMapOutput
}

type ResourceMap map[string]ResourceInput

func (ResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (i ResourceMap) ToResourceMapOutput() ResourceMapOutput {
	return i.ToResourceMapOutputWithContext(context.Background())
}

func (i ResourceMap) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMapOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// A JSON object that contains the request body used to create and update azure resource.
func (o ResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// A `identity` block as defined below.
func (o ResourceOutput) Identity() ResourceIdentityOutput {
	return o.ApplyT(func(v *Resource) ResourceIdentityOutput { return v.Identity }).(ResourceIdentityOutput)
}

// Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
func (o ResourceOutput) IgnoreCasing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.BoolPtrOutput { return v.IgnoreCasing }).(pulumi.BoolPtrOutput)
}

// Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
func (o ResourceOutput) IgnoreMissingProperty() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.BoolPtrOutput { return v.IgnoreMissingProperty }).(pulumi.BoolPtrOutput)
}

// The Azure Region where the azure resource should exist.
func (o ResourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
func (o ResourceOutput) Locks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringArrayOutput { return v.Locks }).(pulumi.StringArrayOutput)
}

// Specifies the name of the azure resource. Changing this forces a new resource to be created.
func (o ResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The output json containing the properties specified in `responseExportValues`. Here're some examples to decode json and extract the value.
func (o ResourceOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Output }).(pulumi.StringOutput)
}

func (o ResourceOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// A list of path that needs to be exported from response body.
// Setting it to `["*"]` will export the full response body.
// Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
func (o ResourceOutput) ResponseExportValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringArrayOutput { return v.ResponseExportValues }).(pulumi.StringArrayOutput)
}

// Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
func (o ResourceOutput) SchemaValidationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.BoolPtrOutput { return v.SchemaValidationEnabled }).(pulumi.BoolPtrOutput)
}

// A mapping of tags which should be assigned to the azure resource.
func (o ResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
// `<api-version>` is version of the API used to manage this azure resource.
func (o ResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ResourceArrayOutput struct{ *pulumi.OutputState }

func (ResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (o ResourceArrayOutput) ToResourceArrayOutput() ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) Index(i pulumi.IntInput) ResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].([]*Resource)[vs[1].(int)]
	}).(ResourceOutput)
}

type ResourceMapOutput struct{ *pulumi.OutputState }

func (ResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (o ResourceMapOutput) ToResourceMapOutput() ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) MapIndex(k pulumi.StringInput) ResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].(map[string]*Resource)[vs[1].(string)]
	}).(ResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceArrayInput)(nil)).Elem(), ResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMapInput)(nil)).Elem(), ResourceMap{})
	pulumi.RegisterOutputType(ResourceOutput{})
	pulumi.RegisterOutputType(ResourceArrayOutput{})
	pulumi.RegisterOutputType(ResourceMapOutput{})
}