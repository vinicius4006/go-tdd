package main

import (
	"appHttp/poquer"
	"bufio"
	"fmt"
	"log"
	"os"
)

const nomeArquivoDB = "jogo.db.json"

func main() {

	armazenamento, close, err := poquer.ArmazenamentoSistemaDeArquivoJogadorAPartirDeArquivo(nomeArquivoDB)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	fmt.Println("Vamos jogar poquer!")
	fmt.Println("Digite '{Nome} venceu' para registrar uma vitória")

	poquer.NovoCLI(armazenamento, bufio.NewScanner(os.Stdin)).JogarPoquer()

}
