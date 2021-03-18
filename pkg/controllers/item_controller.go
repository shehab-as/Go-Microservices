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

type ItemController struct {
	itemService	*services.ItemService
}

func NewItemController(itemService *services.ItemService) *ItemController {
	return &ItemController {
		itemService: itemService,
	}
}

func (h *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
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

func (h *ItemController) GetItemList(w http.ResponseWriter, r *http.Request) {
	var itemList *models.ItemList
	var err error 
	itemList, err = h.itemService.GetAllItems()
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(itemList)
}

func (h *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
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

func (h *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var err error 
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, err := strconv.Atoi(idStr)
	deleted, err := h.itemService.DeleteItem(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if deleted == true {
		fmt.Println("Deleted Item Successfully")
	}
}

func (h *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
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
