package main

import (
	"github.com/Eitol/terraform-internals/pkg/grpcwrap"
	plugin "github.com/Eitol/terraform-internals/pkg/plugin6"
	simple "github.com/Eitol/terraform-internals/pkg/provider-simple-v6"
	"github.com/Eitol/terraform-internals/pkg/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
