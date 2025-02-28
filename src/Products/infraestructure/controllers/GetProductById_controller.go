package controllers

import (
	"API-HEX-GO/src/Products/aplication"
	"API-HEX-GO/src/Products/infraestructure"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewGetProductById(repo)

	product, err := useCase.Execute(id)
	if err != nil {
		http.Error(w, "Error al obtener el producto", http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	productJson, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Error al convertir el producto a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productJson)
}
