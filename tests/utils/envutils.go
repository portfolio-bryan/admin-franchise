package utilstests

import (
	"os"

	"github.com/bperezgo/admin_franchise/config"
)

// LoadEnv load the env files relative to the root folder from any nested test directory.
func LoadEnv() {
	os.Setenv("ENVIRONMENT", "test")
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
}
