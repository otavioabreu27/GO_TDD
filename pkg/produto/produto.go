package produto

import "math"

type Produto struct {
	Nome  string
	Valor float64
}

func NovoProduto(nome string, valor float64) Produto {
	return Produto{Nome: nome, Valor: arredondarDoisDigitosDecimais(valor)}
}

func (p *Produto) SetNome(nome string) {
	p.Nome = nome
}

func (p *Produto) SetValor(valor float64) {
	p.Valor = arredondarDoisDigitosDecimais(valor)
}

func arredondarDoisDigitosDecimais(valor float64) float64 {
	return math.Round(valor*100) / 100
}
