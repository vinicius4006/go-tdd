package poquer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	BancoDeDados *json.Encoder
	liga         Liga
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(arquivo *os.File) (*SistemaDeArquivoDeArmazenamentoDoJogador, error) {

	err := iniciaArquivoBDDJogador(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema inicializando arquivo do jogador, %v", err)
	}

	liga, err := NovaLiga(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema carregando o armazenamento do jogador %s, %v", arquivo.Name(), err)
	}

	return &SistemaDeArquivoDeArmazenamentoDoJogador{BancoDeDados: json.NewEncoder(&Fita{arquivo}), liga: liga}, nil
}

func iniciaArquivoBDDJogador(arquivo *os.File) error {
	arquivo.Seek(0, 0)

	info, err := arquivo.Stat()

	if err != nil {
		return fmt.Errorf("problema ao usar arquivo %s, %v", arquivo.Name(), err)
	}

	if info.Size() == 0 {
		arquivo.Write([]byte("[]"))
		arquivo.Seek(0, 0)
	}

	return nil
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	sort.Slice(s.liga, func(i, j int) bool {
		return s.liga[i].Vitorias > s.liga[j].Vitorias
	})
	return s.liga

}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) ObterPontuacaoJogador(nome string) int {

	jogador := s.ObterLiga().Find(nome)

	if jogador != nil {
		return jogador.Vitorias
	}

	return 0
}

func (s *SistemaDeArquivoDeArmazenamentoDoJogador) RegistrarVitoria(nome string) {

	jogador := s.liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		s.liga = append(s.liga, Jogador{Nome: nome, Vitorias: 1})
	}

	s.BancoDeDados.Encode(s.liga)

}

func NovaLiga(rdr io.Reader) ([]Jogador, error) {
	var liga []Jogador
	err := json.NewDecoder(rdr).Decode(&liga)
	if err != nil {
		err = fmt.Errorf("Problema parseando a liga, %v", err)
	}

	return liga, err
}
