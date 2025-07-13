package item

import (
	"fmt"
	"os"

	"github.com/teokt/gt-items/internal/memory"
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
		return fmt.Errorf("version '%d' is not supported", im.Version)
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
