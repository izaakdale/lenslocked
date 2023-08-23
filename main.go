package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/izaakdale/lenslocked/controllers"
	"github.com/izaakdale/lenslocked/views"
)

func main() {

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	mux.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	mux.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	mux.Get("/faq", controllers.StaticHandler(tpl))

	mux.NotFound(notFoundHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", mux)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Oops! Page not found", http.StatusNotFound)
}
