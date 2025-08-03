package cmd

import (
	"github.com/goodylabs/tug/internal"
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var dockerLongDesc = `Execute Docker-related setup for a given environment.

<envDir> should be the name of an environment directory located under the 'devops/' folder.
For example, if you have a folder 'devops/staging', you should run:

   tug docker staging

This will trigger execution using the configuration from 'devops/staging'.
`

var dockerCmd = &cobra.Command{
	Use:   "docker <envDir>",
	Short: "Run Docker-related commands for a specific environment",
	Long:  dockerLongDesc,
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := internal.InitDependencyContainer("docker")

		err := container.Invoke(func(useModuleUseCase *application.UseModuleUseCase) error {
			return useModuleUseCase.Execute()
		})
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

var useDockerCmd = "true"

func init() {
	if useDockerCmd == "true" {
		rootCmd.AddCommand(dockerCmd)
	}
}
