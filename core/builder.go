package core

import (
	"github.com/matteojoliveau/oiko/core/structures"
	"fmt"
	"os/exec"
	"time"
	"github.com/matteojoliveau/oiko/core/config"
	"os"
)
var log = config.Logger
var buildDir = "build/"
var srcDir = "src/"
var newGopath = "."
var oldGopath = os.Getenv("GOPATH")

type Builder struct {
}

func (b *Builder) Build(oikofile structures.Oikofile) error {
	ns := oikofile.Namespace
	exe := oikofile.Exe
	prepErr := PrepareEnvironment()
	if prepErr != nil {
		return prepErr
	}

	compErr := compile(ns, exe)
	if compErr != nil {
		return compErr
	}

	return nil
}



func compile(namespace string, exe string) error {
	t0 := time.Now()
	srcPath := fmt.Sprintf("%s%s", srcDir, namespace)
	log.Info("Compiling sources from ", srcPath)
	exePath := fmt.Sprintf("%s/%s", buildDir, exe)
	cmd := exec.Command("go", "build", "-o", exePath, namespace)
	out, err := cmd.CombinedOutput()
	log.Info(string(out))
	if err != nil {
		return err
	}
	log.WithField("compilation_time", time.Since(t0)).Infof("Built '%s' in %s", exe, buildDir)

	os.Setenv("GOPATH", oldGopath)
	return nil
}

func NewBuilder() *Builder {
	b := Builder{}
	return &b
}
