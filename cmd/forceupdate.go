/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/spf13/cobra"
)

var forceupdateCmd = &cobra.Command{
	Use:   "forceupdate",
	Short: "Force app update",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.GetReleaser().ForceUpdate(); err != nil {
			fmt.Println("Error updating application:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(forceupdateCmd)
}
