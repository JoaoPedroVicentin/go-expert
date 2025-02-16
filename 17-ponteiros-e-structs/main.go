package main

import "fmt"

type Cliente struct {
	nome  string
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Joao Pedro"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func (c *Cliente) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func NewCliente() *Cliente {
	return &Cliente{
		nome:  "Bruno",
		saldo: 0,
	}
}

func main() {

	joao := Cliente{
		nome:  "Joao",
		saldo: 1000,
	}

	joao.andou()

	joao.simular(2000)

	println(joao.saldo)

	fmt.Printf("O valor da struct com nome é %v\n", joao.nome)
}
