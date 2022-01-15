package main

import (
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("service.env")
	if err != nil {
		panic(err)
	}

	StartServer()
}
