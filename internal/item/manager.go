package item

import (
	"errors"
	"github.com/teokt/gt-items/internal/memory"
	"os"
)

type ItemManager struct {
	Version uint16
	Items   []Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		Version: 0,
		Items:   []Item{},
	}
}

func (im *ItemManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return im.Deserialize(data)
}

func (im *ItemManager) Deserialize(data []byte) error {
	r := memory.NewReader(data)

	if err := r.Read(&im.Version); err != nil {
		return err
	}

	// TODO: supported version check

	var count uint32
	if err := r.Read(&count); err != nil {
		return err
	}

	im.Items = make([]Item, count)
	for i := range im.Items {
		if err := im.Items[i].Deserialize(r, int(im.Version)); err != nil {
			return err
		}
	}

	return nil
}

func (im *ItemManager) GetByID(id int) (*Item, error) {
	if id < 0 || id >= len(im.Items) {
		return nil, errors.New("item not found")
	}
	return &im.Items[id], nil
}
