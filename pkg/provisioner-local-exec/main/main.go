package main

import (
	localexec "github.com/Eitol/terraform-internals/pkg/builtin/provisioners/local-exec"
	"github.com/Eitol/terraform-internals/pkg/grpcwrap"
	"github.com/Eitol/terraform-internals/pkg/plugin"
	"github.com/Eitol/terraform-internals/pkg/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
