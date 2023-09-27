package main

import (
	"log"
	"os"
)

func main() {

	RegisterHandlers()

	log.Printf("starting ...")

	if err := n.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
