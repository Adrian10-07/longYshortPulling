package controllers

import (
	"encoding/json"
	"API-HEX-GO/src/Products/aplication"
	"API-HEX-GO/src/Products/domain"
	"API-HEX-GO/src/Products/infraestructure"
	"net/http"
	"strconv"
	"strings"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido en la URL", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID del producto inválido", http.StatusBadRequest)
		return
	}

	print(productID)

	var updatedProduct domain.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewUpdateProduct(repo)

	err = useCase.Execute(productID, &updatedProduct)
	if err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto actualizado correctamente"))
}
