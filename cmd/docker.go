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
		checkConnectionUseCase := app.NewCheckConnectionUseCase()
		useModuleUseCase := app.NewUseModuleV2UseCase()

		if check, _ := cmd.Flags().GetBool("check"); check == true {
			if err := checkConnectionUseCase.Execute(loadproject.DockerStrategy); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		if host, _ := cmd.Flags().GetString("host"); host != "" {
			user, _ := cmd.Flags().GetString("user")
			if err := useModuleUseCase.ExecuteDirect(user, host, action.Docker); err != nil {
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
	dockerCmd.Flags().Bool("check", false, checkHint)
	dockerCmd.Flags().String("host", "", customHostHint)
	dockerCmd.Flags().String("user", "root", customUserHint)
}
