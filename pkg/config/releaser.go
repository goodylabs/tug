package config

import (
	"fmt"
	"os"
	"path/filepath"

	releaserApi "github.com/goodylabs/releaser/api"
	releaser "github.com/goodylabs/releaser/releaser"
)

var releaserInstance *releaser.ReleaserInstance

func GetReleaser() *releaser.ReleaserInstance {
	if releaserInstance == nil {
		homeDir, err := os.UserHomeDir()
		tugDir := filepath.Join(homeDir, ".tug")
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			os.Exit(1)
		}

		opts := &releaserApi.GithubAppOpts{
			User: "goodylabs",
			Repo: "tug",
		}

		releaserInstance = releaserApi.ConfigureGithubApp(tugDir, opts)
	}
	return releaserInstance
}
