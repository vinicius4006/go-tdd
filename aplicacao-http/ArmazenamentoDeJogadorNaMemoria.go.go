package main

type ArmazenamentoJogadorEmMemoria struct {
	Armazenamento map[string]int
}

// ObterLiga implements ArmazenamentoJogador
func (a *ArmazenamentoJogadorEmMemoria) ObterLiga() []Jogador {
	var liga []Jogador

	for nome, vitorias := range a.Armazenamento {
		liga = append(liga, Jogador{nome, vitorias})
	}

	return liga
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.Armazenamento[nome]
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.Armazenamento[nome]++
}

func NovoArmazenamentoJogadorEmMemoria() ArmazenamentoJogador {

	return &ArmazenamentoJogadorEmMemoria{
		Armazenamento: map[string]int{},
	}
}
