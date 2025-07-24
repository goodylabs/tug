package cmd

import (
	"fmt"
	"os"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tug",
	Short: "CLI tool to manage Docker Swarm environments",
	Long: `Tug is a command-line interface tool designed to simplify Docker Swarm management.
It provides easy access to remote Docker containers and environments through
automated deployment scripts and container selection.

Use 'tug --help' to see available commands and options.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to tug! Use --help to see available commands.")
	},
}

func Execute() {
	adapters.InitializeDependencies()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Toggle verbose output")
}
