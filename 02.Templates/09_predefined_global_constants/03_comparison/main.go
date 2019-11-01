package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type score struct {
	Score1 int
	Score2 int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := score{
		9,
		7,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
