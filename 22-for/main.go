package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}

	for i, numero := range numeros {
		println(i, numero)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Loop infinito")
	}
}
