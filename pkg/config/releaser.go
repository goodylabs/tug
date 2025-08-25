package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goodylabs/releaser"
	releaserGithub "github.com/goodylabs/releaser/providers/github"
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

		releaserInstance = releaser.ConfigureGithubApp(
			tugDir,
			&releaserGithub.GithubOpts{
				User: "goodylabs",
				Repo: "tug",
			})
	}
	return releaserInstance
}
