package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"log"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func Contact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt. Fprint(w, "To get in touch, please send an email " +
	"to  <a href=\"mailto:support@eolas.ca\">" +
	"support@lenslocked.com</a>.")
}

func FAQ(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt. Fprint(w, "<h1>Frequently asked questions</h1>" +
		"<p>Find answers here</p>")
}

func NotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 Page not found!</h1>")
}

var nf http.Handler = http.HandlerFunc(notFound)

func main() {
	r := httprouter.New()
	r.GET("/", Home)
	r.GET("/contact", Contact)
	r.GET("/faq", FAQ)
	r.NotFoundHandler = nf

	log.Fatal(http.ListenAndServe(":3000", r))
}