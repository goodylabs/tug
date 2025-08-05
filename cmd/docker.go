package cmd

import (
	"github.com/goodylabs/tug/internal"
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Abstraction layer for docker operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		check, err := cmd.Flags().GetBool("check")

		container := internal.InitDependencyContainer(
			internal.WithDockerHandler,
		)
		if check {
			err = container.Invoke(func(checkConnectionUseCase *application.CheckConnectionUseCase) error {
				return checkConnectionUseCase.Execute()
			})
		} else {
			err = container.Invoke(func(useModuleUseCase *application.UseModuleUseCase) error {
				return useModuleUseCase.Execute()
			})
		}
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().Bool("check", false, "Check SSH connections before running Docker commands")
}
