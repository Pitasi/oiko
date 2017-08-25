package structures

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/matteojoliveau/oiko/core/config"
)

var log = config.Logger

type Oikofile struct {
	ProjectName string `yaml:"project_name"`
	Namespace   string
	Version     string
	Owner       string
	Email       string
	License     string
	Exe         string `yaml:"executable_name"`
	Vcs struct {
		name string
		url  string
	}
	Dependencies []string
}

func ReadOikoFile() Oikofile {
	oikoPath := "./Oikofile"
	if !filepath.IsAbs(oikoPath) {
		abs, err := filepath.Abs(oikoPath)
		if err != nil {
			log.Error(err)
		}
		oikoPath = abs
	}
	data, ioerr := ioutil.ReadFile(oikoPath)
	if ioerr != nil {
		log.Error(ioerr)
	}
	of := Oikofile{}
	err := yaml.Unmarshal([]byte(data), &of)
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return of
}
