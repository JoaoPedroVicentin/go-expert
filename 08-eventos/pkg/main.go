package main

import "fmt"

func main() {
	events := []string{"teste", "teste2", "teste3", "teste4"}
	events = append(events[:0], events[1:]...)
	fmt.Println(events)
}
