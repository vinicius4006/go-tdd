package test

import (
	"bytes"
	"mocks/entity"
	"reflect"
	"testing"
	"time"
)

const escrita = "escrita"
const pausa = "pausa"

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

func TestContagem(t *testing.T) {
	t.Run("imprime 3 até vai", func(t *testing.T) {
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

	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		entity.Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			"pausa",
			"escrita",
			"pausa",
			"escrita",
			"pausa",
			"escrita",
			"pausa",
			"escrita",
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})

}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	TempoSpy := &TempoSpy{}
	sleeper := entity.SleeperConfiguravel{Duracao: tempoPausa, Pausa: TempoSpy.Pausa}
	sleeper.Sleep()

	if TempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, TempoSpy.duracaoPausa)
	}
}
