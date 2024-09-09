package estoque

import (
	"testing"

	"github.com/otavioabreu27/GO_TDD/pkg/produto"
)

func TestEstoqueComecaComProdutosVazio(t *testing.T) {
	e := NovoEstoque()

	if e.Produtos != nil {
		t.Errorf("Produtos inicializados antes de adicionar")
	}
}

func TestAdicionarProduto(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Produto Teste", 20)

	e.AdicionarProduto(&p, 10)

	if e.Produtos[0] != &p {
		t.Errorf("O produto novo nao foi adicionado")
	}
}

func TestAdicionarProdutoAdicionaQuantidade(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Produto Teste", 5)

	e.AdicionarProduto(&p, 5)

	e.AdicionarProduto(&p, 5)

	if e.QtdProdutos[0] != 10 {
		t.Errorf("Adicionar o mesmo produto nao esta adicionando a quantidade")
	}
}

func TestVenderProdutoDiminuiQuantidade(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Pasta de Dente", 20.1254)

	e.AdicionarProduto(&p, 10)

	e.VenderProduto("Pasta de Dente", 5)

	if e.QtdProdutos[0] != 5 {
		t.Errorf("Vender produto nao esta subtraindo o valor certo")
	}
}

func TestVenderProdutoSemEstoqueChamaErro(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Pasta de Dente", 20.1254)

	e.AdicionarProduto(&p, 2)

	err := e.VenderProduto("Pasta de Dente", 3)

	if err == nil {
		t.Errorf("Vender produto fora do estoque deveria retornar um erro.")
	}

	expectedErrorMessage := "quantidade insuficiente em estoque"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Erro errado encontrado. Esperado: '%s', Obtido: %s", expectedErrorMessage, err.Error())
	}
}

func TestVenderProdutoSemRegistroChamaErro(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Pasta de Dente", 20.1254)

	e.AdicionarProduto(&p, 10)

	err := e.VenderProduto("Escova de Dente", 3)

	if err == nil {
		t.Errorf("Vender produto sem registro deveria retornar um erro.")
	}

	expectedErrorMessage := "produto não encontrado"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Erro errado encontrado. Esperado: '%s', Obtido: %s", expectedErrorMessage, err.Error())
	}
}

func TestEstoqueListaProdutos(t *testing.T) {
	e := NovoEstoque()

	p := produto.NovoProduto("Pasta de Dente", 20.1254)

	e.AdicionarProduto(&p, 10)

	e.ListaProdutos()
}
