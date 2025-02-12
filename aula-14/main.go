package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func (e Empresa) Desativar() {

}

func Desativacao(p Pessoa) {
	p.Desativar()
}

func main() {

	joao := Cliente{
		Nome:  "João",
		Idade: 20,
		Ativo: true,
	}

	fmt.Printf("O nome do cliente é %s, sua idade é %d e ele está ativo? %t\n", joao.Nome, joao.Idade, joao.Ativo)

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)

}
