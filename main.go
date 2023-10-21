package main

import (
	"github.com/oat431/api-tutorial-go/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
