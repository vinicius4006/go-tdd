package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

const tipoDoConteudodoJSON = "application/json"

type EsbocoArmazenamentoJogador struct {
	pontuacoes        map[string]int
	registrosVitorias []string
	liga              []Jogador
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() Liga {
	return e.liga
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.registrosVitorias = append(e.registrosVitorias, nome)
}

// funções dentro do teste
func verificarCorpoResposta(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisicao é inválido obtive '%s' esperava '%s'", recebido, esperado)
	}

}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("não recebeu código de status HTTP esperado %d, recebido %d", esperado, recebido)
	}
}

func novaRequisicaoRegistrarVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func obterLigaDaResposta(t *testing.T, body io.Reader) (liga []Jogador) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&liga)

	if err != nil {
		t.Fatalf("Não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", body, err)
	}

	return

}

func verificaLiga(t *testing.T, obtido, esperado []Jogador) {
	t.Helper()
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v esperado %v", obtido, esperado)
	}
}

func novaRequisicaoDeLiga() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/liga", nil)

	return req
}

func verificaTipoDoConteudo(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	if resposta.Result().Header.Get("content-type") != esperado {
		t.Errorf("resposta não obteve content-type de %s, obtido %v", esperado, resposta.Result().Header)
	}
}

// Testes começam aqui
func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{map[string]int{
		"Maria": 20,
		"Pedro": 10,
	}, nil, nil}

	servidor := NovoServidorJogador(&armazenamento)
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoResposta(t, recebido, esperado)

	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoResposta(t, recebido, esperado)
	})

	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperado := http.StatusNotFound

		if recebido != esperado {
			t.Errorf("recebido status %d esperado %d", recebido, esperado)
		}
	})
}

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		[]string{},
		nil,
	}

	servidor := NovoServidorJogador(&armazenamento)

	t.Run("registra vitorias na chamada ao método HTTP POST", func(t *testing.T) {
		jogador := "Maria"
		requisicao := novaRequisicaoRegistrarVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusAccepted)

		if len(armazenamento.registrosVitorias) != 1 {
			t.Errorf("verifiquei %d chamadas a RegistrarVitoria, esperava %d", len(armazenamento.registrosVitorias), 1)
		}

		if armazenamento.registrosVitorias[0] != jogador {
			t.Errorf("não registrou o vencedor corretamente, recebi '%s', esperava '%s'", armazenamento.registrosVitorias[0], jogador)
		}
	})
}

func TestRegistrarVitoriasEBuscarEstasVitorias(t *testing.T) {
	bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, "")
	defer limpaBancoDeDados()
	armazenamento := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)
	servidor := NovoServidorJogador(armazenamento)
	jogador := "Maria"

	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))

	resposta := httptest.NewRecorder()
	servidor.ServeHTTP(resposta, novaRequisicaoObterPontuacao(jogador))
	verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

	verificarCorpoResposta(t, resposta.Body.String(), "3")

}

func TestLiga(t *testing.T) {

	t.Run("retorna 200 em /liga", func(t *testing.T) {
		armazenamento := EsbocoArmazenamentoJogador{}
		servidor := NovoServidorJogador(&armazenamento)
		requisicao := novaRequisicaoDeLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		var obtido []Jogador
		err := json.NewDecoder(resposta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("Não foi possível fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", resposta.Body, err)
		}

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

	})

	t.Run("retorna a tabela da Liga como JSON", func(t *testing.T) {
		ligaEsperada := []Jogador{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		armazenamento := EsbocoArmazenamentoJogador{nil, nil, ligaEsperada}

		servidor := NovoServidorJogador(&armazenamento)

		requisicao := novaRequisicaoDeLiga()

		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		obtido := obterLigaDaResposta(t, resposta.Body)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

		verificaLiga(t, obtido, ligaEsperada)

		verificaTipoDoConteudo(t, resposta, tipoDoConteudodoJSON)

	})
}

func TestGravaVitoriasEAsRetorna(t *testing.T) {
	bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, "")
	armazenamento := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)
	defer limpaBancoDeDados()
	servidor := NovoServidorJogador(armazenamento)
	jogador := "Pepper"

	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistrarVitoriaPost(jogador))

	t.Run("obter pontuacao", func(t *testing.T) {
		resposta := httptest.NewRecorder()
		servidor.ServeHTTP(resposta, novaRequisicaoObterPontuacao(jogador))

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

		verificarCorpoResposta(t, resposta.Body.String(), "3")

	})

	t.Run("obter liga", func(t *testing.T) {
		resposta := httptest.NewRecorder()
		servidor.ServeHTTP(resposta, novaRequisicaoDeLiga())
		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

		obtido := obterLigaDaResposta(t, resposta.Body)

		esperado := []Jogador{
			{"Pepper", 3},
		}

		verificaLiga(t, obtido, esperado)
	})

}

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	if recebido != esperado {
		t.Errorf("recebido %d esperado %d", recebido, esperado)
	}
}

func criaArquivoTemporario(t *testing.T, dadoInicial string) (io.ReadWriteSeeker, func()) {
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

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {
	bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
		{"Nome": "Cleo", "Vitorias": 10},
		{"Nome": "Chris", "Vitorias": 33}]`)
	armazenamento := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)
	defer limpaBancoDeDados()
	t.Run("/liga de um leitor", func(t *testing.T) {

		recebido := armazenamento.ObterLiga()

		esperado := []Jogador{
			{"Cleo", 10},
			{"Chris", 33},
		}

		verificaLiga(t, recebido, esperado)

		recebido = armazenamento.ObterLiga()
		verificaLiga(t, recebido, esperado)

	})

	t.Run("pegar pontuação do jogador", func(t *testing.T) {
		recebido := armazenamento.ObterPontuacaoJogador("Chris")

		esperado := 33

		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitórias de um jogador existente", func(t *testing.T) {
		armazenamento.RegistrarVitoria("Chris")

		recebido := armazenamento.ObterPontuacaoJogador("Chris")

		esperado := 34

		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitorias de novos jogadores", func(t *testing.T) {
		armazenamento.RegistrarVitoria("Pepper")

		recebido := armazenamento.ObterPontuacaoJogador("Pepper")
		esperado := 1

		definePontuacaoIgual(t, recebido, esperado)
	})
}
