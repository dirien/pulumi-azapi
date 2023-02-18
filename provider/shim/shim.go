package shim

import (
	"github.com/Azure/terraform-provider-azapi/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	p := provider.AzureProvider()
	return p
}
