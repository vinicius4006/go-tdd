package main

import (
	"mocks/entity"
	"os"
	"time"
)

type SleeperPadrao struct {
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &SleeperPadrao{}
	entity.Contagem(os.Stdout, sleeper)
}
