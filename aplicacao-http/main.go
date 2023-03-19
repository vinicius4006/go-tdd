package main

import (
	"appHttp/services"
	"log"
	"net/http"
)

func main() {

	tratador := http.HandlerFunc(services.ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}

}
