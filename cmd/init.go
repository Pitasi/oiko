package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"bufio"
	"os"
	"strings"
	"runtime"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	"github.com/matteojoliveau/quicken/utils"
	"github.com/matteojoliveau/oiko/core/structures"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project. The command will prompt for some informations like name, namespace, license and starting version.`,
	Run: func(cmd *cobra.Command, args []string) {
		existent, exErr := utils.IsFileExistent("Oikofile")
		if exErr != nil {
			log.Fatal(exErr)
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
			log.Fatal(err)
		} else {
			fileErr := ioutil.WriteFile("Oikofile", out, 644)
			if fileErr != nil {
				log.Fatal(fileErr)
			} else {
				fmt.Println()
				fmt.Println("Created 'Oikofile'. Use 'oiko help' to list all available commands.")
			}
		}
	},
}

func prompt() structures.Oikofile {
	in := bufio.NewReader(os.Stdin)
	var nl string
	if runtime.GOOS == "windows" {
		nl = "\r\n"
	} else {
		nl = "\n"
	}
	fmt.Print("Project Name: ")
	pName, _ := in.ReadString('\n')
	pName = strings.Replace(pName, nl, "", -1)
	fmt.Print("Project Namespace (eg: github.com/matteojoliveau/oiko: ")
	nspace, _ := in.ReadString('\n')
	nspace = strings.Replace(nspace, nl, "", -1)
	fmt.Print("Starting Version: ")
	version, _ := in.ReadString('\n')
	version = strings.Replace(version, nl, "", -1)
	fmt.Print("Your Name: ")
	oName, _ := in.ReadString('\n')
	oName = strings.Replace(oName, nl, "", -1)
	fmt.Print("Your Email: ")
	email, _ := in.ReadString('\n')
	email = strings.Replace(email, nl, "", -1)
	fmt.Print("License: ")
	license, _ := in.ReadString('\n')
	license = strings.Replace(license, nl, "", -1)

	info := structures.Oikofile{
		ProjectName: pName,
		Namespace:   nspace,
		Version:     version,
		Owner:       oName,
		Email:       email,
		License:     license,
	}

	return info
}
