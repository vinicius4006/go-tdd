package test

import (
	"context"
	"context-tdd/entity"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "olá, mundo"
	t.Run("avisa a store para cancelar o trabalho se a requisicao for cancelada", func(t *testing.T) {
		store := &entity.SpyStore{Response: data, T: t}
		svr :=
			entity.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.AssertWasCancelled()

	})

	t.Run("retorna dados da store", func(t *testing.T) {
		store := entity.SpyStore{Response: data, T: t}
		svr := entity.Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
		}

		store.AssertWasNotCancelled()
	})

}
