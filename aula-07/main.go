package main

import "fmt"

func main() {
	salarios := map[string]int{"Joao": 1200, "Maria": 2000, "Jose": 1500}

	fmt.Println(salarios["Joao"])

	delete(salarios, "Joao")

	salarios["Pedro"] = 1800

	fmt.Println(salarios)

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
