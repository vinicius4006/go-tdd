package main

type Liga []Jogador

func (l Liga) Find(nome string) *Jogador {
	for i, p := range l {
		if p.Nome == nome {
			return &l[i]
		}
	}

	return nil
}
