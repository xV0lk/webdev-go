package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFs(fs fs.FS, patterns ...string) (Template, error) {
	tmpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("server error while parsingFS file: %w", err)
	}
	return Template{htmlTmpl: tmpl}, nil
}

func ParseTpl(filePath string) (Template, error) {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("server error while parsing file: %w", err)
	}
	return Template{htmlTmpl: tmpl}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Errorf("server error executing template: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}
