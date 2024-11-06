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
	contaTeste := contas.ContaCorrente{}
	contaTeste.Titular = clientes.Titular{Nome: "NomdeTeste"}
	contaTeste.NumeroAgencia = 589
	contaTeste.NumeroConta = 123456
	contaTeste.Saldo = 500
	fmt.Println(contaTeste.Saldo)
	status, valor := contaTeste.Depositar(100)
	fmt.Println(status, valor)
	fmt.Println(contaTeste.Sacar(200))
	//------------------------------------------
	/*contaTeste := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Nometeste"}, Saldo: 50}
	contaTeste2 := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Nometeste2"}, Saldo: 300}
	status := contaTeste.Transferir(200, &contaTeste2)
	if status {
		fmt.Println("Transferência realizada com sucesso")
	} else {
		fmt.Println("Falha na transferência")
	}
	fmt.Printf("Saldo de contaTeste: %.2f\n", contaTeste.Saldo)
	fmt.Printf("Saldo de contaTeste2: %.2f\n", contaTeste2.Saldo)*/
	//------------------------------------------
	/* clienteKadu := clientes.Titular{Nome: "Kadu", CPF: "111.222.333.44", Profissao: "Dev"}
	contaKadu := contas.ContaCorrente{Titular: clienteKadu, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 400}
	fmt.Printf("Conta Corrente de %s:\n%+v\n", contaKadu.Titular.Nome, contaKadu) */
	/* contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	fmt.Println(contaExemplo.ObterSaldo()) */
	//------------------------------------------
	/* contaExemplo := contas.ContaPoupanca{}
	fmt.Println(contaExemplo.Depositar(1000))
	fmt.Println(contaExemplo.ObterSaldo())
	fmt.Println(PagarBoleto(&contaExemplo, 200))
	contaExemplo.Sacar(1000)
	fmt.Println(contaExemplo.ObterSaldo()) */
}
