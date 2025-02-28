package controllers

import (
	"API-HEX-GO/src/Products/aplication"
	"API-HEX-GO/src/Products/infraestructure"
	"net/http"
	"strings"
)


func DeleteProductHandeler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Nombre del producto requerido", http.StatusBadRequest)
		return
	}

	productName := pathParts[2] 


	var Nombreproduct string = productName

	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewDeleteProduct(repo)

	if err := useCase.Execute(Nombreproduct); err != nil {
		http.Error(w, "Error al Borrar el producto", http.StatusInternalServerError)
		return
}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto elminado correctamente"))

}