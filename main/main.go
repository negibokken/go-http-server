package main

import (
	"log"

	httpserver "github.com/negibokken/go-http-server"
)

func main() {
	_, err := httpserver.NewServer()
	if err != nil {
		log.Fatal("server creation error")

	}
}
