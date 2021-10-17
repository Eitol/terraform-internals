package main

import (
	"github.com/Eitol/terraform-internals/pkg/grpcwrap"
	"github.com/Eitol/terraform-internals/pkg/plugin"
	simple "github.com/Eitol/terraform-internals/pkg/provider-simple"
	"github.com/Eitol/terraform-internals/pkg/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
