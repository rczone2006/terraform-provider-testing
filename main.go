package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/rczone2006/terraform-provider-testing/testing"

	
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: testing.Provider,
        
    })
}