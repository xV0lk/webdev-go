package main

import (
	"fmt"
	"net/http"
)

const DPORT = 3000

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Test <em>again 2</em></h1><style>body {background-color: #3B3B3B;color: white}</style>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact request")
	fmt.Fprint(w, "<h1>Contact Page</h1><style>body {background-color: #3B3B3B;color: white}</style>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleHome(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// the following line wont work because cant display as html, only plain text
		// http.Error(w, fmt.Sprintf("<h1>Not Found</h1><p>You requested %s</p><style>body {background-color: #3B3B3B;color: white}</style>", r.URL.Path), http.StatusNotFound)

		// i will try writing to the header the status code
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, "<h1>Path Page</h1><p>You requested %s</p><style>body {background-color: #3B3B3B;color: white}</style>", r.URL.Path)
	}
}

func main() {
	var webRouter http.HandlerFunc = pathHandler
	fmt.Printf("Starting server on port: %v\n", DPORT)
	http.ListenAndServe(fmt.Sprintf(":%v", DPORT), webRouter)
}
