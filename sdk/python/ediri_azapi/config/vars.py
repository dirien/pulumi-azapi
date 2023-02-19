# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('azapi')


class _ExportableConfig(types.ModuleType):
    @property
    def client_certificate_password(self) -> Optional[str]:
        """
        The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return __config__.get('clientCertificatePassword')

    @property
    def client_certificate_path(self) -> Optional[str]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return __config__.get('clientCertificatePath')

    @property
    def client_id(self) -> Optional[str]:
        """
        The Client ID which should be used.
        """
        return __config__.get('clientId')

    @property
    def client_secret(self) -> Optional[str]:
        """
        The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        """
        return __config__.get('clientSecret')

    @property
    def default_location(self) -> Optional[str]:
        return __config__.get('defaultLocation')

    @property
    def default_tags(self) -> Optional[str]:
        return __config__.get('defaultTags')

    @property
    def disable_correlation_request_id(self) -> Optional[bool]:
        """
        This will disable the x-ms-correlation-request-id header.
        """
        return __config__.get_bool('disableCorrelationRequestId')

    @property
    def disable_terraform_partner_id(self) -> Optional[bool]:
        """
        This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        """
        return __config__.get_bool('disableTerraformPartnerId')

    @property
    def environment(self) -> Optional[str]:
        """
        The Cloud Environment which should be used. Possible values are public, usgovernment and china. Defaults to public.
        """
        return __config__.get('environment')

    @property
    def oidc_request_token(self) -> Optional[str]:
        """
        The bearer token for the request to the OIDC provider. For use When authenticating as a Service Principal using OpenID
        Connect.
        """
        return __config__.get('oidcRequestToken')

    @property
    def oidc_request_url(self) -> Optional[str]:
        """
        The URL for the OIDC provider from which to request an ID token. For use When authenticating as a Service Principal
        using OpenID Connect.
        """
        return __config__.get('oidcRequestUrl')

    @property
    def oidc_token(self) -> Optional[str]:
        """
        The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
        """
        return __config__.get('oidcToken')

    @property
    def oidc_token_file_path(self) -> Optional[str]:
        """
        The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
        """
        return __config__.get('oidcTokenFilePath')

    @property
    def partner_id(self) -> Optional[str]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return __config__.get('partnerId')

    @property
    def skip_provider_registration(self) -> Optional[bool]:
        """
        Should the Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
        """
        return __config__.get_bool('skipProviderRegistration')

    @property
    def subscription_id(self) -> Optional[str]:
        """
        The Subscription ID which should be used.
        """
        return __config__.get('subscriptionId')

    @property
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID which should be used.
        """
        return __config__.get('tenantId')

    @property
    def use_oidc(self) -> Optional[bool]:
        """
        Allow OpenID Connect to be used for authentication
        """
        return __config__.get_bool('useOidc')
