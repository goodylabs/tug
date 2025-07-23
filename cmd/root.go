package cmd

import (
	"fmt"
	"os"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tug",
	Short: "CLI tool to manage Docker Swarm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Witaj w tug! Użyj --help, aby zobaczyć dostępne komendy.")
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
