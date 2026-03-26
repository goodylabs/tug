package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/internal/modules/action"
	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Abstraction layer for docker swarm operations related to project repository",
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
			if err := useModuleUseCase.ExecuteDirect(user, host, action.Swarm); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
			return
		}

		if err := useModuleUseCase.Execute(loadproject.DockerStrategy, action.Swarm); err != nil {
			cmd.PrintErrf("%v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(swarmCmd)
	swarmCmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
	swarmCmd.Flags().String("host", "", customHostHint)
	swarmCmd.Flags().String("user", "root", customUserHint)
}
