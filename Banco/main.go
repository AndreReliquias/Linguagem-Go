package main

import (
	c "Banco/clientes"
	Contas "Banco/contas"
	"fmt"
)

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	novoCliente := c.Titular{"João", "000.000.000-00", "Professor"}
	novaContaCorrente := Contas.ContaCorrente{Cliente: novoCliente, NumAgencia: 001, NumConta: 01}
	novaContaCorrente.Depositar(100)
	fmt.Println("Cliente 01: ", novaContaCorrente)

	novoClient := c.Titular{"Maria", "000.000.000-00", "Estudante"}
	novaContaPoupanca := Contas.ContaPoupanca{Cliente: novoClient, NumAgencia: 001, NumConta: 02, Operacao: 0}
	novaContaPoupanca.Depositar(500)
	fmt.Println("Cliente 02: ", novaContaPoupanca)

	PagarBoleto(&novaContaCorrente, 50)
	PagarBoleto(&novaContaPoupanca, 70)
	fmt.Println("Saldo final do cliente 1: ", novaContaCorrente.ObterSaldo())
	fmt.Println("Saldo final do cliente 2: ", novaContaPoupanca.ObterSaldo())

}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
