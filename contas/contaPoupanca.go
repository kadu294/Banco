package contas

import "github.com/kadu294/Banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque feito, seu valor agora é:", c.saldo
	} else {
		return "Valor do Saque invalido, seu valor continua igual:", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Valor Depositado, seu valor agora é:", c.saldo
	} else {
		return "Deposite um valor valido, seu valor continua igual:", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() (string, float64) {
	return "Seu saldo atual é:", c.saldo
}
