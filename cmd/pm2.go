package cmd

import (
	"github.com/goodylabs/tug/internal"
	"github.com/goodylabs/tug/internal/application"
	"github.com/spf13/cobra"
)

var pm2Cmd = &cobra.Command{
	Use:   "pm2 <envDir>",
	Short: "Abstraction layer for pm2 operations related to project repo",
	Run: func(cmd *cobra.Command, args []string) {
		container := internal.InitDependencyContainer("pm2")
		err := container.Invoke(func(GenericUseCase *application.GenericUseCase) error {
			return GenericUseCase.Execute()
		})
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pm2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pm2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pm2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
