package unit

import (
	"go-cli-mgt/pkg/server"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	server.Initialize()

	os.Exit(m.Run())
}
