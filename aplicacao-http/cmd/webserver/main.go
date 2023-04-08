package main

import (
	"appHttp/poquer"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	armazenamento, close, err := poquer.ArmazenamentoSistemaDeArquivoJogadorAPartirDeArquivo(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poquer.NovoServidorJogador(armazenamento)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}

}
