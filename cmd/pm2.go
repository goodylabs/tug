package cmd

import (
	"github.com/goodylabs/tug/internal"
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var pm2Cmd = &cobra.Command{
	Use:   "pm2 <envDir>",
	Short: "Abstraction layer for pm2 operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		check, err := cmd.Flags().GetBool("check")

		if check {
			container := internal.InitDependencyContainer("pm2")
			err = container.Invoke(func(checkConnectionUseCase *application.CheckConnectionUseCase) error {
				return checkConnectionUseCase.Execute()
			})
		} else {
			container := internal.InitDependencyContainer("pm2")
			err = container.Invoke(func(GenericUseCase *application.GenericUseCase) error {
				return GenericUseCase.Execute()
			})
		}
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pm2Cmd)
	pm2Cmd.Flags().Bool("check", false, "Check SSH connections before running PM2 commands")
}
