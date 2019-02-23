package main

import (
	"github.com/caius/terraform-provider-smartos/smartos"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: smartos.Provider,
	})
}
