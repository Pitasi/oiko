package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/matteojoliveau/oiko/core/config"
	"github.com/matteojoliveau/oiko/core/structures"
	"github.com/matteojoliveau/oiko/core"
	"io"
)

func init() {
	RootCmd.AddCommand(installCmd)
	config.InitLogger(IsDebug)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Build and install the project",
	Long:  `Build and install the project. It will compile all the sources from 'src/' and install them in the standard GOPATH/bin folder'`,
	Run: func(cmd *cobra.Command, args []string) {

		file := structures.ReadOikoFile()
		core.SetExeExtension(&file.Exe)

		instErr := core.Install(file)
		core.Check(instErr)

		buildErr := core.Build(file)
		core.Check(buildErr)

		copyExe(file)

	},
}

func copyExe(oikofile structures.Oikofile) {
	exe := oikofile.Exe
	log.Infof("Installing %s", exe)
	srcFile, err := os.Open("build/" + exe)
	core.Check(err)
	defer srcFile.Close()

	destFile, err := os.Create(os.Getenv("GOPATH") + "/bin/" + exe) // creates if file doesn't exist
	core.Check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // core.Check first var for number of bytes copied
	core.Check(err)

	err = destFile.Sync()
	core.Check(err)

	log.Info("Installation finished!")
}


