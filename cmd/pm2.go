/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

// pm2Cmd represents the pm2 command
var pm2Cmd = &cobra.Command{
	Use:   "pm2",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := container.Invoke(func(pm2UseCase *application.Pm2UseCase) error {
			return pm2UseCase.Execute()
		})
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pm2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pm2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pm2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
