package cmd

import (
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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envDir := args[0]

		err := container.Invoke(func(dockerUseCase *application.DockerUseCase) error {
			return dockerUseCase.Execute(envDir)
		})
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
