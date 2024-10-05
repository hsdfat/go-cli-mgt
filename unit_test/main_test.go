package unit

import (
	"go-cli-mgt/pkg/server"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup env, should be absolute path
	server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	os.Exit(m.Run())
}
