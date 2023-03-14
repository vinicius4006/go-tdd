package entity

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperConfiguravel struct {
	Duracao time.Duration
	Pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.Pausa(s.Duracao)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}
