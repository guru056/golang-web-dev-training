package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{
		"Gandhi",
		"Be the change",
		true,
	}

	u2 := user{
		"Buddha",
		"The belief of no beliefs",
		false,
	}

	u3 := user{
		"",
		"Nobody",
		true,
	}

	data := []user{u1, u2, u3}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
