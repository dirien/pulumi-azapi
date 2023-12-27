// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-azapi/sdk/go/azapi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ProviderEndpoint struct {
	ActiveDirectoryAuthorityHost *string `pulumi:"activeDirectoryAuthorityHost"`
	ResourceManagerAudience      *string `pulumi:"resourceManagerAudience"`
	ResourceManagerEndpoint      *string `pulumi:"resourceManagerEndpoint"`
}

// ProviderEndpointInput is an input type that accepts ProviderEndpointArgs and ProviderEndpointOutput values.
// You can construct a concrete instance of `ProviderEndpointInput` via:
//
//	ProviderEndpointArgs{...}
type ProviderEndpointInput interface {
	pulumi.Input

	ToProviderEndpointOutput() ProviderEndpointOutput
	ToProviderEndpointOutputWithContext(context.Context) ProviderEndpointOutput
}

type ProviderEndpointArgs struct {
	ActiveDirectoryAuthorityHost pulumi.StringPtrInput `pulumi:"activeDirectoryAuthorityHost"`
	ResourceManagerAudience      pulumi.StringPtrInput `pulumi:"resourceManagerAudience"`
	ResourceManagerEndpoint      pulumi.StringPtrInput `pulumi:"resourceManagerEndpoint"`
}

func (ProviderEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderEndpoint)(nil)).Elem()
}

func (i ProviderEndpointArgs) ToProviderEndpointOutput() ProviderEndpointOutput {
	return i.ToProviderEndpointOutputWithContext(context.Background())
}

func (i ProviderEndpointArgs) ToProviderEndpointOutputWithContext(ctx context.Context) ProviderEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderEndpointOutput)
}

func (i ProviderEndpointArgs) ToProviderEndpointPtrOutput() ProviderEndpointPtrOutput {
	return i.ToProviderEndpointPtrOutputWithContext(context.Background())
}

func (i ProviderEndpointArgs) ToProviderEndpointPtrOutputWithContext(ctx context.Context) ProviderEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderEndpointOutput).ToProviderEndpointPtrOutputWithContext(ctx)
}

// ProviderEndpointPtrInput is an input type that accepts ProviderEndpointArgs, ProviderEndpointPtr and ProviderEndpointPtrOutput values.
// You can construct a concrete instance of `ProviderEndpointPtrInput` via:
//
//	        ProviderEndpointArgs{...}
//
//	or:
//
//	        nil
type ProviderEndpointPtrInput interface {
	pulumi.Input

	ToProviderEndpointPtrOutput() ProviderEndpointPtrOutput
	ToProviderEndpointPtrOutputWithContext(context.Context) ProviderEndpointPtrOutput
}

type providerEndpointPtrType ProviderEndpointArgs

func ProviderEndpointPtr(v *ProviderEndpointArgs) ProviderEndpointPtrInput {
	return (*providerEndpointPtrType)(v)
}

func (*providerEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderEndpoint)(nil)).Elem()
}

func (i *providerEndpointPtrType) ToProviderEndpointPtrOutput() ProviderEndpointPtrOutput {
	return i.ToProviderEndpointPtrOutputWithContext(context.Background())
}

func (i *providerEndpointPtrType) ToProviderEndpointPtrOutputWithContext(ctx context.Context) ProviderEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderEndpointPtrOutput)
}

type ProviderEndpointOutput struct{ *pulumi.OutputState }

func (ProviderEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderEndpoint)(nil)).Elem()
}

func (o ProviderEndpointOutput) ToProviderEndpointOutput() ProviderEndpointOutput {
	return o
}

func (o ProviderEndpointOutput) ToProviderEndpointOutputWithContext(ctx context.Context) ProviderEndpointOutput {
	return o
}

func (o ProviderEndpointOutput) ToProviderEndpointPtrOutput() ProviderEndpointPtrOutput {
	return o.ToProviderEndpointPtrOutputWithContext(context.Background())
}

func (o ProviderEndpointOutput) ToProviderEndpointPtrOutputWithContext(ctx context.Context) ProviderEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderEndpoint) *ProviderEndpoint {
		return &v
	}).(ProviderEndpointPtrOutput)
}

func (o ProviderEndpointOutput) ActiveDirectoryAuthorityHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderEndpoint) *string { return v.ActiveDirectoryAuthorityHost }).(pulumi.StringPtrOutput)
}

