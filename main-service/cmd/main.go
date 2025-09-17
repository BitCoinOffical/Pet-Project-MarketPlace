package main

import (
	"log"

	"myapp/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
