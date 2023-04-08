package test

import (
	"appHttp/poquer"
	"io/ioutil"
	"os"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	Pontuacoes        map[string]int
	RegistrosVitorias []string
	Liga              []poquer.Jogador
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.Pontuacoes[nome]
	return pontuacao
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() poquer.Liga {
	return e.Liga
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.RegistrosVitorias = append(e.RegistrosVitorias, nome)
}

func VerificaVitoriaJogador(t *testing.T, armazenamento *EsbocoArmazenamentoJogador, vencedor string) {
	t.Helper()

	if len(armazenamento.RegistrosVitorias) != 1 {
		t.Fatalf("recebi %d chamadas de GravarVitoria esperava %d", len(armazenamento.RegistrosVitorias), 1)
	}

	if armazenamento.RegistrosVitorias[0] != vencedor {
		t.Errorf("não armazenou o vencedor correto, recebi '%s' esperava '%s'", armazenamento.RegistrosVitorias[0], vencedor)
	}
}

func CriaArquivoTemporario(t *testing.T, dadoInicial string) (*os.File, func()) {
	t.Helper()

	arquivotmp, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("Não foi possível escrever o arquivo temporário %v", err)
	}

	arquivotmp.Write([]byte(dadoInicial))
	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}
