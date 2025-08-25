package cmd

import (
	"fmt"
	"os"

	"github.com/goodylabs/tug/pkg/config"
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
	updated, err := config.GetReleaser().Run()
	if err != nil {
		fmt.Println("Error checking for updates:", err)
	} else if updated {
		fmt.Println("Application has been updated.")
		os.Exit(0)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version information")
}
