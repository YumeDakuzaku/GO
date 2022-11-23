//main.go vai importar o pacote html/template em nosso projeto, os templates precisam estar acessíveis a partir das rotas
//para isso - criar um objeto chamado templates - ele será usado dentro de cada handler, para chamar o executetemplate do objeto templates e linkar com o html correto para renderizar

package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates_ *template.Template

func main() {
	templates_ = template.Must(template.ParseGlob("templates_/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// request index page handle
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates_.ExecuteTemplate(w, "index.html", nil)
	//fmt.Fprint(w, "This is the index page!")
}

// request contact page handle
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates_.ExecuteTemplate(w, "contact.html", nil)
	//fmt.Fprint(w, "This is the contact page!")
}

// request about page handle
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templates_.ExecuteTemplate(w, "about.html", nil)
	//fmt.Fprint(w, "This is the about page!")
}
