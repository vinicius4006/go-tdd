package main

import (
	"log"
	"net/http"
)

func main() {

	server := &ServidorJogador{Armazenamento: NovoArmazenamentoJogadorEmMemoria()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}

}
