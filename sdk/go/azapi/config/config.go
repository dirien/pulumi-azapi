// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
// Certificate
func GetClientCertificatePassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:clientCertificatePassword")
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
// Principal using a Client Certificate.
func GetClientCertificatePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:clientCertificatePath")
}

// The Client ID which should be used.
func GetClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:clientId")
}

// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
func GetClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:clientSecret")
}
func GetDefaultLocation(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:defaultLocation")
}
func GetDefaultTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:defaultTags")
}

// This will disable the x-ms-correlation-request-id header.
func GetDisableCorrelationRequestId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azapi:disableCorrelationRequestId")
}

// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
func GetDisableTerraformPartnerId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azapi:disableTerraformPartnerId")
}

// The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
func GetEnvironment(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:environment")
}

// The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
// Connect.
func GetOidcRequestToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:oidcRequestToken")
}

// The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
// using OpenID Connect.
func GetOidcRequestUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:oidcRequestUrl")
}

// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
func GetOidcToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:oidcToken")
}

// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
func GetOidcTokenFilePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:oidcTokenFilePath")
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
func GetPartnerId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:partnerId")
}

// Should the Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
func GetSkipProviderRegistration(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azapi:skipProviderRegistration")
}

// The Subscription ID which should be used.
func GetSubscriptionId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:subscriptionId")
}

// The Tenant ID which should be used.
func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azapi:tenantId")
}

// Allow OpenID Connect to be used for authentication
func GetUseOidc(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azapi:useOidc")
}