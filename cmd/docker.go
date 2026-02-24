package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/internal/modules/action"
	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Abstraction layer for docker operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		if check, _ := cmd.Flags().GetBool("check"); check == true {
			checkConnectionUseCase := app.NewCheckConnectionUseCase()
			if err := checkConnectionUseCase.Execute(loadproject.DockerStrategy); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		useModuleUseCase := app.NewUseModuleV2UseCase()

		if input, _ := cmd.Flags().GetBool("input"); input == true {
			if err := useModuleUseCase.Execute(loadproject.InputStrategy, action.Docker); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		if err := useModuleUseCase.Execute(loadproject.DockerStrategy, action.Docker); err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
	dockerCmd.Flags().Bool("input", false, "Manually input host configuration")
}
