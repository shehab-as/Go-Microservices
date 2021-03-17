package services

import (
	"Go-Microservices/pkg/models"
	"Go-Microservices/pkg/repositories"
)

type ItemService struct {
	itemRepo	*repositories.ItemRepo 
}

func NewItemService(itemRepo *repositories.ItemRepo) *ItemService {
	return &ItemService {
		itemRepo: itemRepo,
	}
}

func (s* ItemService) GetSingleItem(ID int) (*models.Item, error) {
	item, _ := s.itemRepo.GetItemById(ID)
	return item, nil
}

func (s *ItemService) GetAllItems() (*models.ItemList, error) {
	itemList, _ := s.itemRepo.GetAllItems()
	return itemList, nil
}

func (s* ItemService) InsertItem(item *models.Item) (bool, error) {
	state, err := s.itemRepo.SaveItem(item)
	return state, err 
}

func (s* ItemService) DeleteItem(item *models.Item) (bool, error) {
	state, err := s.itemRepo.DeleteItem(item)
	return state, err
}

func (s* ItemService) UpdateItem(item *models.Item) (bool, error) {
	state, err := s.itemRepo.UpdateItem(item)
	return state, err 
}

func (s* ItemService) IsItemAvailable(ID int) bool {
	_, err := s.itemRepo.GetItemById(ID)
	if err != nil {
		return true
	}
	return false
}