package main

import (
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/luebken/stpl/pkg/stpl/maven"

	"fmt"
	"log"
)

func main() {

	//hackish import reference stacks from effective poms

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		log.Printf("usage: stplcli pom2referencestack <FILENAME>")
		os.Exit(1)
	}
	filename := argsWithoutProg[1]
	//	log.Printf("parsing : %v\n", filename)

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	projects := maven.Unmarshal(dat)

	fmt.Printf("\nvar %v = `\n", firstToLower(strings.Replace(projects[0].Name, " ", "", -1)))
	fmt.Printf("name: %v\n", projects[0].Name)
	descNoTabs := strings.Replace(projects[0].Description, "\t", "", -1)
	descNoNewLines := strings.Replace(descNoTabs, "\n", " ", -1)
	fmt.Printf("description: %#v\n", descNoNewLines)
	fmt.Print("depdendencies: \n")

	for _, dep := range projects[0].Dependencies.Dependencies {
		if dep.Scope == "compile" {
			fmt.Printf("- %v:%v:%v\n", dep.GroupID, dep.ArtifactID, dep.VersionString)
		}
	}
	fmt.Printf("`\n")
}

func firstToLower(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
