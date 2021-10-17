package local

import (
	"flag"
	"os"
	"testing"

	_ "github.com/Eitol/terraform-internals/pkg/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
