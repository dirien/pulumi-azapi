---
title: Azapi Setup
meta_desc: Information on how to install the Azapi provider.
layout: installation
---

## Installation

The Pulumi Azapi provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ediri/azapi`](https://www.npmjs.com/package/@ediri/azapi)
* Python: [`ediri_azapi`](https://pypi.org/project/ediri_azapi/)
* Go: [`github.com/dirien/pulumi-azapi/sdk/go/azapi`](https://github.com/dirien/pulumi-azapi/sdk)
* .NET: [`ediri.Azapi`](https://www.nuget.org/packages/ediri.Azapi)

### Provider Binary

The Azapi provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource azapi v0.1.7
```

Replace the version string with your desired version.
