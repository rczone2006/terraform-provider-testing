package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/rczone2006/terraform-provider-testing/zoning"

	
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return zoning.Provider()
        },
    })
}