// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("azapi");

export declare const auxiliaryTenantIds: string[] | undefined;
Object.defineProperty(exports, "auxiliaryTenantIds", {
    get() {
        return __config.getObject<string[]>("auxiliaryTenantIds");
    },
    enumerable: true,
});

/**
 * The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
 * Certificate
 */
export declare const clientCertificatePassword: string | undefined;
Object.defineProperty(exports, "clientCertificatePassword", {
    get() {
        return __config.get("clientCertificatePassword");
    },
    enumerable: true,
});

/**
 * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
 * Principal using a Client Certificate.
 */
export declare const clientCertificatePath: string | undefined;
Object.defineProperty(exports, "clientCertificatePath", {
    get() {
        return __config.get("clientCertificatePath");
    },
    enumerable: true,
});

/**
 * The Client ID which should be used.
 */
export declare const clientId: string | undefined;
Object.defineProperty(exports, "clientId", {
    get() {
        return __config.get("clientId");
    },
    enumerable: true,
});

/**
 * The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
 */
export declare const clientSecret: string | undefined;
Object.defineProperty(exports, "clientSecret", {
    get() {
        return __config.get("clientSecret");
    },
    enumerable: true,
});

/**
 * The value of the x-ms-correlation-request-id header (otherwise an auto-generated UUID will be used).
 */
export declare const customCorrelationRequestId: string | undefined;
Object.defineProperty(exports, "customCorrelationRequestId", {
    get() {
        return __config.get("customCorrelationRequestId");
    },
    enumerable: true,
});

export declare const defaultLocation: string | undefined;
Object.defineProperty(exports, "defaultLocation", {
    get() {
        return __config.get("defaultLocation");
    },
    enumerable: true,
});

export declare const defaultName: string | undefined;
Object.defineProperty(exports, "defaultName", {
    get() {
        return __config.get("defaultName");
    },
    enumerable: true,
});

export declare const defaultNamingPrefix: string | undefined;
Object.defineProperty(exports, "defaultNamingPrefix", {
    get() {
        return __config.get("defaultNamingPrefix");
    },
    enumerable: true,
});

export declare const defaultNamingSuffix: string | undefined;
Object.defineProperty(exports, "defaultNamingSuffix", {
    get() {
        return __config.get("defaultNamingSuffix");
    },
    enumerable: true,
});

export declare const defaultTags: {[key: string]: string} | undefined;
Object.defineProperty(exports, "defaultTags", {
    get() {
        return __config.getObject<{[key: string]: string}>("defaultTags");
    },
    enumerable: true,
});

/**
 * This will disable the x-ms-correlation-request-id header.
 */
export declare const disableCorrelationRequestId: boolean | undefined;
Object.defineProperty(exports, "disableCorrelationRequestId", {
    get() {
        return __config.getObject<boolean>("disableCorrelationRequestId");
    },
    enumerable: true,
});

export declare const disableTerraformPartnerId: boolean | undefined;
Object.defineProperty(exports, "disableTerraformPartnerId", {
    get() {
        return __config.getObject<boolean>("disableTerraformPartnerId");
    },
    enumerable: true,
});

export declare const endpoint: outputs.config.Endpoint | undefined;
Object.defineProperty(exports, "endpoint", {
    get() {
        return __config.getObject<outputs.config.Endpoint>("endpoint");
    },
    enumerable: true,
});

/**
 * The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
 */
export declare const environment: string | undefined;
Object.defineProperty(exports, "environment", {
    get() {
        return __config.get("environment");
    },
    enumerable: true,
});

/**
 * The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
 * Connect.
 */
export declare const oidcRequestToken: string | undefined;
Object.defineProperty(exports, "oidcRequestToken", {
    get() {
        return __config.get("oidcRequestToken");
    },
    enumerable: true,
});

/**
 * The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
 * using OpenID Connect.
 */
export declare const oidcRequestUrl: string | undefined;
Object.defineProperty(exports, "oidcRequestUrl", {
    get() {
        return __config.get("oidcRequestUrl");
    },
    enumerable: true,
});

/**
 * The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
 */
export declare const oidcToken: string | undefined;
Object.defineProperty(exports, "oidcToken", {
    get() {
        return __config.get("oidcToken");
    },
    enumerable: true,
});

/**
 * The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
 */
export declare const oidcTokenFilePath: string | undefined;
Object.defineProperty(exports, "oidcTokenFilePath", {
    get() {
        return __config.get("oidcTokenFilePath");
    },
    enumerable: true,
});

/**
 * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
 */
export declare const partnerId: string | undefined;
Object.defineProperty(exports, "partnerId", {
    get() {
        return __config.get("partnerId");
    },
    enumerable: true,
});

/**
 * Should the Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
 */
export declare const skipProviderRegistration: boolean | undefined;
Object.defineProperty(exports, "skipProviderRegistration", {
    get() {
        return __config.getObject<boolean>("skipProviderRegistration");
    },
    enumerable: true,
});

/**
 * The Subscription ID which should be used.
 */
export declare const subscriptionId: string | undefined;
Object.defineProperty(exports, "subscriptionId", {
    get() {
        return __config.get("subscriptionId");
    },
    enumerable: true,
});

/**
 * The Tenant ID which should be used.
 */
export declare const tenantId: string | undefined;
Object.defineProperty(exports, "tenantId", {
    get() {
        return __config.get("tenantId");
    },
    enumerable: true,
});

/**
 * Allow Azure CLI to be used for Authentication.
 */
export declare const useCli: boolean | undefined;
Object.defineProperty(exports, "useCli", {
    get() {
        return __config.getObject<boolean>("useCli");
    },
    enumerable: true,
});

/**
 * Allow Managed Service Identity to be used for Authentication.
 */
export declare const useMsi: boolean | undefined;
Object.defineProperty(exports, "useMsi", {
    get() {
        return __config.getObject<boolean>("useMsi");
    },
    enumerable: true,
});

/**
 * Allow OpenID Connect to be used for authentication
 */
export declare const useOidc: boolean | undefined;
Object.defineProperty(exports, "useOidc", {
    get() {
        return __config.getObject<boolean>("useOidc");
    },
    enumerable: true,
});

