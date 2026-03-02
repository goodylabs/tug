package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/internal/modules/action"
	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/spf13/cobra"
)

var pm2Cmd = &cobra.Command{
	Use:   "pm2",
	Short: "Abstraction layer for pm2 operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		checkConnectionUseCase := app.NewCheckConnectionUseCase()
		useModuleUseCase := app.NewUseModuleV2UseCase()

		if check, _ := cmd.Flags().GetBool("check"); check == true {
			if err := checkConnectionUseCase.Execute(loadproject.Pm2Strategy); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		if host, _ := cmd.Flags().GetString("host"); host != "" {
			user, _ := cmd.Flags().GetString("user")
			if err := useModuleUseCase.ExecuteDirect(user, host, action.Pm2); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		if err := useModuleUseCase.Execute(loadproject.Pm2Strategy, action.Pm2); err != nil {
			cmd.PrintErrf("%v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(pm2Cmd)
	pm2Cmd.Flags().Bool("check", false, checkHint)
	pm2Cmd.Flags().String("host", "", customHostHint)
	pm2Cmd.Flags().String("user", "root", customUserHint)
}
