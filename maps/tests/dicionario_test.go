package main

import (
	"maps/entity"
	"testing"
)

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func comparaDefinicao(t *testing.T, dicionario entity.Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("deveria ter encontrado a palavra adicionada", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
}

func TestBusca(t *testing.T) {
	dicionario := entity.Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")
		comparaErro(t, resultado, entity.ErrNaoEncontrado)
	})

}

func TestAdiciona(t *testing.T) {
	t.Run("adicionar palavra nova", func(t *testing.T) {
		dicionario := entity.Dicionario{}
		dicionario.Adiciona("teste", "isso é apenas um teste")

		comparaDefinicao(t, dicionario, "teste", "isso é apenas um teste")
	})

	t.Run("adicionar palavra existente", func(t *testing.T) {
		dicionario := entity.Dicionario{"teste": "isso é apenas um teste"}
		err := dicionario.Adiciona("teste", "isso é apenas um teste2")

		comparaErro(t, err, entity.ErrPalavraExistente)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		novaDefinicao := "nova definição"
		dicionario := entity.Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)

	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := entity.Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErro(t, err, entity.ErrPalavraInexistente)
	})

}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := entity.Dicionario{palavra: "definicao de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != entity.ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)

	}
}
