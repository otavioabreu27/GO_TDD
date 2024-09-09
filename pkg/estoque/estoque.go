package estoque

import (
	"errors"
	"fmt"

	"github.com/otavioabreu27/GO_TDD/pkg/produto"
)

type Estoque struct {
	Produtos    []*produto.Produto
	QtdProdutos []int64
}

func NovoEstoque() Estoque {
	return Estoque{}
}

func (e *Estoque) AdicionarProduto(p *produto.Produto, quantidade int64) {
	// Se o produto existe apenas adiciona a quantidade
	for i, prod := range e.Produtos {
		if prod.Nome == p.Nome {
			e.QtdProdutos[i] += quantidade
			return
		}
	}

	e.Produtos = append(e.Produtos, p)
	e.QtdProdutos = append(e.QtdProdutos, quantidade)
}

func (e *Estoque) VenderProduto(nomeProduto string, quantidade int64) error {
	for i, prod := range e.Produtos {
		if prod.Nome == nomeProduto {
			if e.QtdProdutos[i] >= quantidade {
				e.QtdProdutos[i] -= quantidade
				return nil
			} else {
				return errors.New("quantidade insuficiente em estoque")
			}
		}
	}

	return errors.New("produto não encontrado")
}

func (e *Estoque) ListaProdutos() {
	fmt.Println("Lista de Produtos")
	for i, prod := range e.Produtos {
		fmt.Printf("%d. %s: R$%.2f [%d]\n", i, prod.Nome, prod.Valor, e.QtdProdutos[i])
	}
}
