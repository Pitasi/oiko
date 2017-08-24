package cmd

import (
	"github.com/spf13/cobra"
	"github.com/matteojoliveau/oiko/core"
	"github.com/matteojoliveau/oiko/core/structures"
	"log"
	"os"
)

func init()  {
	RootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Long:  `Build the project. It will compile all the sources from 'src/' and put the output in 'build/'`,
	Run: func(cmd *cobra.Command, args []string) {
		file := structures.ReadOikoFile()
		builder := core.NewBuilder()
		buildErr := builder.Build(file)
		if buildErr != nil {
			log.Fatal(buildErr)
			os.Exit(1)
		}
	},
}


