package routes

import (
	"net/http"
	"API-HEX-GO/src/Products/infraestructure/controllers"
	"API-HEX-GO/src/Products/infraestructure/adapters/http/middleware"
)

func SetupRoutes() {
	http.HandleFunc("/products", middleware.CorsMiddleware(controllers.CreateProductHandler))
	http.HandleFunc("/view-products", middleware.CorsMiddleware(controllers.GetProductHandler))
	http.HandleFunc("/view-products-id/", middleware.CorsMiddleware(controllers.GetProductByIdHandler))
	http.HandleFunc("/delete-products/", middleware.CorsMiddleware(controllers.DeleteProductHandeler))
	http.HandleFunc("/update-products/", middleware.CorsMiddleware(controllers.UpdateProductHandler))
}
