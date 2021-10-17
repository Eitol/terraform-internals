package etcdv2

import (
	"testing"

	"github.com/Eitol/terraform-internals/pkg/backend"
)

func TestBackend_impl(t *testing.T) {
	var _ backend.Backend = new(Backend)
}
