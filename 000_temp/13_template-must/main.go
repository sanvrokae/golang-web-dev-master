package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "indexfiles.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "about.gohtml", nil)
}
