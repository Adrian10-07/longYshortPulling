package controllers

import (
	"encoding/json"
	"API-HEX-GO/src/Products/aplication"
	"API-HEX-GO/src/Products/domain"
	"API-HEX-GO/src/Products/infraestructure"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	var updatedProduct domain.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewUpdateProduct(repo)

	// Implementación de long polling
	timeout := time.After(30 * time.Second)
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-timeout:
			http.Error(w, "Tiempo de espera agotado", http.StatusRequestTimeout)
			return
		case <-tick:
			err = useCase.Execute(productID, &updatedProduct)
			if err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Producto actualizado correctamente"))
				return
			}
		}
	}
}
