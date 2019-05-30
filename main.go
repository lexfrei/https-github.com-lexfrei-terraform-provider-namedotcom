package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mhumesf/terraform-provider-namedotcom/namedotcom"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}
