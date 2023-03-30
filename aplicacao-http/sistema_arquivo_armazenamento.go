package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	BancoDeDados io.ReadWriteSeeker
	liga         Liga
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados io.ReadWriteSeeker) *SistemaDeArquivoDeArmazenamentoDoJogador {
	bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(bancoDeDados)
	return &SistemaDeArquivoDeArmazenamentoDoJogador{BancoDeDados: bancoDeDados, liga: liga}
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	s.BancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(s.BancoDeDados)
	return liga

}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterPontuacaoJogador(nome string) int {

	jogador := s.ObterLiga().Find(nome)

	if jogador != nil {
		return jogador.Vitorias
	}

	return 0
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) RegistrarVitoria(nome string) {
	liga := s.ObterLiga()
	jogador := liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		liga = append(liga, Jogador{nome, 1})
	}

	s.BancoDeDados.Seek(0, 0)
	json.NewEncoder(s.BancoDeDados).Encode(liga)
}

func NovaLiga(rdr io.Reader) ([]Jogador, error) {
	var liga []Jogador
	err := json.NewDecoder(rdr).Decode(&liga)
	if err != nil {
		err = fmt.Errorf("Problema parseando a liga, %v", err)
	}

	return liga, err
}
