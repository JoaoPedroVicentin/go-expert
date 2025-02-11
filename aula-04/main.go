package main

import "fmt"

type ID int

var (
	e float64 = 1.12
	f ID      = 1
)

func main() {
	fmt.Printf("O valor de E é %T", f)
}
