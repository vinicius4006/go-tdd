package main

import (
	"mocks/entity"
	"os"
	"time"
)

func main() {
	sleeper := &entity.SleeperConfiguravel{Duracao: 1 * time.Second, Pausa: time.Sleep}

	entity.Contagem(os.Stdout, sleeper)
}
