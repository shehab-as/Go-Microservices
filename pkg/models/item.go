package models

type Item struct {
	ID		string		`json:"id"`
	Name	string	 	`json:"name"`
	Cost	float64	 	`json:"cost"`
	Details	string	 	`json:"details"`
}

type ItemList struct {
	ID		string		`json:"id"`
	Items 	[]Item 		`json:"items"`
}

type ItemRepository interface {
	GetItemById(ID int) (*Item, error)
	GetAllItems() (*ItemList, error)
	SaveItem(*Item) (bool, error)
	DeleteItem(ID int) (bool, error)
	UpdateItem(*Item) (bool, error)
}