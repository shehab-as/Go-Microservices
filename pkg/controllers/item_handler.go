package controllers

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"Go-Microservices/pkg/models"
	"Go-Microservices/pkg/services"
)

type ItemHandler struct {
	itemService	*services.ItemService
}

func NewItemHandler(itemService *services.ItemService) *ItemHandler {
	return &ItemHandler {
		itemService: itemService,
	}
}

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT - GET")

	var item *models.Item
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, err := strconv.Atoi(idStr)
	item, err = h.itemService.GetSingleItem(idInt); 
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) GetItemList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT - GET")

	var itemList *models.ItemList
	var err error 
	itemList, err = h.itemService.GetAllItems()
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(itemList)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT - POST")

	var itemToCreate models.Item
	var err error 
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &itemToCreate)
	created, err := h.itemService.InsertItem(&itemToCreate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if created == true {
		fmt.Println("Saved Item Successfully")
	}
	json.NewEncoder(w).Encode(itemToCreate)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT - DELETE")

	var itemToDelete models.Item
	var err error 
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &itemToDelete)
	deleted, err := h.itemService.DeleteItem(&itemToDelete)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if deleted == true {
		fmt.Println("Deleted Item Successfully")
	}
	json.NewEncoder(w).Encode(itemToDelete)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT - PUT")

	var itemToUpdate models.Item
	var err error 
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &itemToUpdate)
	updated, err := h.itemService.UpdateItem(&itemToUpdate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if updated == true {
		fmt.Println("Updated Item Successfully")
	}
	json.NewEncoder(w).Encode(itemToUpdate)
}
