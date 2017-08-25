package core

import (
	"github.com/matteojoliveau/oiko/util"
	"github.com/matteojoliveau/oiko/core/structures"
	"os"
	"go/build"
	"strings"
	"path/filepath"
	"fmt"
	"os/exec"
	"time"
	"github.com/matteojoliveau/oiko/core/config"
)
var log = config.Logger
var buildDir = "build/"
var srcDir = "src/"
var newGopath = "."

type Builder struct {
}

func (b *Builder) Build(oikofile structures.Oikofile) error {
	log.Debug("Test")
	ns := oikofile.Namespace
	exe := oikofile.Exe
	prepErr := b.prepareEnvironment(ns)
	if prepErr != nil {
		return prepErr
	}

	compErr := b.compile(ns, exe)
	if compErr != nil {
		return compErr
	}

	return nil
}

func (b *Builder) prepareEnvironment(namespace string) error {
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

func (b *Builder) compile(namespace string, exe string) error {
	t0 := time.Now()
	srcPath := fmt.Sprintf("%s%s", srcDir, namespace)
	log.Info("Compiling sources from ", srcPath)
	exePath := fmt.Sprintf("%s/%s", buildDir, exe)
	cmd := exec.Command("go", "build", "-o", exePath, namespace)
	out, err := cmd.CombinedOutput()
	log.Debug(string(out))
	if err != nil {
		return err
	}
	log.WithField("compilation_time", time.Since(t0)).Infof("Built '%s' in %s", exe, buildDir)

	return nil
}

func NewBuilder() Builder {
	b := Builder{}
	return b
}
