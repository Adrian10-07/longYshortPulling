package controllers

import (
	"API-HEX-GO/src/pets/aplication"
	"API-HEX-GO/src/pets/infraestructure"
	"net/http"
	"strings"
)

func DeletePetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Nombre de la mascota requerido", http.StatusBadRequest)
		return
	}

	petName := pathParts[2]  

	var NombrePet string = petName

	repo := infraestructure.NewMySQLRepository()  
	useCase := aplication.NewDeletePet(repo)

	if err := useCase.Execute(NombrePet); err != nil {
		http.Error(w, "Error al eliminar la mascota", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mascota eliminada con éxito"))
}
