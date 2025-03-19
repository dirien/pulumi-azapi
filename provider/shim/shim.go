package shim

import (
	"github.com/Azure/terraform-provider-azapi/internal/provider"
	tf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider() tf.Provider {
	return provider.AzureProvider()
}
