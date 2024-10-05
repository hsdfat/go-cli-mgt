package unit

import (
	"go-cli-mgt/pkg/server"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup env, should be absolute path
	server.Initialize(".env")

	os.Exit(m.Run())
}
