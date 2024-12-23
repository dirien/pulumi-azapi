package shim

import (
	azprovider "github.com/Azure/terraform-provider-azapi/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProviderV2() provider.Provider {
	return azprovider.AzureProvider()
}
