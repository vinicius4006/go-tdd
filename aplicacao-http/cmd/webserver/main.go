package main

import (
	"appHttp/poquer"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problema abrindo %s %v", dbFileName, err)
	}
	armazenamento, errSistema := poquer.NovoSistemaDeArquivoDeArmazenamentoDoJogador(db)

	if errSistema != nil {
		log.Fatalf("problema criando o sistema de arquivo do armazenamento do jogador, %v ", err)
	}

	server := poquer.NovoServidorJogador(armazenamento)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}

}
