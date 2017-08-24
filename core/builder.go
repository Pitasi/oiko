package core

import (
	"github.com/matteojoliveau/oiko/util"
	"os"
	"github.com/matteojoliveau/oiko/core/structures"
	"go/build"
	"strings"
	"path/filepath"
	"fmt"
	"os/exec"
)

var buildDir = "build/go"
var srcDir = "src/"
var newGopath = srcDir

type Builder struct {
}

func (b *Builder) Build(oikofile structures.Oikofile) error{
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
		}
	}

	gopath := build.Default.GOPATH
	paths := strings.Split(gopath, string(os.PathListSeparator))
	ng, absErr := filepath.Abs(newGopath)
	if absErr != nil{
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
		os.Setenv("GOPATH", fmt.Sprintf("%s%s%s", gopath, string(os.PathListSeparator), ng))
	}

	return nil
}

func (b *Builder) compile(namespace string, exe string) error {
	srcPath := fmt.Sprintf("%s%s", srcDir, namespace)
	fmt.Printf("Comipiling sources from %s\n", srcPath)
	exePath := fmt.Sprintf("%s/%s", buildDir, exe)
	fmt.Println(exePath)
	cmd := exec.Command("go", "build", namespace, "-o", exePath)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}


func NewBuilder() Builder {
	b := Builder{}
	return b
}
