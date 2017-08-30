package core

import (
	"github.com/matteojoliveau/oiko/core/structures"
	"os/exec"
	"os"
)

type Installer struct {

}

func (i *Installer) Install(oikofile structures.Oikofile) error  {
	ns := oikofile.Namespace
	prj := oikofile.ProjectName
	prepErr := PrepareEnvironment()
	if prepErr != nil {
		return prepErr
	}

	insErr := install(ns, prj)
	if insErr != nil {
		return insErr
	}
	os.Setenv("GOPATH", oldGopath)
	return nil
}

func install(ns string, prj string) error {
	log.Infof("Installing project %s in namespace %s", prj, ns)
	cmd := exec.Command("go", "install", ns)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Info(string(out))

	return nil
}

func NewInstaller() *Installer {
	i := Installer{}
	return &i
}
