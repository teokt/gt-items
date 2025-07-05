package item

type ItemManager struct {
	items []Item
}

func NewManager() *ItemManager {
	return &ItemManager{
		items: []Item{},
	}
}
