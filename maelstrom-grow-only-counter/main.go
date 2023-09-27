package main

import (
	"log"
)

func main() {
	var h = &HandlerRepository{}
	h.RegisterHandlers()
	log.Printf("Starting ...")

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
