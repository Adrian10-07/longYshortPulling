package controllers

import (
	"API-HEX-GO/src/pets/aplication"
	"API-HEX-GO/src/pets/domain"
	"API-HEX-GO/src/pets/infraestructure"
	"encoding/json"
	"net/http"
	"strings"
)

func UpdatePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Nombre de la mascota requerido en la URL", http.StatusBadRequest)
		return
	}



	idPet := pathParts[2]


	var updatedPet domain.Pet
	err := json.NewDecoder(r.Body).Decode(&updatedPet)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMySQLRepository()  
	useCase := aplication.NewEditPet(repo)

	err = useCase.Execute(idPet, &updatedPet)
	if err != nil {
		http.Error(w, "Error al actualizar la mascota", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mascota actualizada correctamente"))
}
