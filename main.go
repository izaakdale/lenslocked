package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/", homeHandler)
	mux.Get("/contact", contactHandler)
	mux.Get("/faq", faqHandler)
	mux.NotFound(notFoundHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("localhost:3000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcomen!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact</h1><p>To get in touch email me at <a href=\"mailto:izaakdale@live.com\">izaakdale@live.com</a></p>")
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Oops! Page not found", http.StatusNotFound)
}
