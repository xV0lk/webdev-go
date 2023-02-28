package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProcessTmpl(n string, w *http.ResponseWriter) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", n))
	if err != nil {
		return err
	}
	err = tmpl.Execute(*w, nil)
	if err != nil {
		return err
	}
	return nil
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home request")
	w.Header().Set("Content-Type", "text/html")
	err := ProcessTmpl("home", &w)
	if err != nil {
		http.Error(w, fmt.Errorf("server error while parsing file: %w", err).Error(), http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact request")
	fmt.Fprint(w, `<h1>Contact Page</h1
	><style>body {background-color: #3B3B3B;color: white}</style>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Faq Page")
	fmt.Fprint(w, `<h1>Faq</h1>
	<ul>
		<li>
			<p><strong>This is a question</strong></p>
			<p>This is the answer</p>
		</li>
		<li>
			<p><strong>This is another question: </strong></p>
			<p>This is another answer</p>
		</li>
	</ul>
	<style>body {background-color: #3B3B3B;color: white}</style>`)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Page")
	testId := chi.URLParam(r, "testId")
	fmt.Fprintf(w, `<h1>Url Param: %s</h1>
	<style>body {background-color: #3B3B3B;color: white}</style>`, testId)
}
