package main

import (
	"fmt"

	"github.com/kadu294/Banco/clientes"
	"github.com/kadu294/Banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) (string, float64) {
	conta.Sacar(valorBoleto)
	return "Um boleto foi debitado da sua conta no valor de:", valorBoleto
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	//contaTeste := contas.ContaCorrente{}
	//contaTeste.Titular = clientes.Titular{Nome: "NomdeTeste"}
	//contaTeste.NumeroAgencia = 589
	//contaTeste.NumeroConta = 123456
	//contaTeste.Saldo = 500
	//fmt.Println(contaTeste.Saldo)
	//status, valor := contaTeste.Depositar(100)
	//fmt.Println(status, valor)
	//fmt.Println(contaTeste.Sacar(200))
	contaTeste := contas.ContaCorrente{}
	contaTeste.Titular = clientes.Titular{Nome: "NomeTeste"}
	contaTeste.Saldo = 500
	contaTeste2 := contas.ContaCorrente{}
	contaTeste2.Titular = clientes.Titular{Nome: "NomeTeste2"}
	contaTeste2.Saldo = 300
	status := contaTeste.Transferir(200, &contaTeste2)
	fmt.Println(status)
	fmt.Println(contaTeste, "\n", contaTeste2)
	//clienteKadu := clientes.Titular{Nome: "Kadu", CPF: "111.222.333.44", Profissao: "Dev"}
	//contaKadu := contas.ContaCorrente{Titular: clienteKadu, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 400}
	//fmt.Println(contaKadu)
	//contaExemplo := contas.ContaCorrente{}
	//contaExemplo.Depositar(100)
	//fmt.Println(contaExemplo.ObterSaldo())
	//contaExemplo := contas.ContaPoupanca{}
	//fmt.Println(contaExemplo.Depositar(100))
	//fmt.Println(PagarBoleto(&contaExemplo, 100))
	//contaExemplo.Sacar(500)
	//fmt.Println(contaExemplo.ObterSaldo())
}
