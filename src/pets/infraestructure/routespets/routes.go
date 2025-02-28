package routes

import (
	"net/http"
	"API-HEX-GO/src/pets/infraestructure/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/pets", controllers.CreatePetHandler)
	http.HandleFunc("/view-pets", controllers.GetPetHandler)
	http.HandleFunc("/delete-pets/", controllers.DeletePetHandler)
	http.HandleFunc("/update-pets/", controllers.UpdatePetHandler)    
}
