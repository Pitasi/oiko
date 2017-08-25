package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

const ver string = "0.1.0"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Oiko",
	Long:  `All software has versions. This is Oiko's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Oiko Build System\n Version: %s\n", ver)
	},
}
