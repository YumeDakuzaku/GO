package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi everyone, i like %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//resumidamente ele pega o que vem depois da URL inicial e preenche na string.
// resultado --> "Hi everyone, i like bananas! da url" http://localhost:8000/banana
