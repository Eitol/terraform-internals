package terraform

import (
	backendInit "github.com/Eitol/terraform-internals/pkg/backend/init"
)

func init() {
	// Initialize the backends
	backendInit.Init(nil)
}
