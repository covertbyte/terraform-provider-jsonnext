package main

import (
	"github.com/covertbyte/terraform-provider-jsonnext/jsonnext"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jsonnext.Provider,
	})
}
