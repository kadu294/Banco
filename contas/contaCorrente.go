package contas

import "github.com/kadu294/Banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	Saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque feito, seu valor agora é:", c.Saldo
	} else {
		return "Valor do Saque invalido, seu valor continua igual:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Valor Depositado, seu valor agora é:", c.Saldo
	} else {
		return "Deposite um valor valido, seu valor continua igual:", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorTranferencia float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valorTranferencia <= c.Saldo && valorTranferencia > 0
	if podeTransferir {
		c.Saldo -= valorTranferencia
		contaDestino.Depositar(valorTranferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() (string, float64) {
	return "Seu Saldo atual é:", c.Saldo
}
