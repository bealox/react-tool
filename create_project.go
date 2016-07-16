package main

import (
	"io/ioutil"
	"log"
	"os"
)

var readme = `
#React-tool

You might want to install ESLint to debug your React files, run this  to install it.
npm install eslint eslint-plugin-react --save-dev

React-tool is a command line that allows you to create project with React js.
-cp=<name> to create project
-cf=<name> to create files

#Create projects
The current folder structure is like this:
- folder
	public/
		-css/
		-js/
		index.html
		index.jsx
	server.go

#Create files
this command line creates jsx and html files.
`

var server = `
	package main

	import (
		"log"
		"net/http"
	)

	func main() {
		http.Handle("/", http.FileServer(http.Dir("./public")))
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
`

var eslint = `
{
    "env": {
        "browser": true,
        "node": true
    },
    "globals": {
        "React":true
    },
    "ecmaFeatures":{
	"jsx": true
    },
    "parser": "babel-eslint",
    "plugins": [
        "react"
    ],
    "rules": {
	        "strict": 0
     }
}
`

func createProject(fileName string) {
	/**
	  Create a project that contains js and css subfolders.
	**/
	err := os.Mkdir(fileName, 0777)

	if err != nil {
		log.Fatalln(err)
	}

	err = os.Chdir(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	//readme
	err = ioutil.WriteFile("README.md", []byte(readme), 0777)
	if err != nil {
		log.Fatalln(err)
	}

	//generate server.go
	err = copyServerFile()
	if err != nil {
		log.Fatalf("unable to copy file %s", err)
	}

	err = os.Mkdir("public", 0777)

	if err != nil {
		log.Fatalln(err)
	}

	err = os.Chdir("public")

	if err != nil {
		log.Fatalln(err)
	}

	//Install stuff here

	err = ioutil.WriteFile(".eslintrc.json", []byte(eslint), 0777)
	if err != nil {
		log.Fatalln("issue creating eslint json file ", err)
	}

	err = os.Mkdir("js", 0777)
	err = os.Mkdir("css", 0777)
	createFile("index")

	if err != nil {
		log.Fatalln("issue creating css/js folder ", err)
	}

}

func copyServerFile() error {
	err := ioutil.WriteFile("server.go", []byte(server), 0777)
	return err
}
