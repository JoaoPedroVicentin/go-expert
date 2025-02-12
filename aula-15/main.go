package main

import "fmt"

func main() {
	a := 10
	var ponteiro *int = &a

	*ponteiro = 20
	b := &a

	fmt.Printf("Valor de A: %d\n", a)

	fmt.Printf("Endereço de A pelo ponteiro: %p\n", ponteiro)
	fmt.Printf("Endereço de A pelo B: %p\n", b)

	fmt.Printf("Valor de A acessado pelo ponteiro: %d\n", *ponteiro)
	fmt.Printf("Valor de A acessado pelo B: %d\n", *b)

	fmt.Printf("Endereço do ponteiro: %p\n", &ponteiro)
	fmt.Printf("Endereço de B: %p\n", &b)
}
