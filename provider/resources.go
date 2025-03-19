// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	shimprovider "github.com/Azure/terraform-provider-azapi/shim"
	"github.com/dirien/pulumi-azapi/provider/pkg/version"
	"github.com/ettle/strcase"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

//go:embed cmd/pulumi-resource-azapi/bridge-metadata.json
var bridgeMetadata []byte

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	mainMod = "index" // the azapi module
)

var module_overrides = map[string]string{}

var name_overrides = map[string]string{}

func convertName(tfname string) (module string, name string) {
	tfNameItems := strings.Split(tfname, "_")
	contract.Assertf(len(tfNameItems) >= 2, "Invalid snake case name %s", tfname)
	contract.Assertf(tfNameItems[0] == "azapi", "Invalid snake case name %s. Does not start with azapi", tfname)
	// if len(tfNameItems) == 2 {
	// 	module = mainMod
	// 	name = tfNameItems[1]
	// } else {
	// 	module = strcase.ToPascal(strings.Join(tfNameItems[1:len(tfNameItems)-1], "_"))
	// 	name = tfNameItems[len(tfNameItems)-1]

	// 	if v, ok := module_overrides[module]; ok {
	// 		module = v
	// 	}
	// }
	// contract.Assertf(!unicode.IsDigit(rune(module[0])), "Pulumi namespace must not start with a digit: %s", name)
	module = mainMod
	name = strings.Join(tfNameItems[1:], "_")
	name = strcase.ToPascal(name)
	if v, ok := name_overrides[name]; ok {
		name = v
	}
	contract.Assertf(!unicode.IsDigit(rune(name[0])), "Pulumi name must not start with a digit: %s", name)
	return
}

func makeDataSource(ds string) tokens.ModuleMember {
	mod, name := convertName(ds)
	return tfbridge.MakeDataSource("azapi", mod, "get"+name)
}

func makeResource(res string) tokens.Type {
	mod, name := convertName(res)
	return tfbridge.MakeResource("azapi", mod, name)
}

func moduleComputeStrategy() tfbridge.Strategy {
	return tfbridge.Strategy{
		Resource: func(tfToken string, elem *tfbridge.ResourceInfo) error {
			if elem.Tok == "" {
				elem.Tok = makeResource(tfToken)
			}
			return nil
		},
		DataSource: func(tfToken string, elem *tfbridge.DataSourceInfo) error {
			if elem.Tok == "" {
				elem.Tok = makeDataSource(tfToken)
			}
			return nil
		},
	}
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pf.ShimProvider(shimprovider.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "azapi",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "AzAPI",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "dirien",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/dirien/pulumi-azapi/img/azure.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/dirien/pulumi-azapi",
		Description:       "A Pulumi package for creating and managing Azapi resources",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords: []string{
			"pulumi",
			"azapi",
			"category/cloud",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/dirien/pulumi-azapi",
		Repository: "https://github.com/dirien/pulumi-azapi",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		Version:              version.Version,
		GitHubOrg:            "Azure",
		MetadataInfo:         tfbridge.NewProviderMetadata(bridgeMetadata),
		TFProviderVersion:    "2.3.0",
		UpstreamRepoPath:     "./upstream",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources:            map[string]*tfbridge.ResourceInfo{},
		DataSources:          map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@ediri/azapi",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "ediri_azapi",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/dirien/pulumi-%[1]s/sdk/", "azapi"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"azapi",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "ediri",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "io.dirien",
		},
	}

	prov.MustComputeTokens(moduleComputeStrategy())
	prov.SetAutonaming(255, "-")

	return prov
}
