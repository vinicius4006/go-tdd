package test

import (
	"appHttp/poquer"
	"io"
	"strings"
	"testing"
)

type CLI struct {
	armazenamento poquer.ArmazenamentoJogador
	entrada       io.Reader
}

func (c *CLI) JogarPoquer() {
	c.armazenamento.RegistrarVitoria("Chris")
}

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris venceu\n")
	armazenamentoJogador := &EsbocoArmazenamentoJogador{}
	cli := &CLI{armazenamentoJogador, in}
	cli.JogarPoquer()

	if len(armazenamentoJogador.registrosVitorias) != 1 {
		t.Fatal("esperando uma chamada de vitoria mas nao recebi nenhuma")
	}

	obtido := armazenamentoJogador.registrosVitorias[0]
	esperado := "Chris"

	if obtido != esperado {
		t.Errorf("nao armazenou o vencedor correto, recebi '%s', esperava '%s'", obtido, esperado)
	}

}
