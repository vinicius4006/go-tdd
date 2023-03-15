package test

import (
	"reflect"
	"testing"
)

// Desafio Golang: escreva uma função percorre(x interface{}, fn func(string)) que recebe uma struct x e chama fn para todos os campos string encontrados dentro dela. nível de dificuldade: recursão.

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	// quantidadeDeValores := 0
	// var obtemCampo func(int) reflect.Value

	percorreValor := func(valor reflect.Value) {
		percorre(valor.Interface(), fn)
	}

	switch valor.Kind() {
	case reflect.Struct:
		for i := 0; i < valor.NumField(); i++ {
			percorreValor(valor.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < valor.Len(); i++ {
			percorreValor(valor.Index(i))
		}
	case reflect.Map:

		for _, chave := range valor.MapKeys() {
			percorreValor(valor.MapIndex(chave))
		}
	case reflect.String:
		fn(valor.String())
	}

}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)
	if valor.Kind() == reflect.Pointer {
		valor = valor.Elem()
	}

	return valor
}

func verificaSeContem(t *testing.T, palheiro []string, agulha string) {
	contem := false
	for _, x := range palheiro {
		if x == agulha {
			contem = true
		}
	}

	if !contem {
		t.Errorf("esperava-se que %+v contivesse '%s', mas não continha", palheiro, agulha)
	}
}

func TestPercorre(t *testing.T) {
	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{"Struct com um campo string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
		{
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Campos aninhados",
			Pessoa{"Chris", struct {
				Idade  int
				Cidade string
			}{33, "Londres"}}, []string{"Chris", "Londres"},
		},
		{
			"Ponteiros para coisas",
			&Pessoa{"Chris", Perfil{33, "Londres"}},
			[]string{"Chris", "Londres"},
		},
		{"Slices", []Perfil{{33, "Londres"}, {34, "Reykjavík"}}, []string{"Londres", "Reykjavík"}},
		{
			"Arrays",
			[2]Perfil{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
		{"Maps", map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}, []string{"Bar", "Boz"}},
	}
	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})
			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado %v, esperado %v", resultado, teste.ChamadasEsperadas)
			}
		})
	}

	t.Run("com maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var resultado []string
		percorre(mapA, func(entrada string) {
			resultado = append(resultado, entrada)
		})

		verificaSeContem(t, resultado, "Bar")
		verificaSeContem(t, resultado, "Boz")
	})

}
