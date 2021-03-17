package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"Go-Microservices/pkg/db"
	"Go-Microservices/pkg/services"
	"Go-Microservices/pkg/controllers"
	"Go-Microservices/pkg/repositories"
)

func main() {

	db := db.ConnectDB()
	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := controllers.NewItemHandler(itemService)

	router := mux.NewRouter()
	fmt.Println("Server is running on Port 8080")

	router.HandleFunc("/api/v1/item", itemHandler.GetItemList).Methods("GET")
	router.HandleFunc("/api/v1/item", itemHandler.CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", itemHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/api/v1/item/{id}", itemHandler.DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item/{id}", itemHandler.GetItem).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}