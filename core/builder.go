package core

import (
	"github.com/matteojoliveau/oiko/util"
	"os"
	"github.com/matteojoliveau/oiko/core/structures"
	"go/build"
	"strings"
	"path/filepath"
	"fmt"
)

var newGopath = "build/go"
var buildDir = newGopath + "src/"

type Builder struct {
}

func (b *Builder) Build(oikofile structures.Oikofile) error{
	prepErr := b.prepareEnvironment(oikofile.Namespace)
	if prepErr != nil {
		return prepErr
	}

	return nil
}

func (b *Builder) prepareEnvironment(namespace string) error {
	existent, exErr := util.IsFileExistent("build/")
	if exErr != nil {
		return exErr
	} else {
		if !existent {
			os.MkdirAll(buildDir+namespace, os.ModePerm)
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
		os.Setenv("GOPATH", fmt.Sprintf("%s%s%s", gopath, os.PathListSeparator, ng))
	}

	return nil
}



func NewBuilder() Builder {
	b := Builder{}
	return b
}
