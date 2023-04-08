package poquer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const jsonContentType = "application/json"

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
	RegistrarVitoria(nome string)
	ObterLiga() Liga
}

type ServidorJogador struct {
	Armazenamento ArmazenamentoJogador
	http.Handler
}

func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	s := new(ServidorJogador)

	s.Armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.manipulaJogadores))

	s.Handler = roteador

	return s

}

func (s *ServidorJogador) manipulaLiga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.Armazenamento.ObterLiga())
	w.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) manipulaJogadores(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]
	switch r.Method {
	case http.MethodPost:
		s.registrarVitoria(w, r, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, r, jogador)
	}
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, r *http.Request, jogador string) {

	pontuacao := s.Armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter, r *http.Request, jogador string) {

	s.Armazenamento.RegistrarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}

func ArmazenamentoSistemaDeArquivoJogadorAPartirDeArquivo(path string) (*SistemaDeArquivoDeArmazenamentoDoJogador, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("falha ao abrir %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(db)

	if err != nil {
		return nil, nil, fmt.Errorf("falha ao criar sistema de arquivos para armazenar jogadores, %v ", err)
	}

	return armazenamento, closeFunc, nil

}
