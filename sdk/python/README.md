# Azapi Resource Provider

The Azapi Resource Provider lets you manage [azapi](https://learn.microsoft.com/en-us/azure/developer/terraform/overview-azapi-provider) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @ediri/azapi
```

or `yarn`:

```bash
yarn add @ediri/azapi
```

### Python

To use from Python, install using `pip`:

```bash
pip install ediri_azapi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/dirien/pulumi-azapi/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package ediri.Azapi
```

## Configuration

The following configuration points are available for the `azapi` provider:

- `azapi:apiKey` (environment: `azapi_API_KEY`) - the API key for `azapi`
- `azapi:region` (environment: `azapi_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/azapi/api-docs/).
