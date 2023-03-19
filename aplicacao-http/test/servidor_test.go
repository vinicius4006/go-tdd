package test

import (
	"appHttp/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObterJogadores(t *testing.T) {

	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := httptest.NewRequest(http.MethodGet, "/jogadores/Maria", nil)
		resposta := httptest.NewRecorder()

		services.ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}

	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Pedro", nil)
		resposta := httptest.NewRecorder()

		services.ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}
	})
}
