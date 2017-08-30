package cmd

import (
	"github.com/spf13/cobra"
	"github.com/matteojoliveau/oiko/core"
	"github.com/matteojoliveau/oiko/core/structures"
	"github.com/matteojoliveau/oiko/core/config"
)



func init() {
	RootCmd.AddCommand(buildCmd)
	config.InitLogger(IsDebug)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Long:  `Build the project. It will compile all the sources from 'src/' and put the output in 'build/'`,
	Run: func(cmd *cobra.Command, args []string) {

		file := structures.ReadOikoFile()
		core.SetExeExtension(&file.Exe)
		buildErr := core.Build(file)
		core.Check(buildErr)
	},
}

