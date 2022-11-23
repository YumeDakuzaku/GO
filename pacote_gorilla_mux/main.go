//o mux (multiplexador de solicitação http) implementa um roteador e um despachante de solicitação
// instalar o gorilla mux --> go get -u github.com/gorilla/mux (se der erro ) no CMD digitar go mod init nomemodule (pq ainda não atualizaram a versão antiga)

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// request index page handle
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the index page!")
}

// request contact page handle
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the contact page!")
}

// request about page handle
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page!")
}

//go run main.go
// E acessar:localhost:8000/      localhost:8000/contact      localhost:8000/about
