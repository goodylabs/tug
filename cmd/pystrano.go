/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/pkg/dependecies"
	"github.com/spf13/cobra"
)

// pystranoCmd represents the pystrano command
var pystranoCmd = &cobra.Command{
	Use:   "pystrano",
	Short: "Abstraction layer for pm2 operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		check, err := cmd.Flags().GetBool("check")

		container := dependecies.InitDependencyContainer(
			dependecies.WithPystranoHandler,
		)
		if check {
			err = container.Invoke(func(checkConnectionUseCase *app.CheckConnectionUseCase) error {
				return checkConnectionUseCase.Execute()
			})
		} else {
			err = container.Invoke(func(useModuleUseCase *app.UseModuleUseCase) error {
				return useModuleUseCase.Execute()
			})
		}
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pystranoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pystranoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pystranoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
