package test

import (
	"bytes"
	"mocks/entity"
	"testing"
)

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func TestContagem(t *testing.T) {
	//buffer porque implementa interface writer
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	entity.Contagem(buffer, sleeperSpy)

	resultado := buffer.String()
	esperado := "3\n2\n1\nVai!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

	if sleeperSpy.Chamadas != 4 {
		t.Errorf("n houve chamadas suficientes do sleeper, esperado 4, resultado %d", sleeperSpy.Chamadas)
	}
}
