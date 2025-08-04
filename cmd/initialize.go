/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal"
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

// initializeCmd represents the initialize command
var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "A brief description of your command",
	Long:  `Initialize configuration for tug`,
	Run: func(cmd *cobra.Command, args []string) {
		container := internal.InitDependencyContainer()
		err := container.Invoke(func(initializeUseCase *application.InitializeUseCase) error {
			return initializeUseCase.Execute()
		})
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initializeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initializeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initializeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
