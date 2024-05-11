package shim

import (
	"context"
	"github.com/Azure/terraform-provider-azapi/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	prov "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	p := provider.AzureProvider()
	return p
}

var _ prov.Provider = &AzapiProvider{}

type AzapiProvider struct{}

func (a AzapiProvider) Metadata(ctx context.Context, request prov.MetadataRequest, response *prov.MetadataResponse) {
}

func (a AzapiProvider) Schema(ctx context.Context, request prov.SchemaRequest, response *prov.SchemaResponse) {

}

func (a AzapiProvider) Configure(ctx context.Context, request prov.ConfigureRequest, response *prov.ConfigureResponse) {

}

func (a AzapiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (a AzapiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func Framework() func() prov.Provider {
	return func() prov.Provider {
		return &AzapiProvider{}
	}
}
