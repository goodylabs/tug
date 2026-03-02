/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure tug",
	Run: func(cmd *cobra.Command, args []string) {
		configureUseCase := app.NewConfigureUseCase()
		err := configureUseCase.Execute()
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
