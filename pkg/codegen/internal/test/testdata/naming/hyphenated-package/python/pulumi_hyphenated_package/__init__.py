# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .resource import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "hyphenated-package",
  "mod": "",
  "fqn": "pulumi_hyphenated_package",
  "classes": {
   "hyphenated-package::Resource": "Resource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "hyphenated-package",
  "token": "pulumi:providers:hyphenated-package",
  "fqn": "pulumi_hyphenated_package",
  "class": "Provider"
 }
]
"""
)
