package main

import (
	"github.com/Eitol/terraform-internals/pkg/builtin/providers/terraform"
	"github.com/Eitol/terraform-internals/pkg/grpcwrap"
	"github.com/Eitol/terraform-internals/pkg/plugin"
	"github.com/Eitol/terraform-internals/pkg/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
