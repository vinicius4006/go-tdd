package poquer

import (
	"bufio"
	"strings"
)

type cli struct {
	armazenamento ArmazenamentoJogador
	entrada       *bufio.Scanner
}

func NovoCLI(arm ArmazenamentoJogador, en *bufio.Scanner) *cli {
	return &cli{arm, en}
}

func (c *cli) JogarPoquer() {

	userInput := c.readLine()
	c.armazenamento.RegistrarVitoria(extrairVencedor(userInput))
}

func extrairVencedor(userInput string) string {

	return strings.Replace(userInput, " venceu", "", 1)
}

func (c *cli) readLine() string {
	c.entrada.Scan()
	return c.entrada.Text()
}
