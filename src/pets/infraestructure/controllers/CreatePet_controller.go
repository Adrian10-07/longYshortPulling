package controllers

import (
	"encoding/json"
	"net/http"
	"API-HEX-GO/src/pets/aplication"
	"API-HEX-GO/src/pets/domain"
	"API-HEX-GO/src/pets/infraestructure"
)

func CreatePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var pet domain.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}
	
	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewCreatePet(repo)

	if err := useCase.Execute(pet); err != nil {
		http.Error(w, "Error al guardar la mascota", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Mascota creada con éxito"))
}
