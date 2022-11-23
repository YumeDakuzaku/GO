package main //diz qual pacote estamos trabalhando, nesse caso - pacote principal

import ( //imports dos pacotes que serão utilizados
	"fmt"
	"log"
	"net/http"
)

// manipulador de request handle, essa função será usada para toda solicitação feita ao nosso servidor
// handler tem como entrada dois objetos writer de resposta e o objeto de solicitação request que não retorna nada
// dentro da handler nós só vamos ter a impressão através do fmt para nosso writer de resposta
// resposewriter monta a resposta ao servidor http, escrevendo nele os dados para o cliente http, ou seja o browser
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

// a função principal main informa ao pacote http para manipular todas as solicitações para a raiz da web("/") com o manipulador
// listenandserver - escutar na porta 8000 em qualquer interface :8000
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil)) //liga servidor e caso de algum erro gera o log
}

//depois de executar o main (go run main.go) acessar no browser: localhost:8000
