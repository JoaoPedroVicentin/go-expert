package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Olá Mundo!"
	}()

	msg := <-channel
	fmt.Println(msg)
}
