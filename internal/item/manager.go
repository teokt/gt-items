package item

import (
	"errors"
	"os"

	"github.com/teokt/gt-items/internal/memory"
)

var (
	ErrVersionNotSupported = errors.New("version not supported")
	ErrItemNotFound        = errors.New("item not found")
)

const SupportedVersion uint16 = 21

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

	if im.Version > SupportedVersion {
		return ErrVersionNotSupported
	}

	var count uint32
	if err := r.Read(&count); err != nil {
		return err
	}

	im.Items = make([]Item, count)
	for i := range im.Items {
		item := &im.Items[i]
		if err := item.Deserialize(r, int(im.Version)); err != nil {
			return err
		}
	}

	return nil
}

func (im *ItemManager) GetByID(id int) (*Item, error) {
	if id < 0 || id >= len(im.Items) {
		return nil, ErrItemNotFound
	}
	return &im.Items[id], nil
}
