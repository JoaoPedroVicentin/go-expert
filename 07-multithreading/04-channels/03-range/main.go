package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	render(ch)
}

func render(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
