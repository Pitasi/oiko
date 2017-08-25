package cmd

import (
	"github.com/spf13/cobra"
)


var IsDebug bool
func init() {
	//Must find a way to set debug loggin level
	//RootCmd.PersistentFlags().BoolVarP(&IsDebug,"debug", "d", false, "Set the logging level to debug")
}

var RootCmd = &cobra.Command{
	Use:   "oiko",
	Short: "Oiko",
	Long:  `Oiko is a simple declarative build system for Go projects`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
