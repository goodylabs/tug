package cmd

import (
	"fmt"
	"os"

	"github.com/goodylabs/tug/internal"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var version = "default"

var container *dig.Container

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
	if container == nil {
		container = internal.InitDependencyContainer()
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version information")
}
