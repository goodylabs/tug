/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var envDir string

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if envDir == "" && len(args) > 0 {
			envDir = args[0]
		}

		err := container.Invoke(func(dockerUseCase *application.DockerUseCase) {
			dockerUseCase.Execute(envDir)
		})
		if err != nil {
			cmd.PrintErrf("Error executing docker command: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().StringVar(&envDir, "envDir", "", "Environment directory name (alternative to positional argument)")
}
