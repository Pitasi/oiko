package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "oiko",
	Short: "Oiko",
	Long:  `Oiko is a simple declarative build system for Go projects`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
