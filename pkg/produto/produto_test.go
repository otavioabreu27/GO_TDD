package produto

import "testing"

func TestNovoProduto(t *testing.T) {
	p := NovoProduto("Lapis", 20.0)

	if p.Nome != "Lapis" {
		t.Errorf("Nome esperado que fosse 'Lapis', obtido %s", p.Nome)
	}

	if p.Valor != 20 {
		t.Errorf("Valor esperado que fosse '20.0', obtido %f", p.Valor)
	}
}

func TestSetNomeProduto(t *testing.T) {
	p := NovoProduto("Teste", 20.0)

	p.SetNome("Outro")

	if p.Nome != "Outro" {
		t.Errorf("Valor esperado 'Outro', obtido %s", p.Nome)
	}
}

func TestSetValorProduto(t *testing.T) {
	p := NovoProduto("Produto Teste", 30.5)

	p.SetValor(15.45)

	if p.Valor != 15.45 {
		t.Errorf("Houve um erro tentando mudar o valor do produto. Esperado: '15.45' Obtido: %f", p.Valor)
	}
}

func TestValorDuasCasasPrecisaoDecimal(t *testing.T) {
	p := NovoProduto("Produto Teste", 12.15134)

	if p.Valor != 12.15 {
		t.Errorf("Precisao nao equivalente. Esperado: '12.15', Obtido: %f", p.Valor)
	}
}

func TestSetValorDuasCasasPrecisaoDecimal(t *testing.T) {
	p := NovoProduto("Produto teste", 0.0)

	p.SetValor(12.45623)

	if p.Valor != 12.46 {
		t.Errorf("Precisao nao equivalente. Esperado: '12.46', Obtido: %f", p.Valor)
	}
}
