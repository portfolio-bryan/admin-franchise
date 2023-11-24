package utilstests

import (
	"os"

	"github.com/bperezgo/admin_franchise/config"
)

// this should match the project root folder name.
const rootFolder = "assessment-cc-go-sr-01"

// LoadEnv load the env files relative to the root folder from any nested test directory.
func LoadEnv() {
	os.Setenv("ENVIRONMENT", "test")
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
}
