package main

import (
	"log"

	"github.com/oat431/api-tutorial-go/routes"
)

func main() {
	r := routes.SetupRouter()
	log.Println("Router created")

	r.Run(":5000")
}
