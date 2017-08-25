package cmd

import (
	"github.com/spf13/cobra"
	"github.com/matteojoliveau/oiko/core"
	"github.com/matteojoliveau/oiko/core/structures"
	"os"
	"github.com/matteojoliveau/oiko/core/config"
	"fmt"
)

var log = config.Logger

func init() {
	RootCmd.AddCommand(buildCmd)
	config.InitLogger(IsDebug)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Long:  `Build the project. It will compile all the sources from 'src/' and put the output in 'build/'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(IsDebug)

		file := structures.ReadOikoFile()
		builder := core.NewBuilder()
		buildErr := builder.Build(file)
		if buildErr != nil {
			log.Error(buildErr)
			os.Exit(1)
		}
	},
}
