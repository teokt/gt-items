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

const SupportedVersion uint16 = 22

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

func (im *ItemManager) GetItemByID(itemID int) (*Item, error) {
	if itemID < 0 || itemID >= len(im.Items) {
		return nil, ErrItemNotFound
	}
	return &im.Items[itemID], nil
}

func (im *ItemManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return im.Deserialize(data)
}

func (im *ItemManager) Deserialize(data []byte) error {
	reader := memory.NewReader(data)

	if err := reader.Read(&im.Version); err != nil {
		return err
	}

	if im.Version > SupportedVersion {
		return ErrVersionNotSupported
	}

	var itemCount uint32
	if err := reader.Read(&itemCount); err != nil {
		return err
	}

	im.Items = make([]Item, itemCount)

	for i := range im.Items {
		item := &im.Items[i]
		if err := item.Deserialize(reader, int(im.Version)); err != nil {
			return err
		}
	}

	return nil
}
