package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Error creating template", err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating new file", err)
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
}
