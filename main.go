package main

import (
	"log"

	"github.com/oat431/api-tutorial-go/configs"
	"github.com/oat431/api-tutorial-go/routes"
)

func main() {
	configs.ConnectDB()
	log.Println("Database connected")

	r := routes.SetupRouter()
	log.Println("Router created")

	r.Run(":5000")
}
