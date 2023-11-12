package main

import (
	"log"

	_ "github.com/oat431/api-tutorial-go/docs"
	"github.com/oat431/api-tutorial-go/routes"
)

// @title Tutorial API
// @version 1.0
// @description.markdown

// @contact.name Sahachan
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := routes.SetupRouter()
	log.Println("CORS setup")
	log.Println("Router setup")
	r.Run(":8080")
}
