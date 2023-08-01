package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", Router{})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcomen!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact</h1><p>To get in touch email me at <a href=\"mailto:izaakdale@live.com\">izaakdale@live.com</a></p>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}
