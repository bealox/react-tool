package main

import (
	"flag"
	"fmt"
	"os"
)

type htmlcontent struct {
	Title      string
	Javascript []string
	JSX        []string
}

//Usage
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
}

var cp string // create project
var cf string //create file

func init() {
	flag.StringVar(&cp, "create_project", "", "-create_project=<string>, create a project")
	flag.StringVar(&cp, "cp", "", "-cp=<string>, create a project..")
	flag.StringVar(&cf, "create_file", "", "-create_file=<value>, create files")
	flag.StringVar(&cf, "cf", "", "-cp=<string>, create files")
}

func main() {

	flag.Parse()

	if cf != "" {
		createFile(cf)
	} else if cp != "" {
		createProject(cp)
	} else {
		flag.Usage()
	}

}
