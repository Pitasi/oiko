package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"bufio"
	"os"
	"strings"
	"runtime"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/matteojoliveau/oiko/util"
	"github.com/matteojoliveau/oiko/core/structures"
	"path/filepath"
	"github.com/matteojoliveau/oiko/core/config"
)

func init() {
	RootCmd.AddCommand(initCmd)
	config.InitLogger(IsDebug)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project. The command will prompt for some informations like name, namespace, license and starting version.`,
	Run: func(cmd *cobra.Command, args []string) {
		existent, exErr := util.IsFileExistent("Oikofile")
		if exErr != nil {
			log.Error(exErr)
			os.Exit(1)
		}

		if existent {
			fmt.Print("There is already an Oikofile in this directory. Exiting...")
			os.Exit(0)
		}
		fmt.Println("Initializing new Go project. You will be asked for some informations.")
		fmt.Println()

		of := prompt()

		out, err := yaml.Marshal(of)
		if err != nil {
			log.Error(err)
		} else {
			fileErr := ioutil.WriteFile("Oikofile", out, 644)
			if fileErr != nil {
				log.Error(fileErr)
			} else {
				fmt.Println()
				fmt.Println("Created 'Oikofile'. Use 'oiko help' to list all available commands.")
			}
		}
		curr, absErr := filepath.Abs("./")
		if absErr != nil {
			log.Error(absErr)
		} else {
			pkgDir := curr + "/src/" + of.Namespace
			mkdirErr := os.MkdirAll(pkgDir, os.ModePerm)
			if mkdirErr != nil {
				log.Error(absErr)
			}

			cErr := createMain(pkgDir)
			if cErr != nil {
				log.Error(cErr)
			}
		}
	},
}

func prompt() structures.Oikofile {
	in := bufio.NewReader(os.Stdin)
	var nl string
	var goexe string
	if runtime.GOOS == "windows" {
		nl = "\r\n"
		goexe = ".exe"
	} else {
		nl = "\n"
		goexe = ""
	}
	fmt.Print("Project Name: ")
	pName, _ := in.ReadString('\n')
	pName = strings.Replace(pName, nl, "", -1)
	fmt.Print("Starting Version: ")
	version, _ := in.ReadString('\n')
	version = strings.Replace(version, nl, "", -1)
	fmt.Print("Your Name: ")
	oName, _ := in.ReadString('\n')
	oName = strings.Replace(oName, nl, "", -1)
	oNameTrimmed := simplifyString(oName)
	pNameTrimmed := simplifyString(pName)
	defNspace := fmt.Sprintf("github.com/%s/%s", oNameTrimmed, pNameTrimmed)
	fmt.Printf("Project Namespace (default: %s): ", defNspace)
	nspace, _ := in.ReadString('\n')
	nspace = strings.Replace(nspace, nl, "", -1)
	if nspace == "" {
		nspace = defNspace
	}

	fmt.Print("Your Email: ")
	email, _ := in.ReadString('\n')
	email = strings.Replace(email, nl, "", -1)
	fmt.Print("License: ")
	license, _ := in.ReadString('\n')
	license = strings.Replace(license, nl, "", -1)
	fmt.Printf("Executable Name (default: %s): ", pNameTrimmed)
	exe, _ := in.ReadString('\n')
	exe = strings.Replace(exe, nl, "", -1)
	if exe == "" {
		exe = pNameTrimmed + goexe
	}
	if !strings.HasSuffix(exe, goexe) {
		exe = exe + goexe
	}

	info := structures.Oikofile{
		ProjectName: pName,
		Namespace:   nspace,
		Version:     version,
		Owner:       oName,
		Email:       email,
		License:     license,
		Exe:         exe,
	}

	return info
}

func simplifyString(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	return s
}

func createMain(pkgDir string) error {
	file, cErr := os.Create(pkgDir + "/" + "main.go")
	if cErr != nil {
		return cErr
	}

	_, fErr := file.WriteString(`package main

func main() {
	//Your Code Here
}
`)
	if fErr != nil {
		return fErr
	}
	return nil
}
