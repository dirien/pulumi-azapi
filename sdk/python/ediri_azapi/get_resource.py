# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetResourceResult',
    'AwaitableGetResourceResult',
    'get_resource',
    'get_resource_output',
]

@pulumi.output_type
class GetResourceResult:
    """
    A collection of values returned by getResource.
    """
    def __init__(__self__, id=None, identity=None, location=None, name=None, output=None, parent_id=None, resource_id=None, response_export_values=None, tags=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output and not isinstance(output, str):
            raise TypeError("Expected argument 'output' to be a str")
        pulumi.set(__self__, "output", output)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if response_export_values and not isinstance(response_export_values, list):
            raise TypeError("Expected argument 'response_export_values' to be a list")
        pulumi.set(__self__, "response_export_values", response_export_values)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> 'outputs.GetResourceIdentityResult':
        """
        An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region where the azure resource should exist.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def output(self) -> str:
        """
        The output json containing the properties specified in `response_export_values`. Here're some examples to decode json and extract the value.
        """
        return pulumi.get(self, "output")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[str]:
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="responseExportValues")
    def response_export_values(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "response_export_values")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags which should be assigned to the azure resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
        """
        return pulumi.get(self, "type")


class AwaitableGetResourceResult(GetResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceResult(
            id=self.id,
            identity=self.identity,
            location=self.location,
            name=self.name,
            output=self.output,
            parent_id=self.parent_id,
            resource_id=self.resource_id,
            response_export_values=self.response_export_values,
            tags=self.tags,
            type=self.type)


def get_resource(identity: Optional[pulumi.InputType['GetResourceIdentityArgs']] = None,
                 name: Optional[str] = None,
                 parent_id: Optional[str] = None,
                 resource_id: Optional[str] = None,
                 response_export_values: Optional[Sequence[str]] = None,
                 type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceResult:
    """
    This resource can access any existing Azure resource manager resource.


    :param pulumi.InputType['GetResourceIdentityArgs'] identity: An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
    :param str name: Specifies the name of the azure resource.
    :param str parent_id: The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources: 
           - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
           - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
           - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
           - subscription scope: `parent_id` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
           - tenant scope: `parent_id` should be `/`
           
           For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
    :param str resource_id: The ID of an existing azure source.
           
           > **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
    :param Sequence[str] response_export_values: A list of path that needs to be exported from response body.
           Setting it to `["*"]` will export the full response body.
           Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
           ```python
           import pulumi
           ```
    :param str type: It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
           `<api-version>` is version of the API used to manage this azure resource.
    """
    __args__ = dict()
    __args__['identity'] = identity
    __args__['name'] = name
    __args__['parentId'] = parent_id
    __args__['resourceId'] = resource_id
    __args__['responseExportValues'] = response_export_values
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azapi:index/getResource:getResource', __args__, opts=opts, typ=GetResourceResult).value

    return AwaitableGetResourceResult(
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        output=pulumi.get(__ret__, 'output'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        response_export_values=pulumi.get(__ret__, 'response_export_values'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_resource)
def get_resource_output(identity: Optional[pulumi.Input[Optional[pulumi.InputType['GetResourceIdentityArgs']]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        parent_id: Optional[pulumi.Input[Optional[str]]] = None,
                        resource_id: Optional[pulumi.Input[Optional[str]]] = None,
                        response_export_values: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        type: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceResult]:
    """
    This resource can access any existing Azure resource manager resource.


    :param pulumi.InputType['GetResourceIdentityArgs'] identity: An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
    :param str name: Specifies the name of the azure resource.
    :param str parent_id: The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources: 
           - resource group scope: `parent_id` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.
           - management group scope: `parent_id` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.
           - extension scope: `parent_id` should be the ID of the resource you're adding the extension to.
           - subscription scope: `parent_id` should be like `/subscriptions/00000000-0000-0000-0000-000000000000`
           - tenant scope: `parent_id` should be `/`
           
           For child level resources, the `parent_id` should be the ID of its parent resource, for example, subnet resource's `parent_id` is the ID of the vnet.
    :param str resource_id: The ID of an existing azure source.
           
           > **Note:** Configuring `name` and `parent_id` is an alternative way to configure `resource_id`.
    :param Sequence[str] response_export_values: A list of path that needs to be exported from response body.
           Setting it to `["*"]` will export the full response body.
           Here's an example. If it sets to `["properties.loginServer", "properties.policies.quarantinePolicy.status"]`, it will set the following json to computed property `output`.
           ```python
           import pulumi
           ```
    :param str type: It is in a format like `<resource-type>@<api-version>`. `<resource-type>` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
           `<api-version>` is version of the API used to manage this azure resource.
    """
    ...
