package main

func main() {
	a := 1
	b := 2
	c := 3

	switch a {
	case 1:
		println("A é 1")
	case 2:
		println("A é 2")
	default:
		println("A é diferente de 1 e 2")
	}

	if a < b || a < c {
		println("A é menor que B e/ou C")
	}
}
