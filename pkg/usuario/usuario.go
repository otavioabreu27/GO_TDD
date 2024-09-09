package usuario

import "errors"

type Usuario struct {
	Nome  string
	Saldo float64
}

func NovoUsuario(nome string, saldo float64) Usuario {
	return Usuario{Nome: nome, Saldo: saldo}
}

func (u *Usuario) SetNome(nome string) {
	u.Nome = nome
}

func (u *Usuario) AdicionaSaldo(saldo float64) {
	u.Saldo += saldo
}

func (u *Usuario) SubtraiSaldo(saldo float64) error {
	if u.Saldo-saldo < 0 {
		return errors.New("saldo insuficiente")
	}

	u.Saldo -= saldo
	return nil
}
