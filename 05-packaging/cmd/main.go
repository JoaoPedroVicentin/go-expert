package main

import (
	"fmt"
	"go-expert/05-packaging/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
