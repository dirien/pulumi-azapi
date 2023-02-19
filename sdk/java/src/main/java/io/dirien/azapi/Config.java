// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi;

import com.pulumi.core.TypeShape;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("azapi");
/**
 * The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
 * Certificate
 * 
 */
    public Optional<String> clientCertificatePassword() {
        return Codegen.stringProp("clientCertificatePassword").config(config).get();
    }
/**
 * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
 * Principal using a Client Certificate.
 * 
 */
    public Optional<String> clientCertificatePath() {
        return Codegen.stringProp("clientCertificatePath").config(config).get();
    }
/**
 * The Client ID which should be used.
 * 
 */
    public Optional<String> clientId() {
        return Codegen.stringProp("clientId").config(config).get();
    }
/**
 * The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
 * 
 */
    public Optional<String> clientSecret() {
        return Codegen.stringProp("clientSecret").config(config).get();
    }
    public Optional<String> defaultLocation() {
        return Codegen.stringProp("defaultLocation").config(config).get();
    }
    public Optional<Map<String,String>> defaultTags() {
        return Codegen.objectProp("defaultTags", TypeShape.<Map<String,String>>builder(Map.class).addParameter(String.class).addParameter(String.class).build()).config(config).get();
    }
/**
 * This will disable the x-ms-correlation-request-id header.
 * 
 */
    public Optional<Boolean> disableCorrelationRequestId() {
        return Codegen.booleanProp("disableCorrelationRequestId").config(config).get();
    }
/**
 * This will disable the Terraform Partner ID which is used if a custom `partner_id` isn&#39;t specified.
 * 
 */
    public Optional<Boolean> disableTerraformPartnerId() {
        return Codegen.booleanProp("disableTerraformPartnerId").config(config).get();
    }
/**
 * The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
 * 
 */
    public String environment() {
        return Codegen.stringProp("environment").config(config).require();
    }
/**
 * The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
 * Connect.
 * 
 */
    public Optional<String> oidcRequestToken() {
        return Codegen.stringProp("oidcRequestToken").config(config).get();
    }
/**
 * The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
 * using OpenID Connect.
 * 
 */
    public Optional<String> oidcRequestUrl() {
        return Codegen.stringProp("oidcRequestUrl").config(config).get();
    }
/**
 * The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
 * 
 */
    public Optional<String> oidcToken() {
        return Codegen.stringProp("oidcToken").config(config).get();
    }
/**
 * The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
 * 
 */
    public Optional<String> oidcTokenFilePath() {
        return Codegen.stringProp("oidcTokenFilePath").config(config).get();
    }
/**
 * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
 * 
 */
    public Optional<String> partnerId() {
        return Codegen.stringProp("partnerId").config(config).get();
    }
/**
 * Should the Provider skip registering all of the Resource Providers that it supports, if they&#39;re not already registered?
 * 
 */
    public Optional<Boolean> skipProviderRegistration() {
        return Codegen.booleanProp("skipProviderRegistration").config(config).get();
    }
/**
 * The Subscription ID which should be used.
 * 
 */
    public Optional<String> subscriptionId() {
        return Codegen.stringProp("subscriptionId").config(config).get();
    }
/**
 * The Tenant ID which should be used.
 * 
 */
    public Optional<String> tenantId() {
        return Codegen.stringProp("tenantId").config(config).get();
    }
/**
 * Allow OpenID Connect to be used for authentication
 * 
 */
    public Optional<Boolean> useOidc() {
        return Codegen.booleanProp("useOidc").config(config).get();
    }
}
