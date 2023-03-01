package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xV0lk/webdev-go/controllers"
	"github.com/xV0lk/webdev-go/views"
)

const DPORT = 3000

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl, err := views.ParseTpl(CreateTmplPath("home"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl, err = views.ParseTpl(CreateTmplPath("contact"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl, err = views.ParseTpl(CreateTmplPath("faq"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tmpl))

	r.Get("/test/{testId}", testHandler)

	fmt.Printf("Started server on port: %v\n", DPORT)
	http.ListenAndServe(fmt.Sprintf(":%v", DPORT), r)
}
