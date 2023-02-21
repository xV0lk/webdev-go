package main

import (
	"fmt"
	"net/http"
)

const DPORT = 3000

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Printf("Starting server on port: %v", DPORT)
	http.ListenAndServe(fmt.Sprintf(":%v", DPORT), nil)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request")
	fmt.Fprint(w, "<h1>Test <em>again</em></h1>")
}
