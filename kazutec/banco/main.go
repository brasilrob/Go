package main

import (
	"fmt"

	"github.com/kazutec/banco/clientes"
	"github.com/kazutec/banco/contas"
)

type contaInterface interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta contaInterface, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {

	/* 	contaDoRoberto := ContaCorrente{titular: "Roberto", numeroAgencia: 589, numeroConta: 123456, saldo: 1000.05}
	   	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5} */

	//var contaDaCris *contas.ContaCorrente

	clienteCris := new(clientes.Titular)
	clienteCris.Nome = "Cris"
	clienteCris.CPF = "123.456.345-12"
	clienteCris.Profissao = "Advogada"

	contaDaCris := new(contas.ContaCorrente)
	contaDaCris.Titular = *clienteCris
	contaDaCris.Depositar(1250.00)

	//var contaDoGustavo *contas.ContaCorrente
	clienteGustavo := new(clientes.Titular)
	clienteGustavo.Nome = "Gustavo"
	clienteGustavo.CPF = "123.998.345-12"
	clienteGustavo.Profissao = "Desenvolvedor"

	contaDoGustavo := new(contas.ContaPoupanca)
	contaDoGustavo.Titular = *clienteGustavo
	contaDoGustavo.Operacao = 100
	contaDoGustavo.Depositar(305.00)

	status := contaDaCris.Sacar(110.00)
	fmt.Println(status)

	status, valor := contaDaCris.Depositar(1000)
	fmt.Println(status)
	fmt.Println(valor)

	PagarBoleto(contaDaCris, 500.)
	fmt.Println(contaDaCris.GetSaldo())

}
