package services

import (
	"strconv"
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

func (s* ItemService) IsItemAvailable(id int) (bool, error) {
	item, err := s.itemRepo.GetItemById(id)
	if err == nil && item != nil  {
		return true, nil
	}
	return false, err
}

func (s* ItemService) GetSingleItem(ID int) (*models.Item, error) {
	item, err := s.itemRepo.GetItemById(ID)
	return item, err
}

func (s *ItemService) GetAllItems() (*models.ItemList, error) {
	itemList, err := s.itemRepo.GetAllItems()
	return itemList, err
}

func (s* ItemService) InsertItem(item *models.Item) (bool, error) {
	state, err := s.itemRepo.SaveItem(item)
	return state, err 
}

func (s* ItemService) DeleteItem(id int) (bool, error) {
	var err error
	found, err := s.IsItemAvailable(id)
	if found == false {
		return false, err
	}
	state, err := s.itemRepo.DeleteItem(id)
	return state, err
}

func (s* ItemService) UpdateItem(item *models.Item) (bool, error) {
	var err error
	idInt, err := strconv.Atoi(item.ID)
	found, err := s.IsItemAvailable(idInt)
	if found == false {
		return false, err
	}
	state, err := s.itemRepo.UpdateItem(item)
	return state, err 
}

