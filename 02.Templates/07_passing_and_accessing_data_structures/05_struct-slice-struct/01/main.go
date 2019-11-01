package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
}

func main() {
	buddha := sage{
		"Buddha",
		"The belief of no beliefs",
	}
	gandhi := sage{
		"Gandhi",
		"Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	ford := car{
		"Ford",
		"Mustang",
		4,
	}

	corolla := car{
		"Toyota",
		"Corolla",
		4,
	}

	data := items{
		Wisdom:    []sage{buddha, gandhi, mlk, jesus, muhammad},
		Transport: []car{ford, corolla},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
