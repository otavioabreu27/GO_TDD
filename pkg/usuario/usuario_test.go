package usuario

import "testing"

func TestNovoUsuario(t *testing.T) {
	u := NovoUsuario("Nome Usuario", 20.50)

	if u.Nome != "Nome Usuario" {
		t.Errorf("Novo Usuario inicializado com nome errado")
	}

	if u.Saldo != 20.50 {
		t.Errorf("Novo Usuario inicializado com saldo errado")
	}
}

func TestSetNome(t *testing.T) {
	u := NovoUsuario("Nome Teste", 20.50)

	u.SetNome("Outro Nome")

	if u.Nome != "Outro Nome" {
		t.Errorf("Erro ao alterar nome do usuario. Esperado: 'Outro Nome', Obtido: %s", u.Nome)
	}
}

func TestAdicionaSaldo(t *testing.T) {
	u := NovoUsuario("Nome Teste", 20.50)

	u.AdicionaSaldo(10.50)

	if u.Saldo != 31.0 {
		t.Errorf("Erro ao adicionar saldo. Esperado: 31.0, Obtido: %f", u.Saldo)
	}
}

func TestSubtraiSaldo(t *testing.T) {
	u := NovoUsuario("Nome Teste", 20.50)

	err := u.SubtraiSaldo(10.50)

	if err != nil || u.Saldo != 10.0 {
		t.Errorf("Erro ao adicionar saldo. Esperado: 10.0, Obtido: %f", u.Saldo)
	}
}

func TestSubtraiSaldoInsuficiente(t *testing.T) {
	u := NovoUsuario("Nome Teste", 10.20)

	err := u.SubtraiSaldo(10.30)

	if err == nil {
		t.Errorf("Subtrair saldo insuficiente deveria mandar um erro")
	} else if err.Error() != "saldo insuficiente" {
		t.Errorf("Outro erro foi lancado deveria ser 'saldo insuficiente'")
	}
}
