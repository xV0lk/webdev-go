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

	tmpl := views.Must(views.ParseTpl(CreateTmplPath("home")))
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseTpl(CreateTmplPath("contact")))
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl = views.Must(views.ParseTpl(CreateTmplPath("faq")))
	r.Get("/faq", controllers.StaticHandler(tmpl))

	r.Get("/test/{testId}", testHandler)

	fmt.Printf("Started server on port: %v\n", DPORT)
	http.ListenAndServe(fmt.Sprintf(":%v", DPORT), r)
}
