/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/pkg/dependecies"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure tug",
	Run: func(cmd *cobra.Command, args []string) {
		container := dependecies.InitDependencyContainer()
		err := container.Invoke(func(configureUseCase *app.ConfigureUseCase) error {
			return configureUseCase.Execute()
		})
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
