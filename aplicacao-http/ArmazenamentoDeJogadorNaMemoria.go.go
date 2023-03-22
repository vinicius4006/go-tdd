package main

type ArmazenamentoJogadorEmMemoria struct {
	Armazenamento map[string]int
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
