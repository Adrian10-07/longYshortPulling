package main

import (
	"log"
	"net/http"
	"API-HEX-GO/src/Products/infraestructure/routes"
	petsroutes "API-HEX-GO/src/pets/infraestructure/routespets"
)
func main() {


	petsroutes.SetupRoutes()

	routes.SetupRoutes()

	log.Print("server listen in puert 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
