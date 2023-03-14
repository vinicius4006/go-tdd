package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Corredor(a, b string) (vencedor string) {

	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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

// func medirTempoDeResposta(URL string) time.Duration {
// 	inicio := time.Now()
// 	http.Get(URL)
// 	return time.Since(inicio)
// }

func TestCorredor(t *testing.T) {

	servidorLento := criarServidorComAtraso(20 * time.Millisecond)

	servidorRapido := criarServidorComAtraso(19 * time.Millisecond)

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida

	resultado := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

	servidorLento.Close()
	servidorRapido.Close()

}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}
