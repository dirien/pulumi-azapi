// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azapi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azapi/sdk/go/azapi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the azapi package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword pulumi.StringPtrOutput `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate.
	ClientCertificatePath pulumi.StringPtrOutput `pulumi:"clientCertificatePath"`
	// The Client ID which should be used.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The value of the x-ms-correlation-request-id header (otherwise an auto-generated UUID will be used).
	CustomCorrelationRequestId pulumi.StringPtrOutput `pulumi:"customCorrelationRequestId"`
	DefaultLocation            pulumi.StringPtrOutput `pulumi:"defaultLocation"`
	DefaultName                pulumi.StringPtrOutput `pulumi:"defaultName"`
	DefaultNamingPrefix        pulumi.StringPtrOutput `pulumi:"defaultNamingPrefix"`
	DefaultNamingSuffix        pulumi.StringPtrOutput `pulumi:"defaultNamingSuffix"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken pulumi.StringPtrOutput `pulumi:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl pulumi.StringPtrOutput `pulumi:"oidcRequestUrl"`
	// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken pulumi.StringPtrOutput `pulumi:"oidcToken"`
	// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath pulumi.StringPtrOutput `pulumi:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId pulumi.StringPtrOutput `pulumi:"partnerId"`
	// The Subscription ID which should be used.
	SubscriptionId pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	// The Tenant ID which should be used.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:azapi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	AuxiliaryTenantIds []string `pulumi:"auxiliaryTenantIds"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword *string `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate.
	ClientCertificatePath *string `pulumi:"clientCertificatePath"`
	// The Client ID which should be used.
	ClientId *string `pulumi:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret *string `pulumi:"clientSecret"`
	// The value of the x-ms-correlation-request-id header (otherwise an auto-generated UUID will be used).
	CustomCorrelationRequestId *string           `pulumi:"customCorrelationRequestId"`
	DefaultLocation            *string           `pulumi:"defaultLocation"`
	DefaultName                *string           `pulumi:"defaultName"`
	DefaultNamingPrefix        *string           `pulumi:"defaultNamingPrefix"`
	DefaultNamingSuffix        *string           `pulumi:"defaultNamingSuffix"`
	DefaultTags                map[string]string `pulumi:"defaultTags"`
	// This will disable the x-ms-correlation-request-id header.
	DisableCorrelationRequestId *bool `pulumi:"disableCorrelationRequestId"`
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	DisableTerraformPartnerId *bool             `pulumi:"disableTerraformPartnerId"`
	Endpoint                  *ProviderEndpoint `pulumi:"endpoint"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
	Environment string `pulumi:"environment"`
	// The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken *string `pulumi:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl *string `pulumi:"oidcRequestUrl"`
	// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken *string `pulumi:"oidcToken"`
	// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath *string `pulumi:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId *string `pulumi:"partnerId"`
	// Should the Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	SkipProviderRegistration *bool `pulumi:"skipProviderRegistration"`
	// The Subscription ID which should be used.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Tenant ID which should be used.
	TenantId *string `pulumi:"tenantId"`
	// Allow Azure CLI to be used for Authentication.
	UseCli *bool `pulumi:"useCli"`
	// Allow Managed Service Identity to be used for Authentication.
	UseMsi *bool `pulumi:"useMsi"`
	// Allow OpenID Connect to be used for authentication
	UseOidc *bool `pulumi:"useOidc"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	AuxiliaryTenantIds pulumi.StringArrayInput
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword pulumi.StringPtrInput
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate.
	ClientCertificatePath pulumi.StringPtrInput
	// The Client ID which should be used.
	ClientId pulumi.StringPtrInput
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret pulumi.StringPtrInput
	// The value of the x-ms-correlation-request-id header (otherwise an auto-generated UUID will be used).
	CustomCorrelationRequestId pulumi.StringPtrInput
	DefaultLocation            pulumi.StringPtrInput
	DefaultName                pulumi.StringPtrInput
	DefaultNamingPrefix        pulumi.StringPtrInput
	DefaultNamingSuffix        pulumi.StringPtrInput
	DefaultTags                pulumi.StringMapInput
	// This will disable the x-ms-correlation-request-id header.
	DisableCorrelationRequestId pulumi.BoolPtrInput
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	DisableTerraformPartnerId pulumi.BoolPtrInput
	Endpoint                  ProviderEndpointPtrInput
	// The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
	Environment pulumi.StringInput
	// The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken pulumi.StringPtrInput
	// The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl pulumi.StringPtrInput
	// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken pulumi.StringPtrInput
	// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath pulumi.StringPtrInput
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId pulumi.StringPtrInput
	// Should the Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	SkipProviderRegistration pulumi.BoolPtrInput
	// The Subscription ID which should be used.
	SubscriptionId pulumi.StringPtrInput
	// The Tenant ID which should be used.
	TenantId pulumi.StringPtrInput
	// Allow Azure CLI to be used for Authentication.
	UseCli pulumi.BoolPtrInput
	// Allow Managed Service Identity to be used for Authentication.
	UseMsi pulumi.BoolPtrInput
	// Allow OpenID Connect to be used for authentication
	UseOidc pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
// Certificate
func (o ProviderOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertificatePassword }).(pulumi.StringPtrOutput)
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
// Principal using a Client Certificate.
func (o ProviderOutput) ClientCertificatePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertificatePath }).(pulumi.StringPtrOutput)
}

// The Client ID which should be used.
func (o ProviderOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
func (o ProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// The value of the x-ms-correlation-request-id header (otherwise an auto-generated UUID will be used).
func (o ProviderOutput) CustomCorrelationRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CustomCorrelationRequestId }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) DefaultLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DefaultLocation }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) DefaultName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DefaultName }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) DefaultNamingPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DefaultNamingPrefix }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) DefaultNamingSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DefaultNamingSuffix }).(pulumi.StringPtrOutput)
}

// The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
func (o ProviderOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
// Connect.
func (o ProviderOutput) OidcRequestToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcRequestToken }).(pulumi.StringPtrOutput)
}

// The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
// using OpenID Connect.
func (o ProviderOutput) OidcRequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcRequestUrl }).(pulumi.StringPtrOutput)
}

// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
func (o ProviderOutput) OidcToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcToken }).(pulumi.StringPtrOutput)
}

// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
func (o ProviderOutput) OidcTokenFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcTokenFilePath }).(pulumi.StringPtrOutput)
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
func (o ProviderOutput) PartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PartnerId }).(pulumi.StringPtrOutput)
}

// The Subscription ID which should be used.
func (o ProviderOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// The Tenant ID which should be used.
func (o ProviderOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
