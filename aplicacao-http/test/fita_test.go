package test

import (
	"appHttp/poquer"
	"io/ioutil"
	"testing"
)

func TestFita_Escrita(t *testing.T) {
	arquivo, limpa := criaArquivoTemporario(t, "12345")

	defer limpa()

	fita := &poquer.Fita{File: arquivo}

	fita.Write([]byte("abc"))

	arquivo.Seek(0, 0)

	novoConteudoDoArquivo, _ := ioutil.ReadAll(arquivo)

	recebido := string(novoConteudoDoArquivo)

	esperado := "abc"

	if recebido != esperado {
		t.Errorf("recebido '%s' esperado '%s'", recebido, esperado)
	}
}
