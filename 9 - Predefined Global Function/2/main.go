package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Mcleod",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
