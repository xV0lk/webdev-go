package controllers

import (
	"net/http"

	"github.com/xV0lk/webdev-go/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What is this?",
			Answer:   "This is a website for learning web development with Go.",
		},
		{
			Question: "Why?",
			Answer:   "Because I can.",
		},
		{
			Question: "How?",
			Answer:   "By studying very hard and writing a lot of code.",
		},
		{
			Question: "When?",
			Answer:   "Every day 4 am",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
