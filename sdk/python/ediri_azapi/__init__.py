# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .get_resource import *
from .get_resource_action import *
from .provider import *
from .resource import *
from .resource_action import *
from .update_resource import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import ediri_azapi.config as __config
    config = __config
else:
    config = _utilities.lazy_import('ediri_azapi.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "azapi",
  "mod": "index/resource",
  "fqn": "ediri_azapi",
  "classes": {
   "azapi:index/resource:Resource": "Resource"
  }
 },
 {
  "pkg": "azapi",
  "mod": "index/resourceAction",
  "fqn": "ediri_azapi",
  "classes": {
   "azapi:index/resourceAction:ResourceAction": "ResourceAction"
  }
 },
 {
  "pkg": "azapi",
  "mod": "index/updateResource",
  "fqn": "ediri_azapi",
  "classes": {
   "azapi:index/updateResource:UpdateResource": "UpdateResource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "azapi",
  "token": "pulumi:providers:azapi",
  "fqn": "ediri_azapi",
  "class": "Provider"
 }
]
"""
)