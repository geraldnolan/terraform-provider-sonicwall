package main

import (
	"github.com/geraldnolan/terraform-provider-sonicwall/sonicwall"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sonicwall.Provider,
	})
}
