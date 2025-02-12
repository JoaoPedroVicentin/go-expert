package main

import "errors"

func main() {

	valor, err := sum(40, 20)

	if err != nil {
		println(err.Error())
	}
	println(valor)

}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("The sum is greater than 50")
	} else {
		return a + b, nil
	}
}
