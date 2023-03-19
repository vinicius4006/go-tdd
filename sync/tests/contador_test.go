package tests

import (
	"sync"
	"sync-cp/entity"
	"testing"
)

func verificaContador(t *testing.T, resultado *entity.Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
	}
}

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := entity.NovoContador()
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()
		verificaContador(t, contador, 3)

	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := entity.NovoContador()

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)
		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)

		}
		wg.Wait()
		verificaContador(t, contador, contagemEsperada)

	})
}
