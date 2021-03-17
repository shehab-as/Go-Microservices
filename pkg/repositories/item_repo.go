package repositories

import (
	"fmt"
	"database/sql"
	"Go-Microservices/pkg/models"
)

type ItemRepo struct {
	dbClient *sql.DB
}

func NewItemRepo(dbClient *sql.DB) *ItemRepo {
	return &ItemRepo {
		dbClient: dbClient,
	}
}

func (i *ItemRepo) GetItemById(ID int) (*models.Item, error) {
	fmt.Println("GetItemById()")
	return &models.Item{}, nil
}

func (i *ItemRepo) GetAllItems() (*models.ItemList, error) {
	fmt.Println("GetAllItems()")
	return &models.ItemList{}, nil
}

func (i *ItemRepo) SaveItem(*models.Item) (bool, error) {
	fmt.Println("SaveItem()")
	return true, nil
}

func (i *ItemRepo) DeleteItem(*models.Item) (bool, error) {
	fmt.Println("DeleteItem()")
	return true, nil
}

func (i *ItemRepo) UpdateItem(*models.Item) (bool, error) {
	fmt.Println("UpdateItem()")
	return true, nil
}