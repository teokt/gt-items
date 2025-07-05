package item

import (
	"errors"
)

type ItemManager struct {
	Items []Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		Items: []Item{},
	}
}

func (im *ItemManager) LoadFromFile(filename string) error {
	return nil
}

func (im *ItemManager) GetByID(id int) (*Item, error) {
	if id < 0 || id >= len(im.Items) {
		return nil, ErrNotFound
	}
	return &im.Items[id], nil
}

var ErrNotFound = errors.New("item not found")
