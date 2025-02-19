package soma

type Carro struct {
	Marca string
}

var A int = 10

func Soma[T int | float64](a, b T) T {
	return a + b
}

func (c Carro) Andar() string {
	return c.Marca + " andou"
}
