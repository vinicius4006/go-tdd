package poquer_test

import (
	"appHttp/poquer"
	"appHttp/test"
	"bufio"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("recorda vencedor chris digitado pelo usuario", func(t *testing.T) {

		in := strings.NewReader("Chris venceu\n")
		armazenamentoJogador := &test.EsbocoArmazenamentoJogador{}
		cli := poquer.NovoCLI(armazenamentoJogador, bufio.NewScanner(in))
		cli.JogarPoquer()

		test.VerificaVitoriaJogador(t, armazenamentoJogador, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		in := strings.NewReader("Cleo venceu\n")
		armazenamentoJogador := &test.EsbocoArmazenamentoJogador{}

		cli := poquer.NovoCLI(armazenamentoJogador, bufio.NewScanner(in))
		cli.JogarPoquer()

		test.VerificaVitoriaJogador(t, armazenamentoJogador, "Cleo")

	})

}
