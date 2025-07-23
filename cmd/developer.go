/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var envDir string

var developerCmd = &cobra.Command{
	Use:   "developer [envDir]",
	Short: "Run developer command with optional env dir",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			envDir = args[0]
		}

		application.DeveloperUseCase(&application.DeveloperOptions{
			EnvDir: envDir,
		})
	},
}

func init() {
	rootCmd.AddCommand(developerCmd)

	developerCmd.Flags().StringVar(&envDir, "envDir", "", "Path to config file")
}
