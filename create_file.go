package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func createFile(cf string) {
	js := cf + ".jsx"
	html := cf + ".html"

	attributes := htmlcontent{
		Title: "Test",
		JSX:   []string{js},
	}

	err := scanJS(&attributes)
	if err != nil {
		log.Fatalln("error when opening a file ", err)
	}

	ioutil.WriteFile(js, nil, 0777)
	ioutil.WriteFile(html, nil, 0777)

	file, err := os.OpenFile(html, os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("error when opening a file ", err)
	}
	tpl, err := template.New("test").Parse(content)
	if err != nil {
		log.Fatalln("error when parsing template ", err)
	}
	err = tpl.Execute(file, attributes)
	if err != nil {
		log.Fatalln("erro when template execute ", err)
	}
}

func scanJS(attribute *htmlcontent) error {

	return filepath.Walk("js/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)

		if strings.EqualFold(ext, ".js") {
			attribute.Javascript = append(attribute.Javascript, path)
		} else if strings.EqualFold(ext, ".jsx") {
			attribute.JSX = append(attribute.JSX, path)
		}

		return nil
	})
}
