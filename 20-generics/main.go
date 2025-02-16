package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {

	m := map[string]int{"João": 10, "Maria": 20, "José": 30}
	m2 := map[string]float64{"João": 10.20, "Maria": 20.0, "José": 30.54}

	m3 := map[string]MyNumber{"João": 10, "Maria": 20, "José": 30}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
