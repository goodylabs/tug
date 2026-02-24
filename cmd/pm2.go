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
		// check, err := cmd.Flags().GetBool("check")

		// container := dependecies.InitDependencyContainer(
		// 	dependecies.WithPm2Handler,
		// )
		// if check {
		// 	err = container.Invoke(func(checkConnectionUseCase *app.CheckConnectionUseCase) error {
		// 		return checkConnectionUseCase.Execute()
		// 	})
		// } else {
		// 	err = container.Invoke(func(useModuleUseCase *app.UseModuleUseCase) error {
		// 		return useModuleUseCase.Execute()
		// 	})
		// }
		// if err != nil {
		// 	cmd.PrintErrf("%v\n", err)
		// }

		if check, _ := cmd.Flags().GetBool("check"); check == true {
			checkConnectionUseCase := app.NewCheckConnectionUseCase()
			if err := checkConnectionUseCase.Execute(loadproject.Pm2Strategy); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		useCase := app.NewUseModuleV2UseCase()
		err := useCase.Execute(loadproject.Pm2Strategy, action.Pm2)

		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(pm2Cmd)
	pm2Cmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
}
