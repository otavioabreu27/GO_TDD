package main

import (
	"github.com/otavioabreu27/GO_TDD/pkg/estoque"
	"github.com/otavioabreu27/GO_TDD/pkg/produto"
)

func main() {
	e := estoque.NovoEstoque()

	p := produto.NovoProduto("Pasta de Dente", 20.24)

	e.AdicionarProduto(&p, 20)

	e.ListaProdutos()

	p.SetNome("Creme Dental")

	e.ListaProdutos()
}
