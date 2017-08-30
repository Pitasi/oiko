package core

import (
	"runtime"
	"github.com/matteojoliveau/oiko/core/structures"
	"os"
	"strings"
	"path/filepath"
	"fmt"
	"go/build"
	"github.com/matteojoliveau/oiko/util"
)

func SetExeExtension(exe *string) {

	goexe := ""
	if runtime.GOOS == "windows" {
		goexe = ".exe"
	}
	*exe += goexe
}


func Install(oikofile structures.Oikofile) error {
	installer := NewInstaller()
	err := installer.Install(oikofile)
	if err != nil {
		return err
	}
	return nil
}

func Build(oikofile structures.Oikofile) error {
	builder := NewBuilder()
	err := builder.Build(oikofile)
	if err != nil {
		return err
	}
	return nil
}

func PrepareEnvironment() error {
	existent, exErr := util.IsFileExistent("build/")
	if exErr != nil {
		return exErr
	} else {
		if !existent {
			os.MkdirAll(buildDir, os.ModePerm)
			log.Debug("Created", buildDir)
		}
	}

	gopath := build.Default.GOPATH

	paths := strings.Split(gopath, string(os.PathListSeparator))
	ng, absErr := filepath.Abs(newGopath)
	if absErr != nil {
		return absErr
	}

	alreadySet := false
	for i := range paths {
		p := paths[i]
		if p == ng {
			alreadySet = true
		}
	}
	if !alreadySet {
		p := fmt.Sprintf("%s%s%s", gopath, string(os.PathListSeparator), ng)

		os.Setenv("GOPATH", p)
		log.Debug("Updated GOPATH with ", ng)
		log.Debug("New GOPATH: ", p)
	}

	return nil
}

func Check(err error) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
