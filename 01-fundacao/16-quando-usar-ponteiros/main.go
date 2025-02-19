package main

func soma(a, b *int) int {
	*a = 50
	*b = 50

	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20

	println(var1)
	println(var2)

	result := soma(&var1, &var2)

	println(result)

	println(var1)
	println(var2)
}
