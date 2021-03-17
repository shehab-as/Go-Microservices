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
	var err error
	var item models.Item
	err = i.dbClient.QueryRow("SELECT * FROM items WHERE id = ?", ID).Scan(&item.ID, &item.Name, &item.Cost, &item.Details)
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	return &item, nil
}

// TODO.
func (i *ItemRepo) GetAllItems() (*models.ItemList, error) {
	var err error
	rows, err := i.dbClient.Query("SELECT * FROM Item")
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	var itemList *models.ItemList
	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Cost, &item.Details)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		fmt.Printf("%+v\n", item)
	}
	return itemList, nil
}

func (i *ItemRepo) SaveItem(item *models.Item) (bool, error) {
	var err error 
	query, err := i.dbClient.Prepare("INSERT INTO Item(id, name, cost, details) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Printf("ERROR INSERT QUERY - %s", err)
		return false, err
	}
	query.Exec(item.ID, item.Name, item.Cost, item.Details)
	return true, nil
}

func (i *ItemRepo) DeleteItem(ID int) (bool, error) {
	var err error
	query, err := i.dbClient.Prepare("DELETE FROM Item WHERE id=?")
	if err != nil {
		fmt.Printf("ERROR DELETE QUERY - %s", err)
		return false, err
	}
	query.Exec(ID)
	return true, nil
}

func (i *ItemRepo) UpdateItem(item *models.Item) (bool, error) {
	var err error
	query, err := i.dbClient.Prepare("UPDATE Item SET name=?, cost=?, details=? WHERE id=?")
	if err != nil {
		fmt.Printf("ERROR UPDATE QUERY - %s", err)
		return false, err
	}
	query.Exec(item.Name, item.Cost, item.Details, item.ID)
	return true, nil
}