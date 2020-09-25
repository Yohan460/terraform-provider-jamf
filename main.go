package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/sioncojp/terraform-provider-jamf/jamf"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jamf.Provider,
	})
}