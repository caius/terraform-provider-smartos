package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-smartos/smartos"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: smartos.Provider,
	})
}
