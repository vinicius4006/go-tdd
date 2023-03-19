package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//revise o código
// Te pediram para fazer uma função chamada Corredor que recebe duas URLs que "competirão" entre si através de uma chamada HTTP GET onde a primeira URL a responder será retornada. Se nenhuma delas responder dentro de 10 segundos a função deve retornar um erro.

var limiteDeDezSegundos = 10 * time.Second

func Corredor(a, b string) (vencedor string, err error) {
	return Configuravel(a, b, limiteDeDezSegundos)
}
func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, err error) {
	select {
	//ele executa os casos, o que retornar primeiro ele manda
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}

// func medirTempoDeResposta(URL string) time.Duration {
// 	inicio := time.Now()
// 	http.Get(URL)
// 	return time.Since(inicio)
// }

func TestCorredor(t *testing.T) {

	t.Run("compara a velocidade de servidores, retornando o endereço mais rápido", func(t *testing.T) {
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)

		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		defer servidorLento.Close()
		defer servidorRapido.Close()

		esperado := URLRapida

		resultado, _ := Corredor(URLLenta, URLRapida)

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidorA := criarServidorComAtraso(11 * time.Second)
		servidorB := criarServidorComAtraso(12 * time.Second)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Configuravel(servidorA.URL, servidorB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}

	})

}
