package main

import (
	"fmt"
	"golang-crud/controllers"
	"golang-crud/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)
	log.Println(fmt.Sprintf("Starting server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
