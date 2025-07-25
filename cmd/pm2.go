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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := container.Invoke(func(pm2UseCase *application.Pm2UseCase) {
			pm2UseCase.Execute()
		})
		if err != nil {
			cmd.PrintErrf("Error executing docker command: %v\n", err)
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
