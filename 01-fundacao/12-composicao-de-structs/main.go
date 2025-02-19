package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {

	Joao := Cliente{
		Nome:  "João",
		Idade: 20,
		Ativo: true,
	}

	fmt.Printf("O nome do cliente é %s, sua idade é %d e ele está ativo? %t\n", Joao.Nome, Joao.Idade, Joao.Ativo)

	Joao.Ativo = false
	Joao.Address.Cidade = "São Paulo"

	fmt.Println(Joao.Ativo)

}
