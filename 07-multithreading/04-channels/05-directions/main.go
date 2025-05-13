package main

import "fmt"

func received(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go received("Hello", hello)
	read(hello)
}
