package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goodylabs/releaser"
	releaserGithub "github.com/goodylabs/releaser/providers/github"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "tug",
	Short: "CLI tool to manage Docker Swarm environments",
	Long:  `Use 'tug --help' to see available commands and options.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Println(version)
			return
		}

		fmt.Println("Welcome to tug! Use --help to see available commands.")
	},
}

func Execute() {
	instance := releaser.ConfigureGithubApp(&releaserGithub.GithubOpts{
		User: "goodylabs",
		Repo: "tug",
	})

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}
	tugDir := filepath.Join(homeDir, ".tug")
	if err = instance.Run(tugDir); err != nil {
		fmt.Println("Error checking for updates:", err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version information")
}
