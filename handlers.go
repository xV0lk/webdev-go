package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProcessTmpl(w http.ResponseWriter, n string) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", n))
	if err != nil {
		http.Error(w, fmt.Errorf("server error while parsing file: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Errorf("server error while parsing file: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home request")
	ProcessTmpl(w, "home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact request")
	ProcessTmpl(w, "contact")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Faq Page")
	ProcessTmpl(w, "faq")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Page")
	testId := chi.URLParam(r, "testId")
	fmt.Fprintf(w, `<h1>Url Param: %s</h1>
	<style>body {background-color: #3B3B3B;color: white}</style>`, testId)
}
