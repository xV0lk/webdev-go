package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/xV0lk/webdev-go/views"
)

func ProcessTmpl(w http.ResponseWriter, filePath string) {
	t, err := views.ParseTpl(filePath)
	if err != nil {
		http.Error(w, "There was a server error while loading your page", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func CreateTmplPath(fileName string) string {
	return filepath.Join("templates", fmt.Sprintf("%s.tmpl", fileName))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tplPath := CreateTmplPath("home")
	fmt.Println("Home request")
	ProcessTmpl(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := CreateTmplPath("contact")
	fmt.Println("Contact request")
	ProcessTmpl(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := CreateTmplPath("faq")
	fmt.Println("Faq Page")
	ProcessTmpl(w, tplPath)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Page")
	testId := chi.URLParam(r, "testId")
	fmt.Fprintf(w, `<h1>Url Param: %s</h1>
	<style>body {background-color: #3B3B3B;color: white}</style>`, testId)
}
