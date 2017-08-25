package src

import (
	"github.com/matteojoliveau/oiko/cmd"
	"os"

	"github.com/matteojoliveau/oiko/core/config"
)

var log = config.Logger

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
