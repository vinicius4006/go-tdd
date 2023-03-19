package services

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch jogador {
	case "Maria":
		fmt.Fprint(w, "20")
	case "Pedro":
		fmt.Fprint(w, "10")
	}
}
