package main

import (
	"log"

	"github.com/oat431/api-tutorial-go/routes"
)

func main() {
	r := routes.SetupRouter()
	log.Println("CORS setup")
	log.Println("Router setup")
	r.Run(":8080")
}
