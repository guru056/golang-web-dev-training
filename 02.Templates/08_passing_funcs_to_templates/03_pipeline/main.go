package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"fsqrt":   sqRoot,
	"fdbl":    double,
	"fsquare": square,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return float64(x * x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
