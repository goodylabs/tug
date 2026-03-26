/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/internal/modules/action"
	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/spf13/cobra"
)

// pystranoCmd represents the pystrano command
var pystranoCmd = &cobra.Command{
	Use:   "pystrano",
	Short: "Abstraction layer for pystrano operations related to project repository",
	Run: func(cmd *cobra.Command, args []string) {

		if check, _ := cmd.Flags().GetBool("check"); check == true {
			checkConnectionUseCase := app.NewCheckConnectionUseCase()
			if err := checkConnectionUseCase.Execute(loadproject.PystranoStrategy); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		useCase := app.NewUseModuleV2UseCase()
		err := useCase.Execute(loadproject.PystranoStrategy, action.Pystrano)

		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pystranoCmd)
	pystranoCmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
}
