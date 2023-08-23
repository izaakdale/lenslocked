package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(path string) (Template, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template error: %w", err)
	}
	return Template{t}, nil
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "template execution error", http.StatusInternalServerError)
		return
	}
}