func (o ProviderEndpointOutput) ResourceManagerAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderEndpoint) *string { return v.ResourceManagerAudience }).(pulumi.StringPtrOutput)
}

func (o ProviderEndpointOutput) ResourceManagerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderEndpoint) *string { return v.ResourceManagerEndpoint }).(pulumi.StringPtrOutput)
}

type ProviderEndpointPtrOutput struct{ *pulumi.OutputState }

func (ProviderEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderEndpoint)(nil)).Elem()
}

func (o ProviderEndpointPtrOutput) ToProviderEndpointPtrOutput() ProviderEndpointPtrOutput {
	return o
}

func (o ProviderEndpointPtrOutput) ToProviderEndpointPtrOutputWithContext(ctx context.Context) ProviderEndpointPtrOutput {
	return o
}

func (o ProviderEndpointPtrOutput) Elem() ProviderEndpointOutput {
	return o.ApplyT(func(v *ProviderEndpoint) ProviderEndpoint {
		if v != nil {
			return *v
		}
		var ret ProviderEndpoint
		return ret
	}).(ProviderEndpointOutput)
}

func (o ProviderEndpointPtrOutput) ActiveDirectoryAuthorityHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryAuthorityHost
	}).(pulumi.StringPtrOutput)
}

func (o ProviderEndpointPtrOutput) ResourceManagerAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.ResourceManagerAudience
	}).(pulumi.StringPtrOutput)
}

func (o ProviderEndpointPtrOutput) ResourceManagerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.ResourceManagerEndpoint
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentity struct {
	// A list of User Managed Identity ID's which should be assigned to the azure resource.
	IdentityIds []string `pulumi:"identityIds"`
	// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	PrincipalId *string `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	TenantId *string `pulumi:"tenantId"`
	// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
	Type string `pulumi:"type"`
}

// ResourceIdentityInput is an input type that accepts ResourceIdentityArgs and ResourceIdentityOutput values.
// You can construct a concrete instance of `ResourceIdentityInput` via:
//
//	ResourceIdentityArgs{...}
type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	// A list of User Managed Identity ID's which should be assigned to the azure resource.
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}

// ResourceIdentityPtrInput is an input type that accepts ResourceIdentityArgs, ResourceIdentityPtr and ResourceIdentityPtrOutput values.
// You can construct a concrete instance of `ResourceIdentityPtrInput` via:
//
//	        ResourceIdentityArgs{...}
//
//	or:
//
//	        nil
type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

// A list of User Managed Identity ID's which should be assigned to the azure resource.
func (o ResourceIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o ResourceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o ResourceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
func (o ResourceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

// A list of User Managed Identity ID's which should be assigned to the azure resource.
func (o ResourceIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o ResourceIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o ResourceIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type GetResourceIdentity struct {
	// A list of User Managed Identity ID's which should be assigned to the azure resource.
	IdentityIds []string `pulumi:"identityIds"`
	// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	PrincipalId string `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	TenantId string `pulumi:"tenantId"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type string `pulumi:"type"`
}

// GetResourceIdentityInput is an input type that accepts GetResourceIdentityArgs and GetResourceIdentityOutput values.
// You can construct a concrete instance of `GetResourceIdentityInput` via:
//
//	GetResourceIdentityArgs{...}
type GetResourceIdentityInput interface {
	pulumi.Input

	ToGetResourceIdentityOutput() GetResourceIdentityOutput
	ToGetResourceIdentityOutputWithContext(context.Context) GetResourceIdentityOutput
}

type GetResourceIdentityArgs struct {
	// A list of User Managed Identity ID's which should be assigned to the azure resource.
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
	// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
	// `<api-version>` is version of the API used to manage this azure resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceIdentity)(nil)).Elem()
}

func (i GetResourceIdentityArgs) ToGetResourceIdentityOutput() GetResourceIdentityOutput {
	return i.ToGetResourceIdentityOutputWithContext(context.Background())
}

func (i GetResourceIdentityArgs) ToGetResourceIdentityOutputWithContext(ctx context.Context) GetResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourceIdentityOutput)
}

func (i GetResourceIdentityArgs) ToGetResourceIdentityPtrOutput() GetResourceIdentityPtrOutput {
	return i.ToGetResourceIdentityPtrOutputWithContext(context.Background())
}

func (i GetResourceIdentityArgs) ToGetResourceIdentityPtrOutputWithContext(ctx context.Context) GetResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourceIdentityOutput).ToGetResourceIdentityPtrOutputWithContext(ctx)
}

