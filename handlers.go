package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func CreateTmplPath(fileName string) string {
	return filepath.Join("templates", fmt.Sprintf("%s.tmpl", fileName))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Page")
	testId := chi.URLParam(r, "testId")
	fmt.Fprintf(w, `<h1>Url Param: %s</h1>
	<style>body {background-color: #3B3B3B;color: white}</style>`, testId)
}
