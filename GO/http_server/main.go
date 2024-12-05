package main

import (
	"fmt"
	"html"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "home, %q", html.EscapeString(req.URL.Path))
}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
