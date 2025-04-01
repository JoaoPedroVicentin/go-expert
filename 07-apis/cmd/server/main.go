package main

import "go-expert/07-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
