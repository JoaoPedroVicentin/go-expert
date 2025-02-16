package main

import (
	"fmt"
	"go-expert/21-pacotes-e-modulos/soma"

	"github.com/google/uuid"
)

func main() {
	s := soma.Soma(10, 20)
	carro := soma.Carro{Marca: "Fiat"}

	fmt.Println("Resultado da soma:", s)
	fmt.Println(soma.A)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}
