package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const DPORT = 3000

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleHome)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	fmt.Printf("Started server on port: %v\n", DPORT)
	http.ListenAndServe(fmt.Sprintf(":%v", DPORT), r)
}