// GetResourceIdentityPtrInput is an input type that accepts GetResourceIdentityArgs, GetResourceIdentityPtr and GetResourceIdentityPtrOutput values.
// You can construct a concrete instance of `GetResourceIdentityPtrInput` via:
//
//	        GetResourceIdentityArgs{...}
//
//	or:
//
//	        nil
type GetResourceIdentityPtrInput interface {
	pulumi.Input

	ToGetResourceIdentityPtrOutput() GetResourceIdentityPtrOutput
	ToGetResourceIdentityPtrOutputWithContext(context.Context) GetResourceIdentityPtrOutput
}

type getResourceIdentityPtrType GetResourceIdentityArgs

func GetResourceIdentityPtr(v *GetResourceIdentityArgs) GetResourceIdentityPtrInput {
	return (*getResourceIdentityPtrType)(v)
}

func (*getResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GetResourceIdentity)(nil)).Elem()
}

func (i *getResourceIdentityPtrType) ToGetResourceIdentityPtrOutput() GetResourceIdentityPtrOutput {
	return i.ToGetResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *getResourceIdentityPtrType) ToGetResourceIdentityPtrOutputWithContext(ctx context.Context) GetResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourceIdentityPtrOutput)
}

type GetResourceIdentityOutput struct{ *pulumi.OutputState }

func (GetResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceIdentity)(nil)).Elem()
}

func (o GetResourceIdentityOutput) ToGetResourceIdentityOutput() GetResourceIdentityOutput {
	return o
}

func (o GetResourceIdentityOutput) ToGetResourceIdentityOutputWithContext(ctx context.Context) GetResourceIdentityOutput {
	return o
}

func (o GetResourceIdentityOutput) ToGetResourceIdentityPtrOutput() GetResourceIdentityPtrOutput {
	return o.ToGetResourceIdentityPtrOutputWithContext(context.Background())
}

func (o GetResourceIdentityOutput) ToGetResourceIdentityPtrOutputWithContext(ctx context.Context) GetResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GetResourceIdentity) *GetResourceIdentity {
		return &v
	}).(GetResourceIdentityPtrOutput)
}

// A list of User Managed Identity ID's which should be assigned to the azure resource.
func (o GetResourceIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourceIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o GetResourceIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o GetResourceIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
// `<api-version>` is version of the API used to manage this azure resource.
func (o GetResourceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type GetResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (GetResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetResourceIdentity)(nil)).Elem()
}

func (o GetResourceIdentityPtrOutput) ToGetResourceIdentityPtrOutput() GetResourceIdentityPtrOutput {
	return o
}

func (o GetResourceIdentityPtrOutput) ToGetResourceIdentityPtrOutputWithContext(ctx context.Context) GetResourceIdentityPtrOutput {
	return o
}

func (o GetResourceIdentityPtrOutput) Elem() GetResourceIdentityOutput {
	return o.ApplyT(func(v *GetResourceIdentity) GetResourceIdentity {
		if v != nil {
			return *v
		}
		var ret GetResourceIdentity
		return ret
	}).(GetResourceIdentityOutput)
}

// A list of User Managed Identity ID's which should be assigned to the azure resource.
func (o GetResourceIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GetResourceIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

// The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o GetResourceIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
func (o GetResourceIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
// `<api-version>` is version of the API used to manage this azure resource.
func (o GetResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GetResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderEndpointInput)(nil)).Elem(), ProviderEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderEndpointPtrInput)(nil)).Elem(), ProviderEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityPtrInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetResourceIdentityInput)(nil)).Elem(), GetResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetResourceIdentityPtrInput)(nil)).Elem(), GetResourceIdentityArgs{})
	pulumi.RegisterOutputType(ProviderEndpointOutput{})
	pulumi.RegisterOutputType(ProviderEndpointPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(GetResourceIdentityOutput{})
	pulumi.RegisterOutputType(GetResourceIdentityPtrOutput{})
}
