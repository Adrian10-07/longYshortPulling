package controllers

import (
	"API-HEX-GO/src/pets/aplication"
	"API-HEX-GO/src/pets/infraestructure"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	repo := infraestructure.NewMySQLRepository()  
	useCase := aplication.NewGetPet(repo)
	pets, err := useCase.Execute()
	if err != nil {
		fmt.Printf("Error al obtener las mascotas: %v\n", err)
		http.Error(w, "Error al obtener las mascotas", http.StatusInternalServerError)
		return
	}

	petsJson, err := json.Marshal(pets)
	if err != nil {
		http.Error(w, "Error al procesar la respuesta JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(petsJson)

	fmt.Printf(string(petsJson))
}
