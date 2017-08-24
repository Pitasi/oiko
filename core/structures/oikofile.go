package structures

import (
	"path/filepath"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

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
			log.Fatal(err)
		}
		oikoPath = abs
	}
	data, ioerr := ioutil.ReadFile(oikoPath)
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	of := Oikofile{}
	err := yaml.Unmarshal([]byte(data), &of)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return of
}
