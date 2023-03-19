package entity

import "sync"

type Contador struct {
	mu    sync.Mutex
	Count int
}

func (c *Contador) Incrementa() {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count++
}

func (c *Contador) Valor() int {
	return c.Count
}
func NovoContador() *Contador {

	return &Contador{}
}
