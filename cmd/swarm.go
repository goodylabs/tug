package cmd

import (
	"github.com/goodylabs/tug/internal/app"
	"github.com/goodylabs/tug/pkg/dependecies"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Abstraction layer for docker swarm operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		check, err := cmd.Flags().GetBool("check")

		container := dependecies.InitDependencyContainer(
			dependecies.WithSwarmHandler,
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
	rootCmd.AddCommand(swarmCmd)
	swarmCmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
}
